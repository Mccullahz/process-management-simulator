package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"process-management-simulator/cmd"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
)

// --> Initialization models---
type appState int

// stealing app state logic from bubbletea examples for sexiness
const (
	stateLoading appState = iota // initial loading
	stateProcessInput // for user input (numProcesses)
	stateMenu
	stateGenerated
	stateFCFS
	stateRR
)

type item string

func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return string(i) }

type model struct {
	processes    []cmd.Process
	scheduled    []cmd.ScheduledProcess
	timeSlices   []cmd.TimeSlice
	schduledRR   []cmd.ScheduledProcess
	processSnapshot []cmd.ProcessStateSnapshot
	cursor       int
	state        appState
	list         list.Model
	numProcesses int
	
	progress     progress.Model
	percent      float64
	width       int
	height      int
}

// --> main function ONLY STARTS the program
func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func initialModel() model {
	procs := []cmd.Process{}
	sched := []cmd.ScheduledProcess{}
	rrSched := []cmd.ScheduledProcess{}
	rrSlices := []cmd.TimeSlice{}

	items := []list.Item{
		item("First Come First Serve"),
		item("Round Robin"),
	}
	delegate := list.NewDefaultDelegate()
	l := list.New(items, delegate, 30, 14)
	l.Title = "Select a Scheduling View"

	prog := progress.New(progress.WithScaledGradient("#FF7CCB", "#FDFF8C"))

	return model{
		state:       stateLoading,
		processes:   procs,
		scheduled:   sched,
		timeSlices:  rrSlices,
		schduledRR:  rrSched,
		cursor:      0,
		list:        l,
		progress:    prog,
		percent:     0,
	}
}

// --> Bubbletea interface ---

func (m model) Init() tea.Cmd {
	if m.state == stateLoading {
		return tickCmd()
	}
	return nil
}

// update function to handle messages, e.g., key presses
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		// window resizing NOW WORKING :)
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	// tick message to update progress bar
	case tickMsg:
		if m.state == stateLoading {
			m.percent += 0.25
			if m.percent >= 1.0 {
				// adjust amount of processes generated, max burst time, and max arrival time
				m.processes = cmd.GenerateProcesses(5, 10, 5)
				m.scheduled, m.processSnapshot = cmd.FCFS(m.processes) // schedule and snapshots from FCFS
				m.schduledRR, m.timeSlices, m.processSnapshot = cmd.RR(m.processes, 2) // schedule, time quantum, and snapshots from round robin
				m.state = stateMenu
				return m, nil
			}
			return m, tickCmd()
		}
	// key press messages to handle user input
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			if m.state == stateMenu {
				return m, tea.Quit
			}
			m.state = stateMenu
			return m, nil
		case "enter":
			if m.state == stateMenu {
				i, ok := m.list.SelectedItem().(item)
				if ok {
					switch i {
					case "First Come First Serve":
						m.state = stateFCFS
					case "Round Robin":
						m.state = stateRR
					}
				}
			}
		case "esc", "backspace":
			m.state = stateMenu
		}
	}
	// update the list model if the state is in the menu
	if m.state == stateMenu {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}

	return m, nil
}

// view function to render (?) the bubbletea model
func (m model) View() string {
	// early exit if dimensions are not yet available
	if m.width == 0 || m.height == 0 {
		return ""
	}

	var b strings.Builder

	// header things, kinda like doiung css with lipgloss
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF7CCB"))

	header := headerStyle.Render("Process Management Simulator")

	// SEXY PROGRESS BAR :sunglasses:
	if m.state == stateLoading {
		loadingView := lipgloss.JoinVertical(
			lipgloss.Center,
			m.progress.ViewAs(m.percent),
			"Loading process data...",
		)

		return lipgloss.Place(
			m.width, m.height,
			lipgloss.Center, lipgloss.Center,
			loadingView,
		)
	}

	// horizontally center the header
	header = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, header)
	b.WriteString(header + "\n")
	b.WriteString(strings.Repeat("\n", 2))

	// so in order to center processes (body), we need to build a main body with lipgloss
	var body strings.Builder

	// always show the unscheduled processes
	body.WriteString("Unscheduled Generated Processes:\n")
	body.WriteString("PID  Arrival  Burst\n")
	for _, p := range m.processes {
		body.WriteString(fmt.Sprintf("%3d  %7d  %5d\n", p.PID, p.ArrivalTime, p.BurstTime))
	}
	body.WriteString("\n")

	// FCFS VIEW
	if m.state == stateFCFS {
		body.WriteString("First Come First Served Scheduled:\n")
		body.WriteString("PID  Arrival  Burst  Start  Complete  Turnaround  Waiting\n")
		for _, p := range m.scheduled {
			body.WriteString(fmt.Sprintf("%3d  %7d  %5d  %5d  %8d  %10d  %7d\n",
				p.PID, p.ArrivalTime, p.BurstTime, p.StartTime, p.CompletionTime, p.TurnaroundTime, p.WaitingTime))
		}
		body.WriteString("\n[esc] to return to menu")

	// RR VIEW
	} else if m.state == stateRR {
		body.WriteString("Round Robin Scheduled:\n")
		body.WriteString("Time Quantum: 2\n")
		body.WriteString("PID  Arrival  Burst  Start  Complete\n")
		for _, ts := range m.timeSlices {
			var original cmd.Process
			for _, p := range m.processes {
				if p.PID == ts.PID {
					original = p
					break
				}
			}
			body.WriteString(fmt.Sprintf("%3d  %7d  %5d  %5d  %3d\n",
				ts.PID, original.ArrivalTime, original.BurstTime, ts.Start, ts.End))
		}
		body.WriteString("\n[esc] to return to menu")

	// LIST MENU VIEW
	} else {
		// center the list view horizontally
		listView := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, m.list.View())
		body.WriteString(listView)
	}

	body.WriteString("\n\nPress [q] to quit.")

	// horizontally center everything except the progress bar
	centeredContent := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, body.String())

	// add centered header and some spacing above the body
	final := header + "\n\n" + centeredContent

	// vertically center the whole view
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Top, final)
}
// necessary tick message for progress bar
type tickMsg struct{}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second/4, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

// ansi color-safe text centering
func centerText(s string) string {
	width := 80
	padding := (width - len(stripANSI(s))) / 2
	if padding < 0 {
		padding = 0
	}
	return strings.Repeat(" ", padding) + s
}
// i am so sorry for this regex. match esc, match literal [, match 0-9 to ; for colors, match literal m that ends color codes
var ansi = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSI(s string) string {
	return ansi.ReplaceAllString(s, "")
}


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
	stateLoading appState = iota
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
	cursor       int
	state        appState
	list         list.Model
	progress     progress.Model
	percent      float64
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

	case tickMsg:
		if m.state == stateLoading {
			m.percent += 0.25
			if m.percent >= 1.0 {
				m.processes = cmd.GenerateProcesses(5, 10, 5) // process data, burst (max), arrival (max)
				m.scheduled = cmd.FCFS(m.processes)
				m.schduledRR, m.timeSlices = cmd.RR(m.processes, 2) // schedule and time quantum from round robin
				m.state = stateMenu
				return m, nil
			}
			return m, tickCmd()
		}

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

	if m.state == stateMenu {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}

	return m, nil
}

// view function to render (?) the bubbletea model
func (m model) View() string {
	var b strings.Builder

	// header things, kinda like doiung css with lipgloss
	// TODO: determine size of the terminal, then use that to center the header
	var (
		headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF00FF")).
			Align(lipgloss.Center, lipgloss.Center).
			Background(lipgloss.Color("#000000")) // why is this not doing anything?
	)

	// SEXY PROGRESS BAR :sunglasses:
	if m.state == stateLoading {
		b.WriteString("\n\n")
		b.WriteString(centerText(m.progress.ViewAs(m.percent)) + "\n\n")
		b.WriteString(centerText("Loading process data..."))
		return b.String()
	}

	b.WriteString(headerStyle.Render("Process Management Simulator") + "\n")
	b.WriteString(strings.Repeat("\n", 7) + "\n")

	// always show the unscheduled processes
	b.WriteString("Unscheduled Generated Processes:\n")
	b.WriteString("PID  Arrival  Burst\n")
	for _, p := range m.processes {
		b.WriteString(fmt.Sprintf("%3d  %7d  %5d\n", p.PID, p.ArrivalTime, p.BurstTime))
	}
	b.WriteString("\n")

	// FCFS VIEW
	if m.state == stateFCFS {
		b.WriteString("First Come First Served Scheduled:\n")
		b.WriteString("PID  Arrival  Burst  Start  Complete  Turnaround  Waiting\n")
		for _, p := range m.scheduled {
			b.WriteString(fmt.Sprintf("%3d  %7d  %5d  %5d  %8d  %10d  %7d\n", // this is disgusting but functional, sorry world
				p.PID, p.ArrivalTime, p.BurstTime, p.StartTime, p.CompletionTime, p.TurnaroundTime, p.WaitingTime))
		}
		b.WriteString("\n[esc] to return to menu")

		// RR VIEW
	} else if m.state == stateRR {
		b.WriteString("Round Robin Scheduled:\n")
		b.WriteString("Time Quantum: 2\n")
		b.WriteString("PID  Arrival  Burst  Start  Complete\n")
		for _, ts := range m.timeSlices {
			var original cmd.Process
			for _, p := range m.processes {
				if p.PID == ts.PID {
					original = p
					break
				}
			}
			b.WriteString(fmt.Sprintf("%3d  %7d  %5d  %5d  %3d\n",
				ts.PID, original.ArrivalTime, original.BurstTime, ts.Start, ts.End))
		}
		b.WriteString("\n[esc] to return to menu")

	} else {
		b.WriteString(m.list.View())
	}

	b.WriteString("\n\nPress [q] to quit.")
	return b.String()
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


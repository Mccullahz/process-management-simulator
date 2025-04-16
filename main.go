package main

import (
	"fmt"
	"os"
	"strings"
	"process-management-simulator/cmd"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/bubbles/list"
)

// --> Initialization models--- 
type appState int
// stealing app state logic from bubbletea examples for sexiness
const (
	stateMenu appState = iota
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
}

func initialModel() model {
	procs := cmd.GenerateProcesses(5, 10, 5) // process data, burst (max), arrival (max)
	sched := cmd.FCFS(procs)
	rrSched, rrSlices := cmd.RR(procs, 2) // schedule and time quantum from round robin

	items := []list.Item{
		item("First Come First Serve"),
		item("Round Robin"),
	}
	delegate := list.NewDefaultDelegate()
	l := list.New(items, delegate, 30, 14)
	l.Title = "Select a Scheduling View"

	return model{
		processes:   procs,
		scheduled:   sched,
		timeSlices:  rrSlices,
		schduledRR:  rrSched,
		cursor:     0,
		state:      stateMenu,
		list:       l,
	}
}

// --> Bubbletea interface ---

func (m model) Init() tea.Cmd {
	return nil
}

// update function to handle messages, e.g., key presses
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
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

	b.WriteString(headerStyle.Render("Process Management Simulator") + "\n")
	b.WriteString(strings.Repeat("\n", 7) + "\n")

	// always show the unscheduled processesa
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


	}else {
		b.WriteString(m.list.View())
	}

	b.WriteString("\n\nPress [q] to quit.")
	return b.String()
}

// --> main function ONLY STARTS the program
func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}




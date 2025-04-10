package main

import (
	"fmt"
	"os"
	"strings"
	"process-management-simulator/cmd"
	tea"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
// --> Initialization --- 
type model struct {
	processes	[]cmd.Process
	scheduled	[]cmd.ScheduledProcess
	cursor		int
	showFCFS	bool
	showRR		bool
}

func initialModel() model {
	procs := cmd.GenerateProcesses(5, 10, 5) // processes, burst (max), arrival (max)
	sched := cmd.FCFS(procs)

	return model{
		processes:     procs,
		scheduled:     sched,
		cursor:        0,
		showFCFS: false,
		showRR: false,
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
		case "f":
			m.showFCFS = !m.showFCFS
		case "r":
			m.showRR = !m.showRR
		}
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
	b.WriteString(strings.Repeat("\n", 10) + "\n")

	if m.showFCFS { // FCFS VIEW
		b.WriteString("First Come First Served Scheduled:\n")
		b.WriteString("PID  Arrival  Burst  Start  Complete  Turnaround  Waiting\n")
		for _, p := range m.scheduled {
			b.WriteString(fmt.Sprintf("%3d  %7d  %5d  %5d  %8d  %10d  %7d\n",
				p.PID, p.ArrivalTime, p.BurstTime, p.StartTime, p.CompletionTime, p.TurnaroundTime, p.WaitingTime))
		}

		b.WriteString("\nPress [f] to go back to Generated Processes")
		b.WriteString("\nPress [r] to view Round Robin Schedule")
	} else if m.showRR { // RR VIEW --> FIXME: printing how we would do it in FCFS, this essentially only prints FCFS, not the round robin logic. 
		b.WriteString("Round Robin Scheduled:\n")
		b.WriteString("Time Quantum: 2\n")
		b.WriteString("PID  Arrival  Burst  Start  Complete  Turnaround  Waiting\n")
		for _, p := range m.scheduled {
			b.WriteString(fmt.Sprintf("%3d  %7d  %5d  %5d  %8d  %10d  %7d\n",
			p.PID, p.ArrivalTime, p.BurstTime, p.StartTime, p.CompletionTime, p.TurnaroundTime, p.WaitingTime))
		}

		b.WriteString("\nPress [r] to go back to Generated Processes")
		b.WriteString("\nPress [f] to view First Come First Serve Schedule")
	}else { // DEFAULT VIEW
		b.WriteString("Unscheduled Generated Processes:\n")
		b.WriteString("PID  Arrival  Burst\n")
		for _, p := range m.processes {
			b.WriteString(fmt.Sprintf("%3d  %7d  %5d\n", p.PID, p.ArrivalTime, p.BurstTime))
		}
		b.WriteString("\nPress [f] to view First Come First Serve Schedule")
		b.WriteString("\nPress [r] to view Round Robin Schedule")
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


package main

import (
	"fmt"
	"os"
	"strings"
	"process-management-simulator/cmd"
	tea"github.com/charmbracelet/bubbletea"
)

type model struct {
	processes         []cmd.Process
	scheduled         []cmd.ScheduledProcess
	cursor            int
	showScheduled     bool
}

func initialModel() model {
	procs := cmd.GenerateProcesses(5, 10, 5) // processes, burst (max), arrival (max)
	sched := cmd.FCFS(procs)

	return model{
		processes:     procs,
		scheduled:     sched,
		cursor:        0,
		showScheduled: false,
	}
}

// --- Bubbletea interface ---

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "s":
			m.showScheduled = !m.showScheduled
		}
	}
	return m, nil
}

func (m model) View() string {
	var b strings.Builder

	if m.showScheduled {
		b.WriteString("Scheduled (FCFS):\n")
		b.WriteString("PID  Arrival  Burst  Start  Complete  Turnaround  Waiting\n")
		for _, p := range m.scheduled {
			b.WriteString(fmt.Sprintf("%3d  %7d  %5d  %5d  %8d  %10d  %7d\n",
				p.PID, p.ArrivalTime, p.BurstTime, p.StartTime, p.CompletionTime, p.TurnaroundTime, p.WaitingTime))
		}
	} else {
		b.WriteString("Generated Processes:\n")
		b.WriteString("PID  Arrival  Burst\n")
		for _, p := range m.processes {
			b.WriteString(fmt.Sprintf("%3d  %7d  %5d\n", p.PID, p.ArrivalTime, p.BurstTime))
		}
		b.WriteString("\nPress [s] to view FCFS Schedule")
	}

	b.WriteString("\n\nPress [q] to quit.")
	return b.String()
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}


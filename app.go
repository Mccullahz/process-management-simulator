package main

import (
	"process-management-simulator/cmd"
	"fmt"
	"context"
)
// App is a struct that holds the processes and provides methods to manipulate them
type App struct{
	processes []cmd.Process
	state []cmd.ProcessStateSnapshot
}
// NewApp is the default Wails constructor for the App struct
func NewApp() *App {
	return &App{}
}

// runs the cmd.GenerateProcesses function to generate the process list and output the unscheduled processes
func (a *App) GeneratedProcesses() string {
	output := "Generated Processes:\n"
	output += "PID  Arrival  Burst\n"
	for _, p := range a.processes {
		output += fmt.Sprintf("%3d  %7d  %5d\n", p.PID, p.ArrivalTime, p.BurstTime)
	}
	return output
}
// runs the cmd.FCFS (first come first serve) function and formats the output for js to display
func (a *App) FCFS() string {
	scheduled, state := cmd.FCFS(a.processes)
	a.state = state

	output := "FCFS Scheduling Result:\n"
	output += "PID  Arrival  Burst  Start  Complete\n"
	for _, p := range scheduled {
		output += fmt.Sprintf("%3d  %7d  %5d  %5d  %8d\n",
			p.PID, p.ArrivalTime, p.BurstTime, p.StartTime, p.CompletionTime)
	}
	return output
}
// runs the cmd.RR (round robin) function and formats the output for js to display
func (a *App) RR() string {
	_, slices, state := cmd.RR(a.processes, 2)
	a.state = state

	output := "Round Robin Scheduling (Time Quantum=2):\n"
	output += "PID  Start  End\n"
	for _, s := range slices {
		output += fmt.Sprintf("%3d  %5d  %3d\n", s.PID, s.Start, s.End)
	}
	return output
}
// TODO: make a button for this in front end. 
// regenerates the process list
func (a *App) Regenerate() string {
	a.processes = cmd.GenerateProcesses(5, 10, 5)
	return "Processes regenerated."
}

// this is a dumb startup function, not actually doing anything, wails needs it in main.go though
func (a *App) startup(ctx context.Context) {
	a.processes = cmd.GenerateProcesses(5, 10, 5)
}

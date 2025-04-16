package main

import (
	"process-management-simulator/cmd"
	"fmt"
	"context"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) FCFS() string {
	procs := cmd.GenerateProcesses(5, 10, 5)
	scheduled := cmd.FCFS(procs)

	output := "FCFS Scheduling Result:\n"
	output += "PID  Arrival  Burst  Start  Complete\n"
	for _, p := range scheduled {
		output += fmt.Sprintf("%3d  %7d  %5d  %5d  %8d\n",
			p.PID, p.ArrivalTime, p.BurstTime, p.StartTime, p.CompletionTime)
	}
	return output
}

func (a *App) RR() string {
	procs := cmd.GenerateProcesses(5, 10, 5)
	_, slices := cmd.RR(procs, 2)

	output := "Round Robin Scheduling (q=2):\n"
	output += "PID  Start  End\n"
	for _, s := range slices {
		output += fmt.Sprintf("%3d  %5d  %3d\n", s.PID, s.Start, s.End)
	}
	return output
}

// this is a dumb startup function, not actually doing anything, wails needs it in main.go though
func (a *App) startup(ctx context.Context) {
}

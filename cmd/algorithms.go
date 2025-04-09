package cmd

import (
	"sort"
)

type ScheduledProcess struct {
	Process
	StartTime     int
	CompletionTime int
	TurnaroundTime int
	WaitingTime    int
}
// will need to comment this up a bit better in the near future, works but isnt the prettiest
func FCFS(processes []Process) []ScheduledProcess {
	// sort by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	currentTime := 0
	schedule := []ScheduledProcess{}

	for _, p := range processes {
		if currentTime < p.ArrivalTime {
			currentTime = p.ArrivalTime
		}
		start := currentTime
		completion := start + p.BurstTime
		turnaround := completion - p.ArrivalTime
		waiting := turnaround - p.BurstTime

		schedule = append(schedule, ScheduledProcess{
			Process:        p,
			StartTime:      start,
			CompletionTime: completion,
			TurnaroundTime: turnaround,
			WaitingTime:    waiting,
		})

		currentTime = completion
	}

	return schedule
}


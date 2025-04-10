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

// first come first serve scheduling algorithm, ezpz
func FCFS(processes []Process) []ScheduledProcess {
	// sort process list by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	currentTime := 0
	schedule := []ScheduledProcess{}
	// loop through all processes and calculate the start time, completion time, turnaround time, and waiting time, then append to schedule
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
		// set current time to completion time, where next process will start
		currentTime = completion
	}
	// return schedule, to be used in main.go output
	return schedule
}
// round robin scheduling algorithm
// TODO: FINISH HIM!!!
func RR(processes []Process) []ScheduledProcess {
	// again, sort process list by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	currentTime := 0
	schedule := []ScheduledProcess{}
	quantum := 2 // time slice for round robin, not sure what to really set this to, but 2 is a good starting point i think
	
	// similar to how we did fcfs above, but we need to keep track of the remaining burst time for each processand only allow them to run for the time quantum
	// loop through all processes and calculate the start time, completion time, turnaround time, and waiting time, then append to schedule
	for _, p := range processes {
		if currentTime < p.ArrivalTime {
			currentTime = p.ArrivalTime
		}
		start := currentTime
		// if the process has a burst time greater than the quantum, we need to set the completion time to the current time + quantum, EDIT: this is wrong, in round robin we can only incriment by the quantum.
		completion := start + quantum
		if p.BurstTime > quantum {
			// set the remaining burst time to the burst time - quantum
			p.BurstTime = p.BurstTime - quantum
			// set the current time to the completion time
			currentTime = completion
		} else {
			// if the process has a burst time less than or equal to the quantum, we need to set the completion time to the current time + burst time
			completion = start + p.BurstTime
			currentTime = completion
			p.BurstTime = 0
		}



	}
	return schedule
}

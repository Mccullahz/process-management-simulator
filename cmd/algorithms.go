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
// time slice struct for round robin
type TimeSlice struct {
	PID    int
	Start  int
	End    int
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
func RR(processes []Process, quantum int) ([]ScheduledProcess, []TimeSlice) {
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	n := len(processes)
	currentTime := 0
	completed := 0
	schedule := []ScheduledProcess{}
	queue := []Process{}
	timeSlices := []TimeSlice{}
	remaining := make(map[int]int)
	visited := make(map[int]bool)
	startTimes := make(map[int]int) // tracks first time a process starts

	for _, p := range processes {
		remaining[p.PID] = p.BurstTime
	}

	i := 0

	for completed < n {
		// add arriving processes to the queue
		for i < n && processes[i].ArrivalTime <= currentTime {
			if !visited[processes[i].PID] {
				queue = append(queue, processes[i])
				visited[processes[i].PID] = true
			}
			i++
		}

		if len(queue) == 0 {
			currentTime++
			continue
		}

		current := queue[0]
		queue = queue[1:]

		runTime := quantum
		if remaining[current.PID] < quantum {
			runTime = remaining[current.PID]
		}

		start := currentTime
		currentTime += runTime
		remaining[current.PID] -= runTime

		// record the time slice
		timeSlices = append(timeSlices, TimeSlice{
			PID:   current.PID,
			Start: start,
			End:   currentTime,
		})

		// track the first time this process was scheduled
		if _, ok := startTimes[current.PID]; !ok {
			startTimes[current.PID] = start
		}

		// add newly arrived processes during this time window --> is this working how i think it should be?
		for i < n && processes[i].ArrivalTime <= currentTime {
			if !visited[processes[i].PID] {
				queue = append(queue, processes[i])
				visited[processes[i].PID] = true
			}
			i++
		}

		if remaining[current.PID] > 0 {
			queue = append(queue, current)
		} else {
			completion := currentTime
			turnaround := completion - current.ArrivalTime
			waiting := turnaround - current.BurstTime

			schedule = append(schedule, ScheduledProcess{
				Process:        current,
				StartTime:      startTimes[current.PID],
				CompletionTime: completion,
				TurnaroundTime: turnaround,
				WaitingTime:    waiting,
			})
			completed++
		}
	}

	return schedule, timeSlices
}


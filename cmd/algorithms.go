package cmd

import (
	"sort"
)
// constants for process state tracking
type ProcessState string
const (
	StateNew 	ProcessState = "New"
	StateReady	ProcessState = "Ready"
	StateRunning	ProcessState = "Running"
	StateWaiting	ProcessState = "Waiting"
	StateTerminated	ProcessState = "Terminated"
)

// struct to hold process information
type ScheduledProcess struct {
	Process
	StartTime     int
	CompletionTime int
	TurnaroundTime int
	WaitingTime    int
	State	ProcessState
}
// struct to hold snapshot of process information
type ProcessStateSnapshot struct {
	Time int `json:"time"`
	PID int	`json:"pid"`
	New []int `json:"new"`
	Ready []int `json:"ready"`
	Running []int `json:"running"`
	Waiting []int `json:"waiting"`
	Terminated []int `json:"terminated"`
}

// time slice struct for round robin
type TimeSlice struct {
	PID    int
	Start  int
	End    int
}


// helper function to build a snapshot of the process state
func buildSnapshot(time int, processes []Process, activePID int, activeState ProcessState, completed []ScheduledProcess) ProcessStateSnapshot {
	newList := []int{}
	readyList := []int{}
	runningList := []int{}
	waitingList := []int{}
	terminatedList := []int{}

	for _, p := range processes {
		switch {
		case isTerminated(p.PID, completed):
			terminatedList = append(terminatedList, p.PID)
		case p.PID == activePID:
			switch activeState {
			case StateRunning:
				runningList = append(runningList, p.PID)
			case StateWaiting:
				waitingList = append(waitingList, p.PID)
			case StateTerminated:
				terminatedList = append(terminatedList, p.PID)
		}
		case p.ArrivalTime <= time:
			readyList = append(readyList, p.PID)
		default:
			newList = append(newList, p.PID)
		}
	}

	return ProcessStateSnapshot{
		Time:       time,
		PID:        activePID,
		New:        newList,
		Ready:      readyList,
		Running:    runningList,
		Waiting:    waitingList,
		Terminated: terminatedList,
	}
}
// helper function to check if a process is terminated, best to do this in a separate function to avoid duplication
func isTerminated(pid int, completed []ScheduledProcess) bool {
	for _, c := range completed {
		if c.PID == pid {
			return true
		}
	}
	return false
}



// first come first serve scheduling algorithm, ezpz
func FCFS(processes []Process) ([]ScheduledProcess, []ProcessStateSnapshot) {
	// sort process list by arrival time
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	currentTime := 0
	schedule := []ScheduledProcess{}
	snapshots := []ProcessStateSnapshot{}
	// loop through all processes and calculate the start time, completion time, turnaround time, and waiting time, then append to schedule
	for _, p := range processes {
		if currentTime < p.ArrivalTime {
			// need an idle snapshot to make waiting work
			snapshots = append(snapshots, buildSnapshot(currentTime, processes, -1, "", schedule))
			currentTime = p.ArrivalTime
		}
		// create a snapshot of the process state before scheduling
		snapshots = append(snapshots, buildSnapshot(currentTime, processes, p.PID, StateRunning, schedule))

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

		// create a snapshot of the process state after scheduling
		snapshots = append(snapshots, buildSnapshot(currentTime, processes, p.PID, StateTerminated, schedule))
	}
	// return schedule, to be used in main.go output
	return schedule, snapshots
}


// round robin scheduling algorithm
func RR(processes []Process, quantum int) ([]ScheduledProcess, []TimeSlice, []ProcessStateSnapshot) {
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].ArrivalTime < processes[j].ArrivalTime
	})

	n := len(processes)
	currentTime := 0
	completed := 0
	schedule := []ScheduledProcess{}
	queue := []Process{}
	timeSlices := []TimeSlice{}
	snapshots := []ProcessStateSnapshot{}

	// initialize maps to track remaining burst time, visited processes, and start times
	remaining := make(map[int]int) // tracks remaining burst time for each process
	visited := make(map[int]bool) // tracks if a process has been added to the rr queue
	startTimes := make(map[int]int) // tracks first time a process starts

	for _, p := range processes {
		remaining[p.PID] = p.BurstTime
		// new state snapshot
		snapshots = append(snapshots, buildSnapshot(currentTime, processes, p.PID, StateNew, schedule))
	}

	i := 0

	for completed < n {
		// add arriving processes to the queue
		for i < n && processes[i].ArrivalTime <= currentTime {
			if !visited[processes[i].PID] {
				queue = append(queue, processes[i])
				visited[processes[i].PID] = true
				// ready state snapshot
				snapshots = append(snapshots, buildSnapshot(currentTime, processes, processes[i].PID, StateReady, schedule))

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
		// running state snapshot
		snapshots = append(snapshots, buildSnapshot(currentTime, processes, current.PID, StateRunning, schedule))

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
				snapshots = append(snapshots, buildSnapshot(currentTime, processes, processes[i].PID, StateReady, schedule))
			}
			i++
		}

		if remaining[current.PID] > 0 {
			queue = append(queue, current)
			// waiting state snapshot
			snapshots = append(snapshots, buildSnapshot(currentTime, processes, current.PID, StateWaiting, schedule))
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
			// terminated state snapshot
			snapshots = append(snapshots, buildSnapshot(currentTime, processes, current.PID, StateTerminated, schedule))
			completed++
		}
	}

	return schedule, timeSlices, snapshots
}


package cmd

import (
	"math/rand"
)

type Process struct {
	PID         int
	BurstTime  int
	ArrivalTime int
}

// generate a set of processes with PID, Burst Time, and Arrival Time
func GenerateProcesses(numProcesses int, maxBurstTime int, maxArrivalTime int) []Process {
	processes := make([]Process, numProcesses)
	for i := 0; i < numProcesses; i++ {
		processes[i] = Process{
			PID:         i + 1,
			BurstTime:   rand.Intn(maxBurstTime) + 1, // burst time between 1 and maxBurstTime
			ArrivalTime: rand.Intn(maxArrivalTime),    // arrival time between 0 and maxArrivalTime
		}
	}
	return processes
}

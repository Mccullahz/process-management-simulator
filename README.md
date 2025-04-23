# CS3200 Operating Systems - Process Management Simulator
hi nate, this line is a test for your pulling :) main Branch for GUI  & linux/amd64 branch for CLI tool
This project is a ..... add more description etc etc 

# 🧠 Operating System Process Management Simulator

A desktop application that simulates how an operating system manages processes using classic CPU scheduling algorithms. Built using [Wails](https://wails.io/) for a modern, lightweight cross-platform GUI experience.

## 🚀 Overview

This simulator demonstrates how an OS handles process scheduling through:

- **First Come First Served (FCFS)**
- **Round Robin (RR)**

Its goal is to visually and functionally simulate CPU scheduling behavior, calculating key performance metrics based on user input, and displaying the results in a clear and structured format.

---

## 🛠️ Backend Boilerplate (Go)

The backend is written in Go, using Wails’ bindings to bridge frontend JavaScript/HTML and backend logic. Here’s a sample backend boilerplate:

```go
package main

import (
	"context"
	"time"
)

// App struct
type App struct{}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called at application startup
func (a *App) Startup(ctx context.Context) {
	// Perform any initialization here
}

// FCFS simulates First Come First Serve scheduling
func (a *App) FCFS(processes []Process) []Result {
	// Implement FCFS logic here
	return simulateFCFS(processes)
}

// RoundRobin simulates Round Robin scheduling
func (a *App) RoundRobin(processes []Process, quantum int) []Result {
	// Implement Round Robin logic here
	return simulateRR(processes, quantum)
}

// Example data types
type Process struct {
	PID         string  `json:"pid"`
	ArrivalTime float64 `json:"arrival_time"`
	BurstTime   float64 `json:"burst_time"`
}

type Result struct {
	PID            string  `json:"pid"`
	WaitingTime    float64 `json:"waiting_time"`
	TurnaroundTime float64 `json:"turnaround_time"`
	FinishTime     float64 `json:"finish_time"`
}
💻 How to Run the Simulator
Prerequisites
•	Go 1.21+
•	Node.js & npm
•	Wails CLI (go install github.com/wailsapp/wails/v2/cmd/wails@latest)
Setup
bash
CopyEdit
git clone https://github.com/your-username/process-scheduler-simulator.git
cd process-scheduler-simulator
wails dev
This launches the simulator in development mode.
________________________________________
🧾 Input Instructions
•	Choose the scheduling algorithm via the dropdown menu.
•	Enter the number of processes or paste custom data.
•	For Round Robin, specify the quantum time.
•	Click Simulate to run the algorithm and view results.
________________________________________
📊 Metrics Calculated
Metric	Description
Waiting Time	Time a process spends waiting in the queue before execution begins.
Turnaround Time	Total time from process arrival to completion (TAT = Finish - Arrival).
Finish Time	Time when a process finishes execution.
Average Metrics	Average waiting and turnaround times across all processes.
These metrics are calculated per algorithm and updated live on the UI after simulation.
________________________________________
📌 Features
•	FCFS and Round Robin algorithms
•	Input validation and dynamic UI
•	Real-time simulation with metric summaries
•	Modular backend in Go using Wails




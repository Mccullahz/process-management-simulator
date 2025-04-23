# Operating System Process Management Simulator

A desktop and web application that simulates how an operating system manages processes. With a GUI tool Built using [Wails](https://wails.io/) for a modern cross-platform GUI experience.

## Overview

This simulator demonstrates how an OS handles process scheduling through:

- **First Come First Served (FCFS)**
- **Round Robin (RR)**

Its goal is to both visually and functionally simulate CPU scheduling behavior. This is done by taking a set amout of processes from a user and displaying the results in a correct and quite *spiffy* format.

This simulator works by taking a set of processes and simulating the scheduling of those processes using the FCFS and Round Robin algorithms on the backend. The simulator generates a set of `n` processes with random arrival and burst times, and then simulates the scheduling of those processes using the selected algorithm. The results are displayed in a table format, showing the start and end times for each process, along with a color coded display showcasing the current queue each process is in.

---

## Backend (Go)

The backend is written in Go using Wails’ bindings to bridge frontend JavaScript/HTML and backend logic. The bindings can be found in app.go,:
# Bindings (app.go):
```go
type App struct {
	processes []cmd.Process
	state     []cmd.ProcessStateSnapshot
}
func NewApp() *App {
    return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.processes = cmd.GenerateProcesses(5, 10, 5)
}

```
- Regenerate(count int): Generates a new list of random processes with configurable count.

- GeneratedProcesses(): Returns a textual summary of all current processes, including PID, arrival time, and burst time.

- FCFS(): Executes the First-Come-First-Serve scheduling algorithm on the current processes and returns a formatted string showing each process’s scheduling timeline.

- RR(): Executes the Round Robin algorithm with a fixed time quantum of 2 and returns each time slice per process.

- GetState(): Returns a full process state timeline snapshot for visualization.

These bindings ensure that the GUI remains reactive and displays current scheduling results and queue switches dynamically.

# Algorithms (cmd/algorithms.go):
The core CPU scheduling algorithms are implemented in the cmd package. This includes:

- **FCFS (First-Come-First-Serve)

    - Processes are sorted by arrival time and run in that order.

    - Each process snapshot includes arrival, start, and completion times.

    - A visual timeline is maintained for later state-based rendering.

- **Round Robin (RR)

    - Processes share CPU time in fixed time slices (quantum = 2).

    - Tracks and returns a timeline of individual time slices.

    - Simulates context switching and manages process queues at each tick.

- **Process Generation

    - GenerateProcesses(count, maxArrival, maxBurst) creates a reproducible list of pseudo-random processes.

    - Useful for testing both the CLI and GUI without needing manual input.

- **Process State Snapshots

    - Both algorithms track queue transitions: New -> Ready -> Running -> Waiting -> Terminated.

State is returned to the frontend for step-by-step visual replay of scheduling events.

## Shared Across GUI and CLI
Both the CLI version (Bubbletea based) and the GUI version (Wails based) share the same backedn logic found in the cmd package. This ensures consistency across both interfaces and simplifies the (short term) development life cycle.

# Installation and Usage

💻 How to Run the Simulator
1. Install the latest build of Process-Management-Simulator from this repository.
2. Run the downloaded executable.
3. The GUI will open, displaying the main interface :)
4. Choose the scheduling algorithm (FCFS or Round Robin) from the dropdown menu.
5. Enter the number of processes.
6. Choose "Simulate" to run the selected algorithm.
7. The results will be displayed in a table format, showing the start and end times for each process.

8. Remember to check the linux CLI build as well ;)


_______________________________________
🧾 Input Instructions
•	Choose the scheduling algorithm via the dropdown menu.
•	Enter the number of processes or paste custom data.
•	For Round Robin, specify the quantum time.
•	Click Simulate to run the algorithm and view results.
________________________________________
📌 Features
•	FCFS and Round Robin algorithms
•	Input validation and dynamic UI
•	Real-time process scheduling
•	Modular backend in Go using Wails




# Operating System Process Management Simulator CLI Edition

A terminal based application that simulates how an operating system manages processes using only the command line. Built using [Bubbletea](https://github.com/charmbracelet/bubbletea) and [Lipgloss](https://github.com/charmbracelet/lipgloss), providing a linux experience that (maybe more) stylish as it is functional.

## Overview

This simulator demonstrates how an OS handles process scheduling through:

- **First Come First Served (FCFS)**
- **Round Robin (RR)**

Its goal is to visually and interactively simulate CPU scheduling behavior in your terminal. The user defines a number of processes and the tool simulates scheduling them in real time with dynamic state transitions and visual queus.

The simulator generates a set of `n` processes with randomized arrival and burst times then displays each scheduling step as the algorithm runs showing which process is **ready**, **running**, **waiting**, or **terminated** all in a fully color coded interface.

## Backend Go

The CLI backend is written in Go and uses the exact same scheduling logic as the GUI version shared via the `cmd` package.

### Bindings (main.go) and CLI models:

Unlike our GUI counterpart, the CLI uses Bubbletea's model update view architecture to handle simulation logic and user interaction:

```go
type model struct {
    state         []cmd.ProcessStateSnapshot
    processes     []cmd.Process
    currentView   viewType
    selectedAlgo  algorithm
    numProcesses  int
    readyQueue    []cmd.Process
    running       *cmd.Process
    waitingQueue  []cmd.Process
    terminated    []cmd.Process
}
```

- `Startup logic` initializes the application with a loading screen and takes user input for the number of processes  
- `Update logic` handles menu selections, process generation, and simulation steps per algorithm  
- `View logic` displays each queue using lipgloss styled layouts with boxes for Ready, Running, Waiting, and Terminated  

Each update cycle animates the state transitions allowing users to watch the simulation unfold tick by tick

## Algorithms (cmd/algorithms.go)

The same scheduling logic is used here and shared with the Wails GUI

- **FCFS First Come First Serve**  
  - Processes are executed in the order of arrival  
  - Tracks start and completion times  
  - Generates a step by step queue state for terminal playback

- **Round Robin RR**  
  - Uses fixed time quantum default is 2  
  - Tracks time slices per process  
  - Maintains state transitions for terminal visualization

- **Process Generation**  
  - `GenerateProcesses(count, maxArrival, maxBurst)` creates reproducible testable process lists for simulation

- **Process State Snapshots**  
  - Every tick records the current state of each process  
  - Used to animate the simulation in a visually appealing manner

## Shared Across CLI and GUI

Both the GUI Wails and CLI Bubbletea versions share the same logic through the `cmd` package. This ensures consistent results and behavior across interfaces while simplifying development and debugging.

## Installation and Usage

### How to Run the CLI Simulator

## Prerequisite 
1. Go version 1.20 or higher
2. Linux / Unix / Bash Shell

## Installation
1. The simplest method to running this application is to install the executable directly from this repository, and then run:
    ```bash
    ./Process-Management-Simulator
    ```
2. Alternatively, you can clone the repository and run it locally following the steps below.
    a. Clone the repository  
    ```bash
    git clone https://github.com/Mccullahz/Process-Management-Simulator/tree/linux/amd64
    ```
    b. Change into the directory  
    ```bash
    cd Process-Management-Simulator
    ```
    c. Build the application  
    ```bash
    go build
    ```
    d. Run the application  
    ```bash
    ./Process-Management-Simulator
    ```

## Usage
1. Use the menu to  
   - Select an algorithm FCFS or Round Robin  
   - Enter the number of processes  
   - Watch the simulation animate in real time  

2. Exit anytime with `ctrl+c`, or `q`

## Input Instructions

- Select the scheduling algorithm via the interactive menu  
- Enter the number of processes 5 to 20 is a good range  
- For RR the time quantum is fixed to 2 ticks configurable via code  
- Hit Simulate to begin the animation  
- Enjoy the fully styled visual state transitions of each process

## Features

- Fully interactive terminal UI using Bubbletea  
- FCFS and Round Robin algorithms with visual timeline  
- Process state visualization Ready Running Waiting Terminated  
- Dynamic animations with each CPU tick  
- Modular backend shared with GUI version  
- Styled with Lipgloss for terminal glam  


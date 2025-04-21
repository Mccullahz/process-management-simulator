**Project Overview:**

- In this term project, you will develop a Process Management Simulator that mimics the behavior of processes in an operating system. The goal is to simulate how processes move between various queues in the system and calculate key scheduling metrics, such as waiting time, response time, and turnaround time.\

This simulator should randomly generate a set of processes with different burst times and arrival times, allowing you to observe and calculate process metrics under different scheduling algorithms (e.g., First-Come, First-Served (FCFS), Shortest Job Next (SJN), Round Robin (RR), etc.). The project aims to help you understand how operating systems manage processes and allocate CPU resources.\

### Project Requirements:
**1. Process Generation:**\
- Randomly generate a set of processes with the following attributes:
    - **Process ID (PID):** A unique identifier for each process.
    - **Burst Time:** The CPU time required by the process for execution.
    - **Arrival Time:** The time at which the process arrives in the ready queue.
- Ensure that burst and arrival times are generated randomly within reasonable limits. 

**2. Queue Management:** \
The simulator should model the movement of processes through the five primary queues in the operating system. These queues will represent different states of the process: \
    - **New Queue:** Processes that are being created and initialized.
    - **Ready Queue:** Processes that are ready to execute and are waiting for CPU time.
    - **Running Queue:** The process currently being executed on the CPU.
    - **Waiting Queue:** Processes that are waiting for some event (e.g., I/O completion) before they can resume execution.
    - **Terminated Queue:** Processes that have completed execution and are no longer active.

Implement the movement of processes between these queues based on the events occurring during the simulation (e.g., process creation, CPU time allocation, I/O operations).

**3. Metrics Calculation:**\

For each process, calculate the following scheduling metrics:\
    -**Waiting Time:** The total time a process spends in the ready queue before execution.
`Waiting Time = Turnaround Time - Burst Time`
    - **Response Time:** The time from when the process arrives in the ready queue until it first gets executed.
`Response Time = First Execution Time - Arrival Time`
    - **Turnaround Time:** The total time spent by a process from its arrival to its completion.
`Turnaround Time = Completion Time - Arrival Time`

**4. Randomized Scheduling:** You are required to implement at **least one scheduling algorithm** for this project, but implementing multiple algorithms will enhance your project and provide a more comprehensive comparison. It's recommended to explore different algorithms, such as **First-Come, First-Served (FCFS), Shortest Job Next (SJN), Round Robin (RR)**, or others, to observe how each handles process scheduling and resource allocation. The more algorithms you include, the more thorough your analysis will be.\

**5. Scheduling Algorithms to Implement:**\
    - **First-Come, First-Served (FCFS):** Processes are executed in the order they arrive in the ready queue.
    - **Shortest Job Next (SJN):** The process with the shortest burst time is executed next.
    - **Round Robin (RR):** Each process gets a fixed time slice, and if not completed, it moves back to the ready queue.
    - **Priority Scheduling:** Processes are executed based on priority, with the highest priority executed first.
    - Additional algorithms of your choice can be added for comparison.

**6. Visualization (Optional but Encouraged):**\
- Include a simple visualization or logging feature that tracks the movement of processes through the queues over time.
- This can include a timeline, tables, or textual descriptions showing when processes enter the ready/waiting queues, when they are executed, and when they are completed.

**7. Output Format:**
- After simulating process management, the simulator should output:
    - A table summarizing the Process ID, Arrival Time, Burst Time, Waiting Time, Response Time, Turnaround Time, and Completion Time for each process.
    - The average waiting time, response time, and turnaround time across all processes.
    - A comparison of the performance of different scheduling algorithms, including average waiting times and turnaround times for each algorithm (optional.)

### Project Guidelines:
**1. Programming Language:**\
- You may use any programming language that you are comfortable with (e.g., Python, Java, C++), as long as the project requirements are met.

**2. User Interface:**\
- A **command-line interface (CLI)** is sufficient for this project. The user should be able to input the number of processes to simulate. Alternatively, the simulator can randomly generate the number of processes, along with their arrival and burst times.
    - **Optional:** A graphical user interface (GUI) can be implemented but not required.

**3. Documentation:**\
- Provide a detailed README file that includes:
    - An explanation of how the simulator works.
    - Instructions on how to run the simulator and provide input(s).
    - A summary of the metrics and how they are calculated.
- Submit well-commented source code along with any supporting files.

### Deliverables:
- 1. Source code for the simulator.
- 2. A README file that explains how to run the program and what the output means.
- 3. Each group is required to present their project in class. The exact date for the presentation will be provided at a later time, either in class or via email.

### Evaluation Criteria:
- **Correctness:** Does the simulator accurately model process management, calculate metrics, and simulate the process movement through queues?
- **Usability:** Is the user interface intuitive and easy to use?
- **Code Quality:** Is the code clean, well-structured, and well-documented?- **Creativity:** Are there additional features, such as visualizations or optimizations, that enhance the simulation?
- **Report:** How well does the report explain the scheduling algorithms, their results, and the comparison between them?
- **Presentation:** Each group will be required to demonstrate their project to the professor in a one-on-one session. This will not involve presenting in front of the entire class. No PowerPoint presentation is required; the focus should be on a live demo of the simulator, showcasing its functionality, the algorithms implemented, and the results.

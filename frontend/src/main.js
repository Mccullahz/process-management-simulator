{/*I AM NOT A JS KING ------- UP TO YOU GUYS TO MAKE THIS SUCKER PRETTY */}
import './style.css';
import './app.css';

import {FCFS, RR, GeneratedProcesses, Regenerate, GetState} from '../wailsjs/go/main/App';

{/* VERY VERY ROUGH HOW TO OUTPUT ALGOS IN WINDOW, CAN CHANGE AS MUCH AS NEEDED*/}

async function simFCFS() {
  try {
    const GenProcesses = await GeneratedProcesses();
    const fcfsResult = await FCFS();
    const state = await GetState();
    const processes = GenProcesses.split('\n');
    const res1 = document.getElementById("result1");
    res1.innerHTML = ''; 
    //Generates a line for each processes (left side)
    processes.forEach((line, i) => {
      const div = document.createElement('div');
      const parts = line.trim().split(/\s+/);
      const pid = parseInt(parts[0]);
      div.id = `Proc-${pid}`; //used to ID this div later for color effects
      div.textContent = line;
      div.style.fontSize = '20px';
      div.style.webkitTextStroke = '.5px black';
      res1.appendChild(div);
    });
    const FCFSProc = fcfsResult.split('\n');
    const res2 = document.getElementById("result2");
    res2.innerHTML = ''; 
    //Generates a line for the state changes (right side)
    FCFSProc.forEach((line, i) => {
      const div = document.createElement('div');
      const parts = line.trim().split(/\s+/);
      const pid = parseInt(parts[0]);
      if (!isNaN(pid)) { //prevents the header and sub headers of the tables to stay visible
        div.id = `FCFSProc-${pid}`;
        div.style.opacity = '0';
      }
      div.style.fontSize = '20px';
      div.style.webkitTextStroke = '.5px black';
      div.textContent = line;
      res2.appendChild(div);
    });
    //goes through the snapshots and then gets changes the color of processes based on state
    state.forEach((snapshot, timeIndex) => {
      setTimeout(() => {
      snapshot.new.forEach((PID) => {
        const target = document.getElementById(`Proc-${PID}`);
        if (target) {
          target.style.color = 'white';
        }
      });
      snapshot.ready.forEach((PID) => {
        const target = document.getElementById(`Proc-${PID}`);
        if (target) {
          target.style.color = 'yellow';
        }
      });
      snapshot.running.forEach((PID) => {
        const target = document.getElementById(`Proc-${PID}`);
        if (target) {
          target.style.color = 'green';
        }
      });
      snapshot.terminated.forEach((PID) => {
        const target = document.getElementById(`Proc-${PID}`);
        if (target) {
          target.style.color = 'red';
          const altTarget = document.getElementById(`FCFSProc-${PID}`)
          altTarget.style.opacity = '100'
        }
      });
      }, (timeIndex)*500)
    });

  } catch (err) {
    console.error("Error in simFCFS:", err);
  }
}

async function simRR() {
  try {
    const GenProcesses = await GeneratedProcesses();
    const RRResult = await RR();
    const state = await GetState();
    console.log(state)
    const processes = GenProcesses.split('\n');
    const res1 = document.getElementById("result1");
    res1.innerHTML = ''; 
     //Generates a line for each processes (left side)
    processes.forEach((line, i) => {
      const div = document.createElement('div');
      const parts = line.trim().split(/\s+/);
      const pid = parseInt(parts[0]);
      div.id = `Proc-${pid}`;
      div.textContent = line;
      div.style.fontSize = '20px';
      div.style.webkitTextStroke = '.5px black';
      res1.appendChild(div);
    });
    const RRProc = RRResult.split('\n');
    const res2 = document.getElementById("result2");
    res2.innerHTML = ''; 
    //Generates a line for the state changes (right side)
    RRProc.forEach((line, i) => {
      const div = document.createElement('div');
      const parts = line.trim().split(/\s+/);
      const pid = parseInt(parts[0]);
      if (!isNaN(pid)) {
        div.id = `RRProc-${i-2}`;
        div.style.opacity = '0';
      }
      div.style.fontSize = '20px';
      div.style.webkitTextStroke = '.5px black';
      div.textContent = line;
      res2.appendChild(div);
    });
    //vars to monitor state change
    let OpIndex=0
    let prevRun=-1
    //goes through the snapshots and then gets changes the color of processes based on state
    state.forEach((snapshot, timeIndex) => {
      setTimeout(() => {
      snapshot.new.forEach((PID) => {
        const target = document.getElementById(`Proc-${PID}`);
        if (target) {
          target.style.color = 'white';
        }
      });
      snapshot.ready.forEach((PID) => {
        const target = document.getElementById(`Proc-${PID}`);
        if (target) {
          target.style.color = 'yellow';
        }
      });
      snapshot.running.forEach((PID) => {
        if(prevRun != PID){
          const alttarget = document.getElementById(`RRProc-${OpIndex}`);
          alttarget.style.opacity = '100'
          prevRun = PID
          OpIndex += 1;
        }
        const target = document.getElementById(`Proc-${PID}`);
        if (target) {
          target.style.color = 'green';
        }
      });
      snapshot.waiting.forEach((PID) => {
        const target = document.getElementById(`Proc-${PID}`);
        if (target) {
          target.style.color = 'orange';
        }
      });
      snapshot.terminated.forEach((PID) => {
        const target = document.getElementById(`Proc-${PID}`);
        if (target) {
          target.style.color = 'red';
        }
      });
      }, (timeIndex)*500)
    });
  } catch (err) {
    console.error("Error in simFCFS:", err);
  }
}

window.Regenerate = function () {
  Regenerate().then((result) => {}).catch(console.error);
};

window.Simulate = function () {
  const selectedAlgo = document.getElementById("algoSelect").value;
  if (selectedAlgo === "FCFS") {
    simFCFS();
  } else if (selectedAlgo === "RR") {
    simRR();
  }
};

document.getElementById("algoSelect").addEventListener("change", function () {})
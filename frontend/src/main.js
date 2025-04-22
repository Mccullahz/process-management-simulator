{/*I AM NOT A JS KING ------- UP TO YOU GUYS TO MAKE THIS SUCKER PRETTY */}
import './style.css';
import './app.css';

import {FCFS, RR, GeneratedProcesses, Regenerate, GetState} from '../wailsjs/go/main/App';

{/* VERY VERY ROUGH HOW TO OUTPUT ALGOS IN WINDOW, CAN CHANGE AS MUCH AS NEEDED*/}


function simFCFS(){
  Promise.all([GeneratedProcesses(), FCFS(), GetState()]).then((result) => {
    const processes = result[0].split('\n');
    const res1 = document.getElementById("result1");
    res1.innerHTML = ''; 
    processes.forEach((line, i) => {
      const div = document.createElement('div');
      const parts = line.trim().split(/\s+/);
      const pid = parseInt(parts[0]);
      div.id = `Proc-${pid}`;
      div.textContent = line;
      res1.appendChild(div);
    });
    const FCFSProc = result[1].split('\n');
    const res2 = document.getElementById("result2");
    res2.innerHTML = ''; 
    FCFSProc.forEach((line, i) => {
      const div = document.createElement('div');
      const parts = line.trim().split(/\s+/);
      const pid = parseInt(parts[0]);
      if (!isNaN(pid)) {
        div.id = `FCFSProc-${pid}`;
        div.style.opacity = '0';
      }
      div.textContent = line;
      res2.appendChild(div);
    });
    result[2].forEach((snapshot, timeIndex) => {
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
  }).catch(console.error);
}
function simRR(){
  Promise.all([GeneratedProcesses(), RR(), GetState()]).then((result) => {
    console.log(result[2])
    const processes = result[0].split('\n');
    const res1 = document.getElementById("result1");
    res1.innerHTML = ''; 
    processes.forEach((line, i) => {
      const div = document.createElement('div');
      const parts = line.trim().split(/\s+/);
      const pid = parseInt(parts[0]);
      div.id = `Proc-${pid}`;
      div.textContent = line;
      res1.appendChild(div);
    });
    const RRProc = result[1].split('\n');
    const res2 = document.getElementById("result2");
    res2.innerHTML = ''; 
    RRProc.forEach((line, i) => {
      const div = document.createElement('div');
      const parts = line.trim().split(/\s+/);
      const pid = parseInt(parts[0]);
      if (!isNaN(pid)) {
        div.id = `RRProc-${i-2}`;
        console.log(div.id)
        div.style.opacity = '0';
      }
      div.textContent = line;
      res2.appendChild(div);
    });
    let OpIndex=0
    let prevRun=-1
    result[2].forEach((snapshot, timeIndex) => {
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
      snapshot.terminated.forEach((PID) => {
        const target = document.getElementById(`Proc-${PID}`);
        if (target) {
          target.style.color = 'red';
        }
      });
      }, (timeIndex)*500)
    });
  }).catch(console.error);
}


window.Simulate = function () {
  const selectedAlgo = document.getElementById("algoSelect").value;
  if (selectedAlgo === "FCFS") {
    simFCFS();
  } else if (selectedAlgo === "RR") {
    simRR();
  }
};

document.getElementById("algoSelect").addEventListener("change", function () {})
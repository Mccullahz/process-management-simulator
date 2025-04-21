{/*I AM NOT A JS KING ------- UP TO YOU GUYS TO MAKE THIS SUCKER PRETTY */}
import './style.css';
import './app.css';

import {FCFS, RR, GeneratedProcesses, Regenerate} from '../wailsjs/go/main/App';

{/* VERY VERY ROUGH HOW TO OUTPUT ALGOS IN WINDOW, CAN CHANGE AS MUCH AS NEEDED*/}
window.runRegenerate = function () {
  //Regenerate().then((result) => {}).catch(console.error);
  GeneratedProcesses().then((result) => {
    document.getElementById("result1").innerHTML = result;
  }).catch(console.error);
  FCFS().then((result) => {
    document.getElementById("result2").innerHTML = result;
  }).catch(console.error);
  RR().then((result) => {
    document.getElementById("result3").innerHTML = result;
  }).catch(console.error);
  Regenerate().then((result) => {}).catch(console.error);
};
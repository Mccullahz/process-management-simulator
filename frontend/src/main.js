{/*I AM NOT A JS KING ------- UP TO YOU GUYS TO MAKE THIS SUCKER PRETTY */}
import './style.css';
import './app.css';

import logo from './assets/images/logo-universal.png';
import { FCFS, RR } from '../wailsjs/go/main/App';

document.querySelector('#app').innerHTML = `
  <img id="logo" class="logo">
  <div class="title">Process Scheduling Simulator</div>
  <div class="button-group">
    <button class="btn" onclick="runFCFS()">Run FCFS</button>
    <button class="btn" onclick="runRR()">Run Round Robin</button>
  </div>
  <pre class="result" id="result">← Select an algorithm to begin</pre>
`;
document.getElementById('logo').src = logo;

let resultElement = document.getElementById("result");



{/* VERY VERY ROUGH HOW TO OUTPUT ALGOS IN WINDOW, CAN CHANGE AS MUCH AS NEEDED*/}
window.runFCFS = function () {
    FCFS().then((result) => {
        resultElement.innerText = result;
    }).catch(console.error);
};

window.runRR = function () {
    RR().then((result) => {
        resultElement.innerText = result;
    }).catch(console.error);
};


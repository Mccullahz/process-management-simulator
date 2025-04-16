(function(){const o=document.createElement("link").relList;if(o&&o.supports&&o.supports("modulepreload"))return;for(const e of document.querySelectorAll('link[rel="modulepreload"]'))i(e);new MutationObserver(e=>{for(const t of e)if(t.type==="childList")for(const r of t.addedNodes)r.tagName==="LINK"&&r.rel==="modulepreload"&&i(r)}).observe(document,{childList:!0,subtree:!0});function s(e){const t={};return e.integrity&&(t.integrity=e.integrity),e.referrerpolicy&&(t.referrerPolicy=e.referrerpolicy),e.crossorigin==="use-credentials"?t.credentials="include":e.crossorigin==="anonymous"?t.credentials="omit":t.credentials="same-origin",t}function i(e){if(e.ep)return;e.ep=!0;const t=s(e);fetch(e.href,t)}})();const c="/assets/logo-universal.157a874a.png";function u(){return window.go.main.App.FCFS()}function d(){return window.go.main.App.RR()}document.querySelector("#app").innerHTML=`
  <img id="logo" class="logo">
  <div class="title">Process Scheduling Simulator</div>
  <div class="button-group">
    <button class="btn" onclick="runFCFS()">Run FCFS</button>
    <button class="btn" onclick="runRR()">Run Round Robin</button>
  </div>
  <pre class="result" id="result">\u2190 Select an algorithm to begin</pre>
`;document.getElementById("logo").src=c;let l=document.getElementById("result");window.runFCFS=function(){u().then(n=>{l.innerText=n}).catch(console.error)};window.runRR=function(){d().then(n=>{l.innerText=n}).catch(console.error)};

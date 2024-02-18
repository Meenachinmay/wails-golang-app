import { useState } from "react";
import logo from "./assets/images/logo-universal.png";
import "./App.css";
import { Greet, StartMonitoring, StopMonitoring } from "../wailsjs/go/main/App";

function App() {
  const [resultText, setResultText] = useState(
    "Please enter your name below 👇"
  );
  const [name, setName] = useState("");
  const updateName = (e: any) => setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);

  function greet() {
    Greet(name).then(updateResultText);
  }

  function startMonitoring() {
    StartMonitoring();
  }

  function stopMonitoring() {
    StopMonitoring();
  }

  return (
    <div id="App">
      <img src={logo} id="logo" alt="logo" />
      <div id="result" className="result">
        {resultText}
      </div>
      <div id="input" className="input-box">
        <input
          id="name"
          className="input"
          onChange={updateName}
          autoComplete="off"
          name="input"
          type="text"
        />
        <button className="btn" onClick={greet}>
          Greet
        </button>
        <button
          className=""
          style={{
            width: "100px",
            height: "50px",
            color: "white",
            backgroundColor: "blue",
            border: "none",
            marginRight: "10px",
            marginLeft: "10px",
            borderRadius: "10px",
            fontSize: "12px",
          }}
          onClick={startMonitoring}
        >
          Start monitoring
        </button>
        <button
          className=""
          style={{
            width: "100px",
            height: "50px",
            color: "white",
            backgroundColor: "blue",
            border: "none",
            borderRadius: "10px",
            fontSize: "12px",
          }}
          onClick={stopMonitoring}
        >
          Stop monitoring
        </button>
      </div>
    </div>
  );
}

export default App;

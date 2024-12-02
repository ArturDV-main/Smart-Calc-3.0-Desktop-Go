import { useState } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { Greet } from "../wailsjs/go/main/App";
import Graph from './Graph'
import { Line } from 'react-chartjs-2';

function App() {
    const [resultText, setResultText] = useState("Please enter your expression below 👇");
    const [expression, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);
    const updateResultText = (result: string) => setResultText(result);

    function greet() {
        Greet(expression).then(updateResultText);
    }

    return (
        <div id="App">
            <div id="logo-container">
                <img src={logo} id="logo" alt="logo" />
            </div>
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <input id="expression" className="input" onChange={updateName} autoComplete="off" name="input" type="text" />
                <button className="btn" onClick={greet}>Calc</button>
            </div>
            <div>
                <h1>Sine Wave Graph</h1>
                <Line data={Graph()} />
            </div>
        </div>
    )
}

export default App

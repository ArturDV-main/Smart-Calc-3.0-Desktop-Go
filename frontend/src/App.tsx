import { useState, useEffect } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { Greet } from "../wailsjs/go/main/App";
import { Graph } from "../wailsjs/go/main/App";
import { Line } from 'react-chartjs-2';
import { Chart, registerables } from 'chart.js';

Chart.register(...registerables);

interface ChartData {
    labels: number[];
    datasets: {
        label: string;
        data: number[];
        borderColor: string;
        backgroundColor: string;
        borderWidth: number;
    }[];
}

function App() {
    const [resultText, setResultText] = useState("Please enter your expression below 👇");
    const [resultGraph, setResultGraph] = useState<ChartData>({
        labels: [],
        datasets: [
            {
                label: 'expression',
                data: [],
                borderColor: 'rgba(75, 192, 192, 1)',
                backgroundColor: 'rgba(75, 192, 192, 0.2)',
                borderWidth: 1,
            },
        ],
    });
    const [expression, setExpression] = useState('');

    const updateExpression = (e: any) => setExpression(e.target.value);

    const updateResultText = (result: string) => setResultText(result);

    function greet() {
        Greet(expression, 0).then(updateResultText);
    }

    async function graph() {
        const xValues: number[] = [];
        const yValues: number[] = [];
        for (let x = -20; x <= 20; x += 0.01) {
            const roundedX = parseFloat(x.toFixed(7));
            xValues.push(roundedX);
            yValues.push(await Graph(expression, roundedX));
        }
        setResultGraph({
            labels: xValues,
            datasets: [
                {
                    label: expression,
                    data: yValues,
                    borderColor: 'rgba(75, 192, 192, 1)',
                    backgroundColor: 'rgba(75, 192, 192, 0.2)',
                    borderWidth: 1,
                },
            ],
        });
    }

    return (
        <div id="App">
            <div id="logo-container">
                <img src={logo} id="logo" alt="logo" />
            </div>
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <input id="expression" className="input" onChange={updateExpression} autoComplete="off" name="input" type="text" />
                <button className="btn" onClick={greet}>Calc</button>
            </div>
            <div>Num X <input id="expression" className="inputx" onChange={updateExpression} autoComplete="off" name="input" type="text" /></div>
            <div>
                <h2>Wave Graph</h2>
                <div id="input" className="input-box">
                    <button className="btn" onClick={graph}>Graph</button>
                </div>
                <Line data={resultGraph} />
            </div>
        </div>
    );
}

export default App;
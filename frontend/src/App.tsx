import { useState } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { Greet } from "../wailsjs/go/main/App";
import { Graph } from "../wailsjs/go/main/App";
import { Line } from 'react-chartjs-2';
import { useEffect } from 'react';
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

function GraphCreate() {
    const [data, setData] = useState<ChartData>({
        labels: [],
        datasets: [
            {
                label: 'sin(x)',
                data: [],
                borderColor: 'rgba(75, 192, 192, 1)',
                backgroundColor: 'rgba(75, 192, 192, 0.2)',
                borderWidth: 1,
            },
        ],
    });

    useEffect(() => {
        const generateSineData = () => {
            const xValues: number[] = [];
            const yValues: number[] = [];

            for (let x = 0; x <= 360; x += 1) {
                xValues.push(x);
                yValues.push(Math.sin((x * Math.PI) / 180));
            }

            setData({
                labels: xValues,
                datasets: [
                    {
                        label: 'sin(x)',
                        data: yValues,
                        borderColor: 'rgba(75, 192, 192, 1)',
                        backgroundColor: 'rgba(75, 192, 192, 0.2)',
                        borderWidth: 1,
                    },
                ],
            });
        };

        generateSineData();
    }, []);

    return (
        data
    );
}

function App() {
    const [data, setData] = useState<ChartData>({
        labels: [],
        datasets: [
            {
                label: 'sin(x)',
                data: [],
                borderColor: 'rgba(75, 192, 192, 1)',
                backgroundColor: 'rgba(75, 192, 192, 0.2)',
                borderWidth: 1,
            },
        ],
    });
    const [resultText, setResultText] = useState("Please enter your expression below 👇");
    const [resultGraph, setResultGraph] = useState(data);

    const [expression, setName] = useState('');
    const updateName = (e: any) => setName(e.target.value);

    const updateResultText = (result: string) => setResultText(result);
    const updateResultGraph = (result_graph: ChartData) => setResultGraph(result_graph);


    function greet() {
        Greet(expression).then(updateResultText);
    }

    function graph() {

        GraphCreate(expression).then(updateResultGraph);
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
                <h2>Sine Wave Graph</h2>
                <div id="result" className="result">{resultText}</div>
                <div id="input" className="input-box">
                    <input id="expression" className="input" onChange={updateName} autoComplete="off" name="input" type="text" />
                    <button className="btn" onClick={graph}>Calc</button>
                </div>
                <Line data={GraphCreate()}/>
            </div>
        </div>
    )
}

export default App

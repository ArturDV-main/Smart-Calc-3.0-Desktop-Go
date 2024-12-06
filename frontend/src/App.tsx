import './App.css'

import { useCallback, useRef, useState } from 'react'
import logo from './assets/images/logo-universal.png'
import { Graph, Greet } from "../wailsjs/go/main/App"
import { ChartConfiguration} from 'chart.js'
import { ChartGraph } from './components/graph'

// --------------------------------------------------------------------------------

export const App: React.FC = () => {
    const inpurRef = useRef<HTMLInputElement>(null)
    const xInputRef = useRef<HTMLInputElement>(null)
    const aInputRef = useRef<HTMLInputElement>(null)
    const bInputRef = useRef<HTMLInputElement>(null)
    const [graphData, setGraphData] = useState<ChartConfiguration>()

    const [resultText, setResultText] = useState("Please enter your expression below 👇")

    const handleCalcClick = useCallback(async () => {
        if (!inpurRef.current || !xInputRef.current) return

        const x = xInputRef.current.value ? +xInputRef.current.value : 0

        const result = await Greet(inpurRef.current.value, x)
        setResultText(result)
    }, [])

    const handleGraphClick = useCallback(async () => {
        if (!inpurRef.current) return

        setGraphData(undefined)

        const expression = inpurRef.current.value

        const rangeA = aInputRef.current ? +aInputRef.current.value : 0
        const rangeB = bInputRef.current ? +bInputRef.current.value : 0

        const diff = rangeB - rangeA

        const disc = []
        for (let x = rangeA; x <= rangeB; x += diff / 10000) {
            disc.push(x)
        }

        const xvalues = disc.map(x => +x.toFixed(7))
        const yvalues = await Promise.all(disc.map(async (x) => await Graph(expression, x)))

        setGraphData({
            type: 'line',
            data: {                
                labels: xvalues,
                datasets: [
                    {
                        label: expression,
                        data: yvalues,
                        borderColor: 'rgba(75, 192, 192, 1)',
                        backgroundColor: 'rgba(75, 192, 192, 0.2)',
                        borderWidth: 1
                    },
                ],
            }                    
        })
    }, [])

    return (
        <div id="App">
            <div id="logo-container">
                <img src={logo} id="logo" alt="logo" />
            </div>
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <input id="expression" className="input" ref={inpurRef} autoComplete="off" name="input" type="text" />
                <label>
                    Num x <input id="xval" className="inputx" ref={xInputRef} autoComplete="off" name="input" type="text" />
                </label>
                <button className="btn" onClick={handleCalcClick}>Calc</button>
            </div>
            <div>
                <h2>Wave Graph</h2>
                <div className='graph-input'>
                    <label>
                        <span>Range A</span>
                        <input id="xval" className="inputx" ref={aInputRef} autoComplete="off" name="input" type="text" />
                    </label>
                    <label>
                        <span>Range B</span>
                        <input id="xval" className="inputx" ref={bInputRef} autoComplete="off" name="input" type="text" />
                    </label>
                    <button className="btn" onClick={handleGraphClick}>Graph</button>
                </div>
                <ChartGraph data={graphData} />
            </div>
        </div>
    )
}
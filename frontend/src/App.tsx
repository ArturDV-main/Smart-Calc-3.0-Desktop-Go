import './App.css'

import React, { useCallback, useEffect, useRef, useState } from 'react'
import logo from './assets/images/logo-universal.png'
import { Greet, GraphicCalc, HistoryRead, HistoryClean } from "../wailsjs/go/main/App"
import { main } from '../wailsjs/go/models'
import { Graph } from './components/graph'

// --------------------------------------------------------------------------------

export const App: React.FC = () => {
    const inpurRef = useRef<HTMLInputElement>(null)
    const xInputRef = useRef<HTMLInputElement>(null)
    const aInputRef = useRef<HTMLInputElement>(null)
    const bInputRef = useRef<HTMLInputElement>(null)
    const rangeA = useRef<number>(0)
    const rangeB = useRef<number>(0)
    const [graphData, setGraphData] = useState<main.GraphData>()
    const [historyList, setHistoryList] = useState<string[]>([])
    const [listOpened, setListOpened] = useState(false)

    const [resultText, setResultText] = useState("Please enter your expression below")

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
        rangeA.current = aInputRef.current ? +aInputRef.current.value : 0
        rangeB.current = bInputRef.current ? +bInputRef.current.value : 0

        const data = await GraphicCalc(expression, rangeA.current, rangeB.current)


        setGraphData(data)

    }, [])

    const handleHistoryBtnClick = useCallback(async () => {
        if (listOpened) {
            setListOpened(false)
            return
        }
    }, [])

    useEffect(() => {
        if (listOpened)
            HistoryRead().then(list => setHistoryList(list))
        else
            setHistoryList([])

    }, [listOpened])

    const handleInput = useCallback((e: React.KeyboardEvent<HTMLInputElement>) => {
        const beforeDot = e.currentTarget.value.split(".")[0]
        const afterDot = e.currentTarget.value.split(".")[1] ?? 0
        const qwe = +beforeDot >= 1000000 || afterDot.length >= 7
        if (qwe && e.key !== "Backspace")
            e.preventDefault()
    }, [])

    return (
        <div id="app">
            <div id="logo-container">
                <img src={logo} id="logo" alt="logo" />
            </div>
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <label style={{ position: 'relative' }}>
                    <input id="expression" className="input" ref={inpurRef} autoComplete="off" name="input" type="text" />
                    <button className='list-button' onClick={() => setListOpened(!listOpened)}>
                        <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#1b2636">
                            <path d="M480-360 280-560h400L480-360Z" />
                        </svg>
                    </button>
                    {listOpened &&
                        <div className='history-list'>
                            <ul>
                                {historyList.map((item, index) =>
                                    <li
                                        key={'history-item' + index}
                                        onClick={() => {
                                            if (inpurRef.current)
                                                inpurRef.current.value = item

                                            setListOpened(false)
                                        }}
                                    >
                                        {item}
                                    </li>
                                )}
                            </ul>

                            <button
                                className='clear-history'
                                onClick={() => {
                                    HistoryClean()
                                    .finally(() => {
                                        setListOpened(false) 
                                    })
                                }}
                            >
                                Clear history
                            </button>
                        </div>}
                </label>
                <label>
                    Num x <input id="xval" className="inputx" ref={xInputRef} autoComplete="off" name="input" type="text" />
                </label>
                <button className="btn" onClick={handleCalcClick}>Calc =</button>
            </div>
            <div>
                <h2>Wave Graph</h2>
                <div className='graph-input'>
                    <label>
                        <span>Range A</span>
                        <input id="xval" className="inputx" ref={aInputRef} max="1000000" min="-1000000" type="number" onKeyDown={handleInput} />
                    </label>
                    <label>
                        <span>Range B</span>
                        <input id="xval" className="inputx" ref={bInputRef} max="1000000" min="-1000000" type="number" onKeyDown={handleInput} />
                    </label>
                    <button className="btn" onClick={handleGraphClick}>Graph</button>
                </div>
                <div className='graph-container'>
                    <Graph data={graphData} rangeA={rangeA.current} rangeB={rangeB.current} />
                </div>
            </div>
        </div>
    )
}
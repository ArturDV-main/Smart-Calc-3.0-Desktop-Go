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
    const a_yInputRef = useRef<HTMLInputElement>(null)
    const b_yInputRef = useRef<HTMLInputElement>(null)
    const rangeA = useRef<number>(0)
    const rangeB = useRef<number>(0)
    const rangeY_A = useRef<number>(0)
    const rangeY_B = useRef<number>(0)
    const [graphData, setGraphData] = useState<main.GraphData>()
    const [historyList, setHistoryList] = useState<string[]>([])
    const [listOpened, setListOpened] = useState(false)
    const [helpOpened, setHelpOpened] = useState(false)

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
        rangeY_A.current = a_yInputRef.current ? +a_yInputRef.current.value : 0
        rangeY_B.current = b_yInputRef.current ? +b_yInputRef.current.value : 0

        const data = await GraphicCalc(expression, rangeA.current, rangeB.current, rangeY_A.current, rangeY_B.current)

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
        <>
            <div id="app">
                <button className='help' onClick={() => setHelpOpened(true)}>
                    <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e8eaed">
                        <path
                            d="M478-240q21 0 35.5-14.5T528-290q0-21-14.5-35.5T478-340q-21 0-35.5 14.5T428-290q0 21 14.5 35.5T478-240Zm-36-154h74q0-33 7.5-52t42.5-52q26-26 
                            41-49.5t15-56.5q0-56-41-86t-97-30q-57 0-92.5 30T342-618l66 26q5-18 22.5-39t53.5-21q32 0 48 17.5t16 38.5q0 20-12 37.5T506-526q-44 39-54 59t-10 73Zm38 
                            314q-83 0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 54 85.5 127T880-480q0 83-31.5 
                            156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 0-227 93t-93 227q0 134 93 227t227 93Zm0-320Z"/>
                    </svg>
                </button>
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
                            <span> Range X from </span>
                            <input id="xval" className="inputx" ref={aInputRef} max="1000000" min="-1000000" type="number" onKeyDown={handleInput} defaultValue={-17}/>
                        </label>
                        <label>
                            <span> to </span>
                            <input id="xval" className="inputx" ref={bInputRef} max="1000000" min="-1000000" type="number" onKeyDown={handleInput} defaultValue={17}/>
                        </label>
                        <button className="btn" onClick={handleGraphClick}>Graph</button>
                        <label>
                            <span>Range Y from </span>
                            <input id="xval" className="inputx" ref={a_yInputRef} max="1000000" min="-1000000" type="number" onKeyDown={handleInput} defaultValue={15}/>
                        </label>
                        <label>
                            <span> to </span>
                            <input id="xval" className="inputx" ref={b_yInputRef} max="1000000" min="-1000000" type="number" onKeyDown={handleInput} defaultValue={-15}/>
                        </label>
                    </div>
                    <div className='graph-container'>
                        <Graph data={graphData} rangeA={rangeA.current} rangeB={rangeB.current} />
                    </div>
                </div>
            </div>

            {helpOpened &&
                <div className='help-container'>
                    <button className='help' onClick={() => setHelpOpened(false)}>
                        <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e8eaed">
                            <path
                                d="m336-280 144-144 144 144 56-56-144-144 144-144-56-56-144 144-144-144-56 56 144 144-144 144 56 56ZM480-80q-83 
                                0-156-31.5T197-197q-54-54-85.5-127T80-480q0-83 31.5-156T197-763q54-54 127-85.5T480-880q83 0 156 31.5T763-763q54 
                                54 85.5 127T880-480q0 83-31.5 156T763-197q-54 54-127 85.5T480-80Zm0-80q134 0 227-93t93-227q0-134-93-227t-227-93q-134 
                                0-227 93t-93 227q0 134 93 227t227 93Zm0-320Z" />
                        </svg>
                    </button>
                    <article>
                    <div>Умный калькулятор 3.0. </div>
                    <div>Необходимо ввести выражение в поле и кликнусь Calc = ,
                        результат будет выведен над полем ввода выражения.
                        Для построения графика, введите выражение в поле для расчета, введите оласти зачений и область определений для графика, используйте кнопку Graph.
                        Если была допущена ошибка или выражение не имеет решений, будет выведена ошибка, график так же не будет построен.
                        </div>
                    </article>
                </div>}
        </>

    )
}
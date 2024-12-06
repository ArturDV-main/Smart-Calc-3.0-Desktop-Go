import { Chart, ChartConfiguration, registerables } from "chart.js"
import { useEffect, useRef } from "react"

// --------------------------------------------------------------------------------

Chart.register(...registerables)

interface ChartGraphProps {
    data?: ChartConfiguration
}

export const ChartGraph: React.FC<ChartGraphProps> = ({ data }) => {
    const canvas = useRef<HTMLCanvasElement>(null)

    useEffect(() => {
        if (canvas.current && data)
            new Chart(canvas.current, data)
    }, [!data])

    if (!data) return null

    return (
        <canvas ref={canvas}></canvas>    
    )
    
}
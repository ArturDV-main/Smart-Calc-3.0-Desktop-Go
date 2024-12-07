import { useEffect, useRef } from "react"
import * as d3 from "d3"
import { main } from "../../wailsjs/go/models"

interface GraphProps {
    data?: main.GraphData
    rangeA: number
    rangeB: number
}

export const Graph: React.FC<GraphProps> = ({ data, rangeA, rangeB }) => {
    const svgRef = useRef<SVGSVGElement>(null)
    const width = 1000
    const height = 600

    useEffect(() => {
        if (!data) return

        const { points, maxY, minY } = data
        const margin = { top: 20, right: 20, bottom: 30, left: 40 }

        const xScale = d3.scaleLinear()
            .domain([rangeA, rangeB]) // диапазон x от 0 до 2π
            .range([margin.left, width - margin.right])

        const yScale = d3.scaleLinear()
            .domain([minY, maxY]) // диапазон y от -1 до 1
            .range([height - margin.bottom, margin.top])

        const line = d3.line()
            .x((d: any) => xScale(d.x))
            .y((d: any) => yScale(d.y))

        const svg = d3.select(svgRef.current)
        // .attr('width', width)
        // .attr('height', height)

        svg.selectAll('*').remove() // очищаем SVG перед отрисовкой

        svg.append('g')
            .append('path')
            .datum(points)
            .attr('fill', 'none')
            .attr('stroke', 'steelblue')
            .attr('stroke-width', 2)
            .attr('d', line as any)

        // Добавим оси
        svg.append('g')
            .attr('transform', `translate(0,${height - margin.bottom})`)
            .call(d3.axisBottom(xScale))

        svg.append('g')
            .attr('transform', `translate(${margin.left},0)`)
            .call(d3.axisLeft(yScale))

    }, [!data])

    if (!data) return null

    return <svg ref={svgRef} viewBox={`0 0 ${width} ${height}`} className="graph"></svg>
}

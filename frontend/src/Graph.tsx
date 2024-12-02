import { useState } from 'react';
import './Graph.css';

import React, { useEffect } from 'react';
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

function Graph() {
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

export default Graph
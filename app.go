package main

import (
	"context"
	"fmt"
	"smartcalc/pkg/calcadapter"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(expression string, num_x float64) string {
	got, err := calcadapter.Calculate(expression, num_x)
	if err != nil {
		return fmt.Sprintf("Calculate error:  %s", fmt.Sprintln(err.Error()))
	}
	return fmt.Sprintln("Got: ", got)
}

type GraphData struct {
	Points []calcadapter.Point `json:"points"`
	MaxY   float64             `json:"maxY"`
	MinY   float64             `json:"minY"`
}

func (a *App) GraphicCalc(expression string, range_a float64, range_b float64, value_range_a float64, value_range_b float64) GraphData {
	got, err := calcadapter.GraphicCalc(calcadapter.GraphSrc{Str_r: expression, Range_a: range_a, Range_b: range_b, Value_range_a: value_range_a, Value_range_b: value_range_b})
	var data GraphData
	if err != nil {
		data.Points = []calcadapter.Point{{X: 0, Y: 0}}
		return data
	}
	if len(got) > 0 {
		data.MaxY = got[0].Y
		data.MinY = got[0].Y
	}
	for _, v := range got {
		if v.Y > data.MaxY {
			data.MaxY = v.Y
		}
		if v.Y < data.MinY {
			data.MinY = v.Y
		}
	}
	data.Points = got
	return data
}

func (a *App) HistoryRead() ([]string, error) {
	return calcadapter.HistoryRead()
}

func (a *App) HistoryClean() {
	calcadapter.CleanHistory()
}

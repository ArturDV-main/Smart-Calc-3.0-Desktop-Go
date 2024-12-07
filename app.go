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
		return fmt.Sprintf("Calculate error:  %s", err.Error())
	}
	return fmt.Sprintln("Got: ", got)
}

func (a *App) Graph(expression string, num_x float64) float64 {
	got, err := calcadapter.Calculate(expression, num_x)
	if err != nil {
		return 0
	}
	return got
}

func (a *App) GraphicCalc(expression string, range_a float64, range_b float64) []float64 {
	got, err := calcadapter.GraphicCalc(expression, range_a, range_b)
	if err != nil {
		return nil
	}
	return got
}

func (a *App) HistoryRead() ([]string, error) {
	return calcadapter.HistoryRead()
}

func (a *App) HistoryClean() {
	calcadapter.CleanHistory()
}

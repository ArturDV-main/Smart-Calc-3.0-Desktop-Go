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
func (a *App) Greet(expression string) string {
	got, err := calcadapter.Calculate(expression)
	if err != nil {
		return fmt.Sprintf("Calculate error:  %s", err.Error())
	}
	return fmt.Sprintln("Got: ", got)
}

func (a *App) Graph(expression string) (float64, error) {
	got, err := calcadapter.Calculate(expression)
	if err != nil {
		return 0, err
	}
	return got, nil
}

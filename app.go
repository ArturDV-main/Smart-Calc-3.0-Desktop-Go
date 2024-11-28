package main

import (
	"context"
	"fmt"
	calcadapter "leftrana/smartcalc/pkg/calcadapter"
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
	if err != nil || got != 15 {
		return fmt.Sprintf("Calculate error:  %s, %s", err.Error(), expression)
	}

	return fmt.Sprintf("Hello %s, It's %f", expression, got)
}

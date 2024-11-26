package main

import (
	"fmt"
	calcadapter "leftrana/smartcalc/pkg/calcadapter"
)

func main() {
	str := "5+5*2"
	got, err := calcadapter.Calculate(str)
	if err != nil || got != 15 {
		fmt.Printf("5+5*2 = %f; want 15 ", got)
	}
}

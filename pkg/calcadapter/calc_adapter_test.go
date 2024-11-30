package calcadapter_test

import (
	calcadapter "leftrana/smartcalc/pkg/calcadapter"
	"testing"
)

func TestAbs(t *testing.T) {
	str := "5+5*2"
	got, err := calcadapter.Calculate(str)
	if err != nil || got != 15 {
		t.Errorf("5+5*2 = %f; want 15 ", got)
	}
	str = "-5+5*2"
	got, err = calcadapter.Calculate(str)
	if err != nil || got != 5 {
		t.Errorf("-5+5*2 = %f; want -10 ", got)
	}
	str = "5-5*2"
	got, err = calcadapter.Calculate(str)
	if err != nil || got != -5 {
		t.Errorf("5-5*2 = %f; want 0 ", got)
	}
}

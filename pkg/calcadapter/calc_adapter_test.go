package calcadapter_test

import (
	"log"
	"os"
	calcadapter "smartcalc/pkg/calcadapter"
	"testing"
)

func TestAbs(t *testing.T) {
	str := "5+5*2"
	got, err := calcadapter.Calculate(str, 0)
	if err != nil || got != 15 {
		t.Errorf("5+5*2 = %f; want 15 ", got)
	}
	str = "-5+5*2"
	got, err = calcadapter.Calculate(str, 0)
	if err != nil || got != 5 {
		t.Errorf("-5+5*2 = %f; want -10 ", got)
	}
	str = "5-5*2"
	got, err = calcadapter.Calculate(str, 0)
	if err != nil || got != -5 {
		t.Errorf("5-5*2 = %f; want 0 ", got)
	}
	got_graph, err := calcadapter.GraphicCalc("sin ( x )", -10, 10)
	if err != nil {
		t.Errorf("err = %v; want nil", err)
	}
	if len(got_graph) == 0 {
		t.Errorf("got = %d; want not 0", len(got_graph))
	}
	str = "atan(x)"
	for i := 1; i < 3000; i++ {
		_, err = calcadapter.Calculate(str, 2)
		if err != nil {
			t.Errorf("atan calc err: %f ", err)
		}
	}
	str = "tan(x)"
	for i := 1; i < 10000; i++ {
		r, err := calcadapter.Calculate(str, 2)
		if err != nil {
			t.Errorf("atan calc err: %f ", err)
		}
		if r+2.185040 > 0.00001 {
			t.Errorf("atan calc want: -2.185040, got: %f ", r)
		}
	}
	str = "sin(x)"
	for i := 1; i < 10000; i++ {
		r, err := calcadapter.Calculate(str, 3.1415)
		if err != nil {
			t.Errorf("atan calc err: %f ", err)
		}
		if r-0.000093 > 0.000001 {
			t.Errorf("atan calc want: -0.000093, got: %f ", r)
		}
	}
	str = "sin(x)"
	for i := 1; i < 10000; i++ {
		r, err := calcadapter.Calculate(str, 3.1416)
		if err != nil {
			t.Errorf("atan calc err: %f ", err)
		}
		if r+0.000007 > 0.000001 {
			t.Errorf("atan calc want: -0.000093, got: %f ", r)
		}
	}
	str = "atan(100)"
	for i := 1; i < 10000; i++ {
		r, err := calcadapter.Calculate(str, 0.0)
		if err != nil {
			t.Errorf("atan calc err: %f ", err)
		}
		if r-1.560797 > 0.000001 {
			t.Errorf("atan calc want: 1.560797, got: %f ", r)
		}
	}
	got_graph, err = calcadapter.GraphicCalc("atan(x)", -10, 10)
	if err != nil {
		t.Errorf("err = %v; want != nil", err)
	}
	if got_graph[0].X != -10 {
		t.Errorf("got = %f; want != -10", got_graph[0].X)
	}

	_, err = calcadapter.Calculate("kavabanga", -10)
	if err == nil {
		t.Errorf("err = %v; want != nil", err)
	}
}

func TestHistory(t *testing.T) {
	err := os.Remove(calcadapter.History)
	if err != nil && err.Error() != "remove history.txt: no such file or directory" {
		log.Println("unable to delete file: ", err)
		return
	}
	err = calcadapter.HistoryWrite("5+5*2")
	if err != nil {
		t.Errorf("err = %v; want nil", err)
	}

	err = calcadapter.HistoryWrite("10+10*8")
	if err != nil {
		t.Errorf("err = %v; want nil", err)
	}
	str, err := calcadapter.HistoryRead()

	if len(str) != 2 {
		t.Errorf("got = %d; want 2 ", len(str))
	}
	if err != nil || str[0] != "5+5*2" {
		t.Errorf("got = %s; want 5+5*2", str[0])
	}
	calcadapter.CleanHistory()
	_, err = calcadapter.HistoryRead()
	if err == nil {
		t.Errorf("err = %v; want not nil", err)
	}
}

func TestUno(t *testing.T) {
	for i := 0; i < 1000; i++ {
		_, err := calcadapter.Calculate("kavabanga", -10)
		if err == nil {
			t.Errorf("err = %v; want != nil", err)
		}
	}
}

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
}

func TestHistory(t *testing.T) {
	err := os.Remove(calcadapter.History)
	if err != nil && err.Error() != "remove history.txt: no such file or directory" {
		log.Println("Не удалось удалить файл:", err)
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

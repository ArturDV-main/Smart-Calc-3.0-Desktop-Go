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
}

func TestHistory(t *testing.T) {
	err := os.Remove("history.txt")
	if err != nil {
		log.Println("Не удалось удалить файл:", err)
		return
	}
	err = calcadapter.HistoryWrite("5+5*2")
	if err != nil {
		t.Errorf("err = %v; want nil", err)
	}

	err = calcadapter.HistoryWrite("10+10*2")
	if err != nil {
		t.Errorf("err = %v; want nil", err)
	}
	// str, err = calcadapter.HistoryRead()
	// if err != nil || got != 5 {
	// 	t.Errorf("-5+5*2 = %f; want -10 ", got)
	// }
	// str = "5-5*2"
	// got, err = calcadapter.Calculate(str, 0)
	// if err != nil || got != -5 {
	// 	t.Errorf("5-5*2 = %f; want 0 ", got)
	// }
}

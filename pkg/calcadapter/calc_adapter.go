package calcadapter

/*
   #cgo LDFLAGS: -L.  -lsmart_calc
   #include "../../cpp/src/s21_calc_controller.h"
*/
import "C"
import (
	"errors"
	"log"
)

func Calculate(str string) (float64, error) {
	cstr := C.CString(str)
	c := C.StartCalc(cstr, 0)
	if c.err == 1 {
		log.Println("error")
		return 0, errors.New("error")
	}

	return float64(c.result), nil
}

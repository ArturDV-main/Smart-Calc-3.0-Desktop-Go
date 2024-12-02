package calcadapter

/*
   #cgo CXXFLAGS: -std=c++17
   #cgo LDFLAGS: -L.  -lsmart_calc
   #include "../../cpp/src/s21_calc_controller.h"
*/
import "C"
import (
	"errors"
)

type Resp struct {
	Result float64
	Err    error
}

func Calculate(str string) (float64, error) {
	cstr := C.CString(str)
	c := C.StartCalc(cstr, 0)
	if c.err == 1 {
		return 0.0, errors.New("error")
	}
	return float64(c.result), nil
}

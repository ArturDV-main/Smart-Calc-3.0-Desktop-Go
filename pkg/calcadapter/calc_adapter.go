package calcadapter

/*
   #cgo CXXFLAGS: -std=c++17
   #cgo LDFLAGS: -lsmart_calc
   #include "../../cpp/src/s21_calc_controller.h"
*/
import "C"
import (
	"errors"
	"strings"
)

type Resp struct {
	Result float64
	Err    error
}

func Calculate(str string, num_x float64) (float64, error) {
	cstr := C.CString(replaceMathFunctions(str))
	num := C.double(num_x)
	c := C.StartCalc(cstr, num)
	if c.err == 1 {
		return 0.0, errors.New("error")
	}
	return float64(c.result), nil
}

type TrigonCode rune

const (
	COS  TrigonCode = '@'
	SIN  TrigonCode = 'A'
	TAN  TrigonCode = 'B'
	ACOS TrigonCode = 'C'
	ASIN TrigonCode = 'D'
	ATAN TrigonCode = 'E'
	SQRT TrigonCode = 'F'
	LN   TrigonCode = 'G'
	LOG  TrigonCode = 'H'
)

func replaceMathFunctions(input string) string {
	replacements := map[string]TrigonCode{
		"cos":  COS,
		"sin":  SIN,
		"tan":  TAN,
		"acos": ACOS,
		"asin": ASIN,
		"atan": ATAN,
		"sqrt": SQRT,
		"ln":   LN,
		"log":  LOG,
	}

	for funcName, code := range replacements {
		input = strings.ReplaceAll(input, funcName, string(code))
	}

	return input
}

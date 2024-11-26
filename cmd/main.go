package main

/*
   #cgo LDFLAGS: -L.  -lsmart_calc
   #include "../cpp/src/s21_calc_controller.h"
*/
import "C"
import "fmt"

func main() {
	str := "5+5*5"
	cstr := C.CString(str)
	c := C.StartCalc(cstr, 0)
	if c.err == 1 {
		fmt.Println("error")
	} else {
		fmt.Println("hello", c.result)
	}
}

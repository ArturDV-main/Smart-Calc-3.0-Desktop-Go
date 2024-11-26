package main

/*
   #cgo LDFLAGS: -L.  -lsmart_calc
   #include "./s21_calc_controller.h"
*/
import "C"
import "fmt"

func main() {
	str := "5+5"
	cstr := C.CString(str)
	c := C.StartCalc(cstr, 0)
	fmt.Println("hello", c)
}

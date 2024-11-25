package main

/*
   #cgo CXXFLAGS: -std=c++17
   #cgo LDFLAGS: -L./libs  -lsmart_calc -lstdc++
   #include "c/src/s21_smartcalc.h"
*/
import "C"
import "fmt"

func main() {
	str := "5+5"
	cstr := C.CString(str)
	c := C.Start_calc(cstr, 0)
	fmt.Println("hello", c)
}

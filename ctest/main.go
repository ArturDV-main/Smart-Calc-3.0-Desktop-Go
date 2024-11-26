package main

/*
#cgo linux LDFLAGS: -L. -ladd
#include "./add.h"
*/
import "C"
import (
	"fmt"
)

func main() {
	a := 10
	b := 20
	sum := C.add(C.int(a), C.int(b))
	fmt.Printf("Сумма %d и %d равна %d\n", a, b, sum)
}

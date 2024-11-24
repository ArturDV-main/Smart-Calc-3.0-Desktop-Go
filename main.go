package main

/*
   #cgo LDFLAGS: -L. -lmylib
   extern void hallo();
*/
import "C"

func main() {
	C.hallo()
}

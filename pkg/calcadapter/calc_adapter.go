package calcadapter

/*
   #cgo CXXFLAGS: -std=c++17
   #cgo LDFLAGS: -L.  -lsmart_calc
   #include "../../cpp/src/s21_calc_controller.h"
*/
import "C"
import (
	"errors"
	"log"
	"sync"
)

type Resp struct {
	Result float64
	Err    error
}

func Calculate(str string) (float64, error) {
	var r Resp
	var wg sync.WaitGroup
	wg.Add(1)
	go calculating(&wg, str, &r)
	wg.Wait()
	if r.Err != nil {
		log.Println("error")
		return 0, r.Err
	}
	return r.Result, nil
}

func calculating(wg *sync.WaitGroup, str string, res *Resp) {
	if wg == nil || res == nil {
		log.Println("wg nil", ">")

		return
	}
	defer wg.Done()

	cstr := C.CString(str)
	c := C.StartCalc(cstr, 0)
	if c.err == 1 {
		log.Println("error")
		res.Err = errors.New("error")
	}
	res.Result = float64(c.result)
}

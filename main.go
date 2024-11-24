package main

/*
   #cgo CXXFLAGS: -std=c++17
   #cgo LDFLAGS: -L./lib  -lsmart_calc
   #include "cpp/controllers/s21_calc_controller.h"
   s21::CalcController controller;
   void StartCalc(const std::string& a, const double x) {
      controller.StartCalc(a, x);
   }
   double GetResult() const noexcept { return controller.GetResult(); }
*/
import "C"

func main() {
	C.StartCalc("2+2", 0)
	a := C.GetResult()
}

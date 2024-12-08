#include "./s21_calc_controller.h"
#include "./s21_calc_model.h"
#include <iostream>

#ifdef __cplusplus
extern "C" {
#endif

struct Resp StartCalc(const char* a, const double x) {
  s21::CalcModel calc_model;
  std::string s(a);
  struct Resp result;
  result.err = 0;
  result.errors = NULL;
  result.result = 0.0;
  try {
    calc_model.StartCalc(s, x);
  } catch (const std::exception& e) {
    result.err = 1;
    result.errors = e.what();
    return result;
  }
  
  result.result = calc_model.GetData();

  return result;
}

#ifdef __cplusplus
}
#endif
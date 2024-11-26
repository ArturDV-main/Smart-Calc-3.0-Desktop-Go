#include "./s21_calc_controller.h"
#include "./s21_calc_model.h"

#ifdef __cplusplus
extern "C" {
#endif

double StartCalc(const char* a, const double x) {

s21::CalcModel calc_model;
  std::string s(a);
  double result = 0;
  try {
    calc_model.StartCalc(s, x);
  } catch (const std::exception& e) {
    return result;
  }
  result = calc_model.GetData();
  return result;
}

#ifdef __cplusplus
}
#endif
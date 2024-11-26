#include "./s21_calc_controller.h"
#include "./s21_calc_model.h"

#ifdef __cplusplus
extern "C" {
#endif

Response StartCalc(const char* a, const double x) {

s21::CalcModel calc_model;
  std::string s(a);
  Response result;
  try {
    calc_model.StartCalc(s, x);
  } catch (const std::exception& e) {
    result.err = 1;
    
    return result;
  }
  result.result = calc_model.GetData();

  return result;
}

#ifdef __cplusplus
}
#endif
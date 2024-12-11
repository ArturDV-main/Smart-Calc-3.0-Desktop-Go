#include "./s21_calc_controller.h"
#include <stdlib.h>
#include <cstring>
#include "./s21_calc_model.h"

#ifdef __cplusplus
extern "C" {
#endif

struct Resp StartCalc(const char* a, const double x) {
  s21::CalcModel calc_model;
  std::string s(a);
  struct Resp result;
  result.errors = nullptr;
  result.result = 0.0;
  std::string r = calc_model.Calculating(s, x);
  result.err = 0;
  if (r != "") {
    char* copy = (char*)malloc((r.length() + 1) * sizeof(char));
    if (copy == NULL) {
      return result;
    }
    strcpy(copy, r.c_str());
    result.err = 1;
    result.errors = copy;
    return result;
  }
  result.err = 0;
  result.result = calc_model.GetData();
  return result;
}

#ifdef __cplusplus
}
#endif
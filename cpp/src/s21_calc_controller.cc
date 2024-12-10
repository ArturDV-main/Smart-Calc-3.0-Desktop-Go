#include "./s21_calc_controller.h"
#include "./s21_calc_model.h"
#include <stdlib.h>
#include <string>
#include <stdio.h>

#ifdef __cplusplus
extern "C" {
#endif

Response StartCalc(const char* a, const double x) {
  s21::CalcModel calc_model;
  std::string s(a);
  Response result;
  result.err = 0;
  result.errors = nullptr;
  result.result = 0.0;
  std::string r = calc_model.Calculating(s, x);

  if (r != "") {
    result.err = 1;
    std::string err = result.errors;
    size_t length = err.length() + 1;
    char * destination = (char *)malloc(length);
    if(destination == nullptr) {
      result.errors = nullptr;
       return result;
    }
    strcpy();
    return result;
  }

  result.result = calc_model.GetData();
  return result;
}

#ifdef __cplusplus
}
#endif
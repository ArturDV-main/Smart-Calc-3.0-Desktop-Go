#ifndef SRC_S21_CALC_CONTROLLER_H_
#define SRC_S21_CALC_CONTROLLER_H_

#ifdef __cplusplus
extern "C" {
#endif


struct Resp {
  int err;
  const char* errors;
  double result;
};

struct Resp StartCalc(const char* a, const double x);

#ifdef __cplusplus
}
#endif

#endif  //  SRC_S21_CALC_CONTROLLER_H_

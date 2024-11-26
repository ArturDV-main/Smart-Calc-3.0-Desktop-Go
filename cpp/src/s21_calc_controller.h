#ifndef SRC_S21_CALC_CONTROLLER_H_
#define SRC_S21_CALC_CONTROLLER_H_

#ifdef __cplusplus
extern "C" {
#endif

  typedef struct {
    int err;
    const char* errors;
    double result;
  } Response;
  Response StartCalc(const char* a, const double x);

#ifdef __cplusplus
}
#endif

#endif  //  SRC_S21_CALC_CONTROLLER_H_

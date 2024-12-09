#ifndef SRC_S21_CALC_CONTROLLER_H_
#define SRC_S21_CALC_CONTROLLER_H_

  struct Resp{
    int err;
    const char* errors;
    double result;
  };

  struct Resp StartCalc(const char* source, const double num_x);

#endif  //  SRC_S21_CALC_CONTROLLER_H_

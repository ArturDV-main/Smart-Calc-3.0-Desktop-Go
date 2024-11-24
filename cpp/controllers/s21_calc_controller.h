#ifndef SRC_S21_CALC_CONTROLLER_H_
#define SRC_S21_CALC_CONTROLLER_H_

#include <stdexcept>

#include "../models/s21_calc_model.h"

#ifdef __cplusplus
extern "C" {
#endif

namespace s21 {

class CalcController {
 public:
  CalcController(s21::CalcModel* m = new(s21::CalcModel)) : calc_model(m) {};
  void Reset() noexcept { calc_model->Reset(); }
  void StartCalc(const std::string& a, const double x);
  void CalcCredit(std::array<double, 3> data) { calc_model->CalcCredit(data); }
  void DifferenCalc(std::array<double, 3> data) {
    calc_model->DifferenCalc(data);
  }
  double GetResult() const noexcept { return calc_model->GetData(); }
  std::array<double, 3> GetCredit() const noexcept {
    return calc_model->GetCredit();
  }
  std::vector<double> GetDifferent() const noexcept {
    return calc_model->GetDifferent();
  }

 private:
  s21::CalcModel* calc_model;
};

}  // namespace s21

#ifdef __cplusplus
}
#endif

#endif  //  SRC_S21_CALC_CONTROLLER_H_

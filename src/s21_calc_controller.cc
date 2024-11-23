#include "s21_calc_controller.h"

namespace s21 {

void s21::CalcController::StartCalc(const std::string& a, const double x) {
  try {
    calc_model->StartCalc(a, x);
  } catch (const std::exception& e) {
    throw std::runtime_error(e.what());
  }
}

}  // namespace s21

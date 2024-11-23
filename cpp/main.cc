#include <iostream>

#include "./controllers/s21_calc_controller.h"
#include "./models/s21_calc_model.h"

int main(int argc, char *argv[]) {
  s21::CalcModel calc_model;
  s21::CalcController calc_controller(&calc_model);
  return 0;
}

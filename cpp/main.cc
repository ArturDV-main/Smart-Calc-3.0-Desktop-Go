#include <QApplication>
#include <iostream>

#include "s21_calc_controller.h"
#include "s21_calc_model.h"
#include "s21_view_qt/s21_view_qt.h"

int main(int argc, char *argv[]) {
  s21::CalcModel calc_model;
  s21::CalcController calc_controller(&calc_model);
  QApplication a(argc, argv);
  MainWindow w(nullptr, &calc_controller);
  w.show();
  return a.exec();
}

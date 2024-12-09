#include <iostream>

#include "../src/s21_calc_model.h"
#include "gtest/gtest.h"

TEST(CalcTest, GetDataCalc) {
  s21::CalcModel a;
  EXPECT_EQ(a.GetData(), 0);
}

TEST(CalcTest, SimpleCalc) {
  s21::CalcModel a;
  a.StartCalc("-2+2", 0);
  EXPECT_EQ(a.GetData(), 0);
  a.StartCalc("0.2/2", 0);
  EXPECT_EQ(a.GetData(), 0.1);
  a.StartCalc("@(0)+A(0)*B(0)", 0);
  EXPECT_EQ(a.GetData(), 1);
  a.StartCalc("C(1)-B(1)+E(1)", 0);
  EXPECT_DOUBLE_EQ(a.GetData(), -0.77200956125745379);
  a.StartCalc("2^2+x", 0);
  EXPECT_DOUBLE_EQ(a.GetData(), 4);
  a.StartCalc("H100+G2.7182818284", 0);
  EXPECT_DOUBLE_EQ(a.GetData(), 2.9999999999782783);
}

TEST(CalcTest, TrigoneCalc) {
  s21::CalcModel a;
  a.StartCalc("E(x)", 2);
  EXPECT_DOUBLE_EQ(a.GetData(), 1.1071487177940904);
}

TEST(CalcTest, ResetCalc) {
  s21::CalcModel a;
  a.StartCalc("2+2", 0);
  EXPECT_EQ(a.GetData(), 4);
  a.Reset();
  EXPECT_EQ(a.GetData(), 0);
}

TEST(CalcTest, Credit) {
  s21::CalcModel a;
  std::array<double, 3> data = {10.1, 10.1, 10.1};
  a.CalcCredit(data);
  std::array<double, 3> result = a.GetCredit();
  EXPECT_DOUBLE_EQ(result.at(0), 10.577792933162558);
}

TEST(CalcTest, Credit2) {
  s21::CalcModel a;
  std::array<double, 3> data = {10.1, 10.1, 10.1};
  a.DifferenCalc(data);
  std::vector<double> result = a.GetDifferent();
  EXPECT_DOUBLE_EQ(result.at(0), 11.481230410958904);
}

TEST(CalcTest, Exeptions) {
  s21::CalcModel a;
  EXPECT_ANY_THROW(a.StartCalc("22++R", 0););
  // std::array<double, 3> data = {-10.1, 10.1, 10.1};
  // EXPECT_ANY_THROW(a.DifferenCalc(data););
  // EXPECT_ANY_THROW(a.CalcCredit(data););
  // data[0] = 10.1;
  // data[1] = -10.1;
  // EXPECT_ANY_THROW(a.DifferenCalc(data););
  // EXPECT_ANY_THROW(a.CalcCredit(data););
}

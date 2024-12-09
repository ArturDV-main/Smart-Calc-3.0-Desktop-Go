// Copyright [2022] <Smithjan>
#include <check.h>
#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "./s21_smartcalc.h"

START_TEST(s21_smartcalc1) {
  char *string = "10*10/10+10-10+(-10)";
  double value1 = 0;
  double value2 = start_calc(string, 0);
  ck_assert_double_eq(value1, value2);
}
END_TEST

START_TEST(s21_smartcalc2) {
  char string[20] = "10^2M30";
  double value1 = fmod(pow(10.0, 2), 30.0);
  double value2 = start_calc(string, 0);
  ck_assert_double_eq(value1, value2);
}
END_TEST

START_TEST(s21_smartcalc3) {
  char string[30] = "@(1)+A(1)+B(1/2)";
  double value1 = 1.9280758;
  double value2 = start_calc(string, 0);
  int check = value1 - value2;
  ck_assert_int_eq(0, check);
}
END_TEST

START_TEST(s21_smartcalc4) {
  char string[30] = "D(1/2)+C(0)+E(1)";
  double value1 = 2.8797933;
  double value2 = start_calc(string, 0);
  int check = value1 - value2;
  ck_assert_int_eq(0, check);
}
END_TEST

START_TEST(s21_smartcalc5) {
  char string[30] = "F(4)+H(10)+G(10)";
  double value1 = 5.3025851;
  double value2 = start_calc(string, 0);
  int check = value1 - value2;
  ck_assert_int_eq(0, check);
}
END_TEST

START_TEST(s21_smartcalc6) {
  char string[50] = "16-(7+38*(6/2))*((3*6)-(9*(4+9-99)))/(96*9)";
  double value1 = -94.9166667;
  double value2 = start_calc(string, 0);
  int check = value1 - value2;
  ck_assert_int_eq(0, check);
}
END_TEST
START_TEST(s21_smartcalc7) {
  char string[100] = "F(4)+9.235+23*A(@(10))-564-F(25)+28.6*35";
  if (validator(string)) {
    double value1 = 426.122;
    double value2 = start_calc(string, 0);
    int check = value1 - value2;
    ck_assert_int_eq(0, check);
  }
}
END_TEST
START_TEST(s21_smartcalc8) {
  char string[100] = "F(4)+9.235+23^A(@(10^2))-564-F(25)+28.6*35+2^8";
  if (validator(string)) {
    double value1 = 710.050;
    double value2 = start_calc(string, 0);
    int check = value1 - value2;
    ck_assert_int_eq(0, check);
  }
}
END_TEST

START_TEST(s21_smartcalc9) {
  char *string = "2+3+(3+5)*2+3/(1-4)^3*(3^((2+1)+1))/(((1)))+H(2)";
  if (validator(string)) {
    double value1 = 12.6931;
    double value2 = start_calc(string, 0);
    int check = value1 - value2;
    ck_assert_int_eq(0, check);
  }
}
END_TEST
START_TEST(s21_smartcalc10) {
  char *string = "5+(1+2)*4*@(2*7.5-2)+A(@(2*5))-F(2^5)+D(0.5)+B(50)";
  if (validator(string)) {
    double value1 = 9.74018;
    double value2 = start_calc(string, 0);
    int check = value1 - value2;
    ck_assert_int_eq(0, check);
  }
}
END_TEST

START_TEST(s21_smartcalc11) {
  char *string = "5+(1+2)*4*@(2*7.5-2)+A(@(2*5))-F(2M5)+D(0.5)+B(50)";
  if (validator(string)) {
    double value1 = 13.9828;
    double value2 = start_calc(string, 0);
    int check = value1 - value2;
    ck_assert_int_eq(0, check);
  }
}
END_TEST

int main(void) {
  Suite *s1 = suite_create("Core");
  TCase *tc1_1 = tcase_create("Core");
  SRunner *sr = srunner_create(s1);
  int nf;

  suite_add_tcase(s1, tc1_1);
  tcase_add_test(tc1_1, s21_smartcalc1);
  tcase_add_test(tc1_1, s21_smartcalc2);
  tcase_add_test(tc1_1, s21_smartcalc3);
  tcase_add_test(tc1_1, s21_smartcalc4);
  tcase_add_test(tc1_1, s21_smartcalc5);
  tcase_add_test(tc1_1, s21_smartcalc6);
  tcase_add_test(tc1_1, s21_smartcalc7);
  tcase_add_test(tc1_1, s21_smartcalc8);
  tcase_add_test(tc1_1, s21_smartcalc9);
  tcase_add_test(tc1_1, s21_smartcalc10);
  tcase_add_test(tc1_1, s21_smartcalc11);

  srunner_run_all(sr, CK_ENV);
  nf = srunner_ntests_failed(sr);
  srunner_free(sr);

  return nf == 0 ? 0 : 1;
}

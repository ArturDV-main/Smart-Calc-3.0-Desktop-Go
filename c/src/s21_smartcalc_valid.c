#include "s21_smartcalc.h"

int in_line_start(const char *str, int str_len) {
  int in_start = FALSE;
  if (str[0] == '0' && str_len == 1) in_start = TRUE;
  return in_start;
}

int is_simp_oper(char oper) {
  int it_oper = FALSE;
  int oper_string = position_counter(oper);
  if (oper_string > 0 && oper_string < 7) it_oper = TRUE;
  return it_oper;
}

int char_counter(const char *str_line, char res) {
  int counter = 0, i = 0;
  while (str_line[i]) {
    if (str_line[i] == res) counter++;
    i++;
  }
  return counter;
}

int is_nums(char num) {
  int is_num = FALSE;
  if (('0' <= num && num <= '9') || num == '.') is_num = TRUE;
  return is_num;
}

int last_is(char res) {
  int is_number = FALSE;
  if (is_nums(res) || res == '.' || res == ')' || res == 'X') is_number = TRUE;
  return is_number;
}

int smart_bracket(const char *str_line) {
  int closed_bracket = FALSE;
  int len = strlen(str_line);
  int open_bracket_val = char_counter(str_line, '(');
  int close_bracket_val = char_counter(str_line, ')');
  int last_sym = last_is(str_line[len - 1]);

  if ((open_bracket_val > close_bracket_val) && last_sym) {
    closed_bracket = TRUE;
  } else if (in_line_start(str_line, len) || is_simp_oper(str_line[len - 1]) ||
             str_line[len - 1] == '(') {
    closed_bracket = FALSE;
  } else {
    closed_bracket = ERROR;
  }
  return closed_bracket;
}

int valid_simp_oper(const char *str_line) {
  int valid_oper = FALSE;
  int len = strlen(str_line);
  char lastch = str_line[len - 1];
  int last = last_is(lastch);
  if (last && !in_line_start(str_line, len)) valid_oper = TRUE;
  return valid_oper;
}

int valid_func(const char *str_line) {
  int validfunc = FALSE;
  int len = strlen(str_line);
  if (is_simp_oper(str_line[len - 1]) || str_line[len - 1] == '(' ||
      in_line_start(str_line, len))
    validfunc = TRUE;
  return validfunc;
}

int valid_nums(const char *str_line) {
  int validfunc = FALSE;
  int len = strlen(str_line);
  char lastchar = str_line[len - 1];
  if ((is_simp_oper(lastchar) || lastchar == '(' ||
       in_line_start(str_line, len) || is_nums(lastchar)) &&
      lastchar != 'X')
    validfunc = TRUE;
  return validfunc;
}

int valid_equals(const char *str_line) {
  int validequals = FALSE;
  int len = strlen(str_line);
  char lastchar = str_line[len - 1];
  if (lastchar == ')' || is_nums(lastchar) || lastchar == 'X')
    validequals = TRUE;
  return validequals;
}

int valid_unar(const char *str_line) {
  int validunar = FALSE;
  int len = strlen(str_line);
  char lastchar = str_line[len - 1];
  if (lastchar == '(' || in_line_start(str_line, len)) validunar = TRUE;
  return validunar;
}

int valid_dot(const char *str_line) {
  int validdot = TRUE;
  int i = 0;
  int len = strlen(str_line);
  while (is_nums(str_line[len - 1 - i])) {
    if (str_line[len - 1 - i] == '.') {
      validdot = FALSE;
      break;
    }
    i++;
  }
  return validdot;
}

int valid_dot_line(const char *str_line) {
  int validdot = TRUE;
  int i = 0;
  int len = strlen(str_line);
  while (is_nums(str_line[len - 1 - i])) {
    if (str_line[len - 1 - i] == '.') {
      validdot++;
    }
    i++;
  }
  if (validdot > 2) validdot = FALSE;
  return validdot;
}

int valid_input_line(double maxval, double minval, const char *str_line) {
  int validline = TRUE;
  int len = strlen(str_line);
  char lastchar = str_line[len - 1];
  if (str_line[1] == '\0' && lastchar == '-') {
    validline = TRUE;
  } else if (lastchar == '.') {
    validline = valid_dot_line(str_line);
  } else if (validline) {
    validline = is_nums(lastchar);
  }
  if (validline) {
    double test = atof(str_line);
    if (minval > test || test > maxval) validline = FALSE;
  }
  return validline;
}

int super_valid(double maxval, double minval, const char *str_line) {
  int validline = TRUE;
  //    int len = strlen(str_line);
  int i = 0;
  if (str_line[0] == '-') i = 1;
  while (str_line[i]) {
    if (is_nums(str_line[i]) != 1) {
      validline = FALSE;
      break;
    }
    i++;
  }
  if (validline) {
    int validdot = valid_dot_line(str_line);
    if (!validdot) validline = FALSE;
  }
  if (validline) {
    double test = atof(str_line);
    if (minval > test || test > maxval) validline = FALSE;
  }
  return validline;
}

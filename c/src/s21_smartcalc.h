#ifndef SRC_S21_SMARTCALC_H_
#define SRC_S21_SMARTCALC_H_

enum truefalse {
  FALSE,
  TRUE,
  ERROR,
  COS = '@',
  SIN = 'A',
  TAN = 'B',
  ACOS,
  ASIN = 'D',
  ATAN = 'E',
  SQRT,
  LN,
  LOG = 'H'
};

typedef struct Node_stack {
  double val_dub;
  int prio;
  struct Node_stack *next;
} stack_type;

int validator(const char *str);
double pop_val(stack_type **stack);
stack_type *push_sta(stack_type *plist, double val_dub, int prio);
int position_counter(char src_string);
int prio_check(char src_string);
stack_type parser_uno(const char *calculation_src, int *position, double X);
void destroy_node(stack_type *stack1);
stack_type *del_point(stack_type *stack1);
double simple_math(double second_num, double first_num, char operation);
double trigon_calc(double x, char operation);
double math_operations(stack_type **num_sta, stack_type **oper_sta);
int unar_check(char check, const char *oper_st, int position);
double calc(const char *calculation_src, double X);
double Start_calc(const char * src, double X);

// Validation
int last_is(char res);
int char_counter(const char * str_line, char res);
int in_line_start(const char * str, int str_len);
int is_simp_oper(char oper);
int is_nums(char num);
int smart_bracket(const char * str_line);
int valid_simp_oper(const char * str_line);
int valid_func(const char * str_line);
int valid_nums(const char * str_line);
int valid_equals(const char * str_line);
int valid_unar(const char * str_line);
int valid_dot(const char * str_line);
int valid_input_line(double maxval, double minval, const char * str_line);
int valid_dot_line(const char * str_line);
int super_valid(double maxval, double minval, const char * str_line);

#endif  // SRC_S21_SMARTCALC_H_

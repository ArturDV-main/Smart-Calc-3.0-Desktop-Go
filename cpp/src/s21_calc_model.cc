#include "./s21_calc_model.h"

namespace s21 {

CalcModel::CalcModel() {}

CalcModel::~CalcModel() {}

void CalcModel::StartCalc(const std::string& src_str, double X_num) {
  CleanStacks();
  setlocale(LC_NUMERIC, "C");
  if (ValidationEqual(src_str)) {
    try {
      result_ = Calc(src_str, X_num);
      
      if (std::isnan(result_))
        PushError("error: undefined");
    } catch (const std::exception& e) {
      PushError(e.what());
    }
  } else {
    throw std::invalid_argument("expression error");
  }
}

void CalcModel::CalcCredit(std::array<double, 3> data) {
  if (data[summa] < 1 || data[srok] < 1 || data[percent] < 0)
    PushError("uncorrect data");
  if (data[summa] > 1000000000000 || data[srok] > 420 || data[percent] > 100)
    PushError("redundant data");
  data[percent] = data[percent] / 1200;
  credit_data_[monthly] =
      data[summa] * (data[percent] * pow((1 + data[percent]), data[srok]) /
                     (pow((1 + data[percent]), data[srok]) - 1));
  credit_data_[itog] = credit_data_[monthly] * data[srok];
  credit_data_[pereplata] = credit_data_[itog] - data[summa];
}

bool CalcModel::ValidationEqual(const std::string& str) const noexcept {
  bool valid(false);
  std::string tmp("+-/*M^@ABCDEFGH)(1234567890.eX");
  for (std::string::const_iterator it = str.begin(); it != str.end(); ++it) {
    const char c = *it;
    if (tmp.find(c) == std::string::npos) {
      valid = false;
      break;
    }
    valid = true;
  }
  return valid;
}

void CalcModel::DifferenCalc(std::array<double, 3> data) {
  if (data[summa] < 1 || data[srok] < 1 || data[percent] < 0)
    PushError("uncorrect data");
  if (data[summa] > 1000000000000 || data[srok] > 420 || data[percent] > 100)
    PushError("redundant data");
  int debt_sum = data[summa] / data[srok];
  different_data_.clear();
  different_data_.resize(3 + data[srok]);
  for (int i = 0; i < data[srok]; i++) {
    different_data_[monthly + i] =
        (data[summa] * data[percent] / 100 * 31 / 365) + debt_sum;
    data[summa] -= debt_sum;
    different_data_[itog] += different_data_[monthly + i];
  }
  different_data_[pereplata] = different_data_[itog] - debt_sum * data[srok];
}

double CalcModel::Calc(const std::string& calc_src, double X_num) {
  int position = 0;
  while (calc_src[position]) {  //  Главный цикл вычисления
    StackType st_buf =
        ParserUno(calc_src, &position, X_num);  //  Парсим одну лексемму
    if (st_buf.prio) {  //  Если получили операцию или скобку
      while (st_buf.oper_val) {
        if (st_buf.oper_val == ')' && BracketFinder()) {
          //  Если пришла скобка закр а в стеке скобка откр
          oper_stack_.pop();
          st_buf.oper_val = 0.0;
        } else if (UnarCheck(st_buf.oper_val, calc_src, position)) {
          oper_stack_.push({0.0, st_buf.oper_val, st_buf.prio});
          num_stack_.push(0.0);  //  Получили унарный знак
          st_buf.oper_val = 0.0;
        } else if (oper_stack_.empty() || oper_stack_.top().oper_val == '(') {
          // Если стэк пуст или в нём скобка
          oper_stack_.push({0.0, st_buf.oper_val, st_buf.prio});
          st_buf.oper_val = 0.0;
        } else if (st_buf.prio >
                   oper_stack_.top().prio) {  //  Если приоритет опреации
          oper_stack_.push({0.0, st_buf.oper_val,  //  больше приоритета
                            st_buf.prio});  //  в стеке
          st_buf.oper_val = 0.0;
        } else {
          double buf_num = MathOperations();  //  Выполнить расчёт
          num_stack_.push(buf_num);
        }  //  т.к. остальные условия не прошли
      }
      position++;
    } else {  //  Если получили число
      num_stack_.push(st_buf.val_dub);
    }
  }

  while (!oper_stack_.empty()) {  //  Расчёт оставшегося содержимого стеков
    if (BracketFinder()) {
      oper_stack_.pop();
      //  Если забыли поставить скобки в конце уравнения
      continue;
    }
    //  Если пришло число, просто отправляем в стек чисел
    num_stack_.push(MathOperations());
  }

  double result = 0.0;
  if (!num_stack_.empty()) {
    result = num_stack_.top();
    num_stack_.pop();
  } else {
    PushError("numbers stack empty");
  }
  if (!num_stack_.empty())
    PushError("numbers stack invalid");
  return result;
}

//  Парсер одной лексеммы
CalcModel::StackType CalcModel::ParserUno(const std::string& calc_src,
                                          int* position, double X_num) {
  StackType stack1{};
  int prio = PrioCheck(calc_src[*position]);
  if (prio) {
    stack1.prio = prio;
    stack1.oper_val = calc_src[*position];
  } else {
    if (calc_src[*position] == 'X') {
      stack1.prio = 0;
      stack1.val_dub = X_num;
      *position += 1;
    } else {
      std::string buf{};
      *position = *position + BufferingNumber(&calc_src[*position], buf);
      stack1.prio = prio;
      stack1.val_dub = std::stod(buf);
    }
  }
  return stack1;
}

//  Определение приоритета опреатора
int CalcModel::PrioCheck(const char src_string) const noexcept {
  int prior{};
  int position_num = PositionCounter(src_string);
  if (position_num > 16)
    prior = 0;
  else if (position_num == 16)
    prior = 5;
  else if (position_num > 6)
    prior = 4;
  else if (position_num == 6)
    prior = 3;
  else if (position_num > 2)
    prior = 2;
  else if (position_num >= 0)  //  Низкий приоритет для закрывающей скобки
    prior = 1;  //  запустит подсчёт
  return prior;
}

int CalcModel::PositionCounter(char src_string)
    const noexcept {  //  Подсчёт позиции операции строке приоритетов
  const char* operators = ")+-/*M^@ABCDEFGH(";
  int counter{};
  while (operators[counter]) {
    if (operators[counter] == src_string) {
      break;
    }
    counter++;
  }
  return counter;
}

int CalcModel::BufferingNumber(
    const char* src_string,
    std::string& out) {  //  Сборка числа в строку, возвращает длинну числа
  int i = 0;
  while ((src_string[i] >= '0' && src_string[i] <= '9') ||
         src_string[i] == '.' || src_string[i] == 'e') {
    if (src_string[i] == 'e') {
      out += src_string[i];
      i++;
    }
    out += src_string[i];
    i++;
  }
  return i;
}

int CalcModel::BracketFinder() {
  int finded = 0;
  if (!oper_stack_.empty())
    if (oper_stack_.top().oper_val == '(')
      finded = 1;
  return finded;
}

int CalcModel::UnarCheck(char check, const std::string& calc_str,
                         int position) {
  int unar_minus_find{};
  if ((check == '-' || check == '+') && !position)
    unar_minus_find = 1;
  if ((check == '-' || check == '+') && position > 0)
    if (calc_str[position - 1] == '(')
      unar_minus_find = 1;
  return unar_minus_find;
}

double CalcModel::MathOperations() {
  double buf_num = 0.0;
  if (oper_stack_.top().prio < 4) {
    if (num_stack_.size() < 2)
      PushError("Math err, nums empty");
    double second = num_stack_.top();
    num_stack_.pop();
    double first = num_stack_.top();
    num_stack_.pop();
    char operat = oper_stack_.top().oper_val;
    oper_stack_.pop();
    buf_num = SimpleMath(second, first, operat);
  } else if (oper_stack_.top().prio < 5) {
    if (num_stack_.empty())
      PushError("Math err, expression");
    buf_num = num_stack_.top();
    num_stack_.pop();
    char oper_buf = oper_stack_.top().oper_val;
    oper_stack_.pop();
    buf_num = TrigonCalc(buf_num, oper_buf);
  }
  return buf_num;
}

double CalcModel::SimpleMath(double second_num, double first_num,
                             char operation) {
  double out_num = 0.0;
  double epsilon = 0.00000001;
  switch (operation) {
    case '+':
      out_num = first_num + second_num;
      break;
    case '-':
      out_num = first_num - second_num;
      break;
    case '*':
      out_num = first_num * second_num;
      break;
    case '/':
      if (std::abs(second_num - 0.0) < epsilon)
        PushError("Error: /0");
      if (std::abs(first_num - 0.0) < epsilon)
        PushError("Error: 0/");
      out_num = first_num / second_num;
      break;
    case '^':
      out_num = pow(first_num, second_num);
      break;
    case 'M':
      out_num = fmod(first_num, second_num);
      break;
  }
  return out_num;
}

double CalcModel::TrigonCalc(double x, char operation) {
  double buf_num = 0.0;
  switch (operation) {
    case COS:
      buf_num = cos(x);
      break;
    case SIN:
      buf_num = sin(x);
      break;
    case TAN:
      buf_num = tan(x);
      break;
    case ACOS:
      if (x < -1 || x > 1)
        PushError("error: interval [-1,+1]");
      buf_num = acos(x);
      break;
    case ASIN:
      if (x < -1 || x > 1)
        PushError("error: interval [-1,+1]");
      buf_num = asin(x);
      break;
    case ATAN:
      buf_num = atan(x);
      break;
    case SQRT:
      if (x < -1 || x > 1)
        PushError("error: interval > 1");
      buf_num = sqrt(x);
      break;
    case LN:
      if (x < 1)
        PushError("error: interval > 1");
      buf_num = log(x);
      break;
    case LOG:
      if (x < 1)
        PushError("error: interval > 1");
      buf_num = log10(x);
      break;
  }
  return buf_num;
}

void CalcModel::CleanStacks() {
  while (!num_stack_.empty()) {
    num_stack_.pop();
  }
  while (!oper_stack_.empty()) {
    oper_stack_.pop();
  }
}

void CalcModel::PushError(std::string error) {
  throw std::runtime_error(error);
}

}  // namespace s21
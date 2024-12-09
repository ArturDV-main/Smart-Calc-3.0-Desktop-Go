#include "./s21_calc_controller.h"
#include "./s21_smartcalc.h"

struct Resp StartCalc(const char *source, const double num_x)
{
  struct Resp result;
  result.err = 0;
  result.errors = NULL;
  result.result = 0.0;
  size_t length = strlen(source) + 1; // +1 для терминального нуля
  char *destination = malloc(length);
  if (destination == NULL)
  {
    result.err = 5;
    result.errors = "memory alloc error";
    return result;
  }
  strcpy(destination, source); // Можно использовать strncpy(destination, source, length)
  struct Response r;
  r = calculating(destination, num_x);
  free(destination);
  if (r.err != 0) {
    result.err = 1;
    result.errors = r.errors;
    return result;
  }
  return result;
}

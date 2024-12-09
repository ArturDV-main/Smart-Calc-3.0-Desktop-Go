CC = gcc
FLAGS = -Wall -Werror -Wextra -std=c11
GCOV = --coverage
UNIT = unit_tests
STLIB = s21_smart_calc.a
OS = $(shell uname -s)
BUILD_DIR = ../build

ifeq ($(OS), Darwin)
    LIBS := -lcheck
	APPLICATION = cppfront.app
	OPEN = open
else
    LIBS := -lcheck_pic -lpthread -lrt -lm -lsubunit -g
	APPLICATION = cppfront
	OPEN = ./
endif

all:

lib:
	mkdir -p $(BUILD_DIR)
ifeq ($(OS), Darwin)
	cp ./src/* $(BUILD_DIR)
	cd $(BUILD_DIR) && gcc -std=c11 -dynamiclib -fPIC -o libsmart_calc.dylib ./*.c -lm
	cd $(BUILD_DIR) && cp ./libsmart_calc.dylib ../
	cd $(BUILD_DIR) && cp ./libsmart_calc.dylib ../pkg/calcadapter/
else
	cp ./src/* $(BUILD_DIR)
	cd $(BUILD_DIR) && gcc -std=c11 -shared -fPIC -o libsmart_calc.so ./*.c -lm
	cd $(BUILD_DIR) && cp ./libsmart_calc.so ../
	cd $(BUILD_DIR) && cp ./libsmart_calc.so ../pkg/calcadapter/
endif


s21_smart_calc.a: build_s21_smart_calc

build_s21_smart_calc:
	$(CC) -c s21_smartcalc.c
	ar rcs $(STLIB) s21_*.o
	ranlib $(STLIB)

unit_tests.o:
	$(CC) $(FLAGS) -c unit_tests.c -lm

test: $(STLIB)
	$(CC) $(FLAGS) unit_tests.c $(STLIB) -o $(UNIT) $(LIBS)
	./$(UNIT)

gcov_report: $(STLIB) unit_tests.o
	$(CC) $(FLAGS) $(GCOV) unit_tests.o s21_smartcalc.c -o $(UNIT) $(LIBS)
	./$(UNIT)
	lcov -t "test" -o test.info -c -d .
	genhtml -o report test.info
	open report/index.html

clang:
	clang-format -n *.c
	clang-format -i *.c
	
install:
	make clean
	mkdir smartcalc
	cd cppfront && qmake && make && make clean && rm Makefile && cd ../ && mv cppfront/$(APPLICATION) ./smartcalc
	open ./smartcalc

open:
	cd smartcalc && $(OPEN) $(APPLICATION)

uninstall:
	rm -rf smartcalc

dvi:
	open dvi.html

dist:
	rm -rf Archive_SmartCalc_v1.0/
	mkdir Archive_SmartCalc_v1.0/
	mkdir Archive_SmartCalc_v1.0/src
	mv ./smartcalc/$(APPLICATION) Archive_SmartCalc_v1.0/src/
	tar cvzf Archive_SmartCalc_v1.0.tgz Archive_SmartCalc_v1.0/
	rm -rf Archive_SmartCalc_v1.0/

coverage:
	open report/index.html

clean:
	rm -rf *.o *.a *.gc* *.info test report *dSYM $(UNIT) Archive_SmartCalc_v1.0 *tgz smartcalc

g: clean
	git add .
	git commit -m"develop-s21_smartcalc.c"
	git push origin develop


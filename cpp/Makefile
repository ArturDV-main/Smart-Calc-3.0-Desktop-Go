CXX = gcc
CPP_STD = -std=c++17
TARGET = SmartCalc2_0
CXXFLAGS = -lm -g -Wall -Wextra -Werror -lstdc++ --coverage
GT_FLAGS = -lgtest -lgtest_main -lm

#  Project directories
BUILD_DIR = build
SRC_DIRS = src src/s21_view_qt
GT_DIRS = src src/google_tests
#  Project sourses
SRCS := $(shell find $(SRC_DIRS) -maxdepth 1 -name *.cc)
OBJS := $(SRCS:%=$(BUILD_DIR)/%.o)

#  Google test sourses
GT_SRCS := $(shell find $(GT_DIRS) -maxdepth 1 -name s21_*.cc)
GT_OBJS := $(GT_SRCS:%=$(BUILD_DIR)/%.o)

OS := $(shell uname -s)

ifeq ($(OS), Darwin)
	APPLICATION := SmartCalc2_0.app
	OPEN = open $(BUILD_DIR)/$(APPLICATION)
else
	APPLICATION := SmartCalc2_0
	OPEN = ./$(BUILD_DIR)/$(APPLICATION)
endif

all: apple

install: apple
ifeq ($(OS), Darwin)
	cp -r build/$(APPLICATION) ~/Desktop/
else
	mkdir -p SmartCalc
	cp build/$(APPLICATION) ./SmartCalc/
endif

uninstall:
ifeq ($(OS), Darwin)
	rm -rf ~/Desktop/$(APPLICATION)
else
	rm -rf ./SmartCalc/
endif

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR) test.info report src/s21_view_qt/.qmake.stash \
	src/build-SmartCalc2_0-Desktop-Debug build-SmartCalc2_0-Desktop_x86_darwin_generic_mach_o_64bit-Debug

dvi:
	open dvi.html

dist: all
	rm -rf archive_smart_calc_2_0
	mkdir -p archive_smart_calc_2_0
	mkdir -p archive_smart_calc_2_0/src
	cp $(BUILD_DIR)/$(APPLICATION) ./archive_smart_calc_2_0/
	cp -r ./src ./archive_smart_calc_2_0/
	cp ./*.html ./archive_smart_calc_2_0/
	cp ./*.md ./archive_smart_calc_2_0/
	tar cvzf archive_smart_calc_2_0.tgz archive_smart_calc_2_0
	rm -rf archive_smart_calc_2_0/

#  Google tests
tests: clean $(GT_OBJS)
	mkdir -p $(BUILD_DIR)
	$(CXX) $(CPP_STD) $(GT_OBJS) $(GT_FLAGS) $(CXXFLAGS) -o $(BUILD_DIR)/gtest.out
	./$(BUILD_DIR)/gtest.out

gcov_report: clean tests
ifeq ($(OS), Darwin)
	cd $(BUILD_DIR) && lcov --ignore-errors mismatch -t "test"  -o test.info -c -d .
else
	cd $(BUILD_DIR) && lcov -t "test"  -o test.info -c -d .
endif
	cd $(BUILD_DIR) && lcov --remove test.info '/usr/local/include/*' -o test.info
	cd $(BUILD_DIR) && genhtml -o report test.info
	open $(BUILD_DIR)/report/index.html

#  SmartCallc2.0 application
apple:
	mkdir -p $(BUILD_DIR)
	cd $(BUILD_DIR) && qmake CONFIG+=qtquickcompiler ../src/s21_view_qt/SmartCalc2_0.pro && make

# Build step for C++ source
$(BUILD_DIR)/%.cc.o: %.cc
	mkdir -p $(dir $@)
	$(CXX) $(CPP_STD) $(CXXFLAGS) -c $< -o $@

clang:
	clang-format -style=file:materials/linters/.clang-format -n src/*.cc src/google_tests/*.cc src/*h src/s21_view_qt/*.cc src/s21_view_qt/*.h
	clang-format -style=file:materials/linters/.clang-format -i src/*.cc src/google_tests/*.cc src/*h src/s21_view_qt/*.cc src/s21_view_qt/*.h

open: all
	$(OPEN)

leak:
ifeq ($(OS), Darwin)
	echo $(OS)
	echo "For Aple --------------------"
	leaks -atExit -- ./$(BUILD_DIR)/$(APPLICATION)/Contents/MacOS/SmartCalc2_0
else
	echo $(OS)
	echo "For Ubuntu --------------------"
	CK_FORK=no valgrind --vgdb=no --leak-check=full --show-leak-kinds=all --track-origins=yes --verbose --log-file=$(BUILD_DIR)/RESULT_VALGRIND.txt $(BUILD_DIR)/$(APPLICATION)
	grep errors $(BUILD_DIR)/RESULT_VALGRIND.txt
endif

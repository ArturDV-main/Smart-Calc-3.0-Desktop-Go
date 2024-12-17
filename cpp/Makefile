CXX = g++
CPP_STD = -std=c++17
TARGET = smart_calc
CXXFLAGS = -lm -g -Wall -Wextra -Werror
GCOVFLAGS = --coverage
GT_FLAGS = -lgtest -lgtest_main -lm

#  Project directories
LIB_DIR = ../libs
BUILD_DIR = ../build
SRC_DIRS = ./src
GT_DIRS = ./src ./google_tests

#  Project sourses
SRCS := $(shell find $(SRC_DIRS)  -type f -name "*.cc")
HEADERS := $(shell find $(SRC_DIRS)  -type f -name "*.h")
OBJS := $(SRCS:%=$(BUILD_DIR)/%.o)

#  Google test sourses
GT_SRCS := $(shell find $(GT_DIRS) -type f -name "*.cc")
GT_OBJS := $(GT_SRCS:%=$(BUILD_DIR)/%.o)

OS := $(shell uname -s)

.PHONY: clean
all: lib

# -Wl,-install_name,@rpath/libsmart_calc.dylib

lib:
	mkdir -p $(BUILD_DIR)
ifeq ($(OS), Darwin)
	cp ./src/* $(BUILD_DIR)
	cd $(BUILD_DIR) && g++ -static-libstdc++ -static-libgcc -lstdc++ -lpthread -dynamiclib -fPIC -o libsmart_calc.dylib ./*.cc
	cd $(BUILD_DIR) && cp ./libsmart_calc.dylib ../
	cd $(BUILD_DIR) && cp ./libsmart_calc.dylib ../pkg/calcadapter/
else
	cp ./src/* $(BUILD_DIR)
	cd $(BUILD_DIR) && g++ -static-libstdc++ -static-libgcc -lstdc++ -lpthread -shared -fPIC -o libsmart_calc.so ./*.cc
	cd $(BUILD_DIR) && cp ./libsmart_calc.so ../
	cd $(BUILD_DIR) && cp ./libsmart_calc.so ../pkg/calcadapter/
	cd $(BUILD_DIR) && sudo cp ./libsmart_calc.so /usr/lib/
endif

clean:
	rm -rf $(BUILD_DIR) report $(shell find $(GT_DIRS) -type f -name "*.o")

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
	$(CXX) $(CPP_STD) $(GT_OBJS) $(GT_FLAGS) $(CXXFLAGS) -o $(BUILD_DIR)/gtest.out $(GCOVFLAGS)
	./$(BUILD_DIR)/gtest.out

# Coverage report create
gcov_report: clean tests
ifeq ($(OS), Darwin)
	cd $(BUILD_DIR) && lcov --ignore-errors mismatch --ignore-errors inconsistent -t "test"  -o test.info -c -d .
else
	cd $(BUILD_DIR) && lcov -t "test"  -o test.info -c -d .
endif
	cd $(BUILD_DIR) && lcov --remove test.info '/usr/local/include/*' -o test.info
	cd $(BUILD_DIR) && genhtml -o report test.info
	open $(BUILD_DIR)/report/index.html

# Build step for C++ source
$(BUILD_DIR)/%.cc.o: %.cc
	mkdir -p $(dir $@)
	$(CXX) $(CPP_STD) $(CXXFLAGS) -fPIC -c $< -o $@ $(GCOVFLAGS)

# Clang check and fix
clang:
	clang-format -style=file:./.clang-format -n $(SRCS) $(HEADERS) $(GT_SRCS)
	clang-format -style=file:./.clang-format -i $(SRCS) $(HEADERS) $(GT_SRCS)

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

all: lib open
	wails build
	./build/bin/smartcalc

test:
	cd ./cpp && make
	sudo cp ./build/libsmart_calc.so /usr/lib/
	cd ./pkg/calcadapter && go test

clean:
	rm -rf ./build/*

lib:
	cd ./cpp && make lib

open:
ifeq ($(OS), Darwin)
	open ./build/bin/leftrana/smartcalc.app
else
    LIBS := -lgtest -lstdc++ -lcheck_pic -lrt -lpthread -lsubunit -lm -g
	APPLICATION = 3dviewer
	OPEN = ./$(APPLICATION)
endif
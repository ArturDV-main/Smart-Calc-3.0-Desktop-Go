OS := $(shell uname -s)

all: lib app open

open:
ifeq ($(OS), Darwin)
	export DYLD_LIBRARY_PATH="/Users/leftrana/projects/Smart-Calc-3.0-Desktop-Go/pkg/calcadapter" && open /Users/leftrana/projects/Smart-Calc-3.0-Desktop-Go/build/bin/smartcalc.app
else
	./build/bin/smartcalc
endif

app:
	cp ./appicon.png ./build
	wails build -tags webkit2_41

gotest:
	cd ./pkg/calcadapter && go test

test:
	cd ./cpp && make
	sudo cp ./build/libsmart_calc.so /usr/lib/
	cd ./pkg/calcadapter && go test

clean:
	rm -rf ./build/*

lib:
	cd ./cpp && make lib

dev:
	wails dev -tags webkit2_41
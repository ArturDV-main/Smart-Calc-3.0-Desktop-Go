OS := $(shell uname -s)

all: lib app open

open:
ifeq ($(OS), Darwin)
	open /Users/leftrana/projects/Smart-Calc-3.0-Desktop-Go/build/bin/smartcalc.app
else
	./build/bin/smartcalc
endif

app:
	wails build -tags webkit2_41


test:
	cd ./cpp && make
	sudo cp ./build/libsmart_calc.so /usr/lib/
	cd ./pkg/calcadapter && go test

clean:
	rm -rf ./build/*

lib:
	cd ./cpp && make lib
ifeq ($(OS), Darwin)
	export DYLD_LIBRARY_PATH="/Users/leftrana/projects/Smart-Calc-3.0-Desktop-Go/pkg/calcadapter"
endif

dev:
	wails dev -tags webkit2_41
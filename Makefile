all: lib app open

open:
ifeq ($(OS), Darwin)
	open ./build/bin/leftrana/smartcalc.app
else
	./build/bin/smartcalc
endif

app:
	wails build


test:
	cd ./cpp && make
	sudo cp ./build/libsmart_calc.so /usr/lib/
	cd ./pkg/calcadapter && go test

clean:
	rm -rf ./build/*

lib:
	cd ./cpp && make lib

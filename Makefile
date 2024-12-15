OS := $(shell uname -s)
.PHONY: clean pkg

all: lib app open

pkg:
	pkgbuild --root ./build/my.pkg/smartcalc.app --identifier com.yourcompany.myapp --version 1.0 My.pkg

open:
ifeq ($(OS), Darwin)
	export DYLD_LIBRARY_PATH="/Users/leftrana/projects/Smart-Calc-3.0-Desktop-Go/pkg/calcadapter" && open /Users/leftrana/projects/Smart-Calc-3.0-Desktop-Go/build/bin/smartcalc.app
else
	./build/bin/smartcalc
endif

dvi:
	cd cpp && make dvi

# lib app
dist: 
	rm -rf archive_smart_calc_3_0
	mkdir -p archive_smart_calc_3_0
	mkdir -p archive_smart_calc_3_0/src
	rsync -av ./* ./archive_smart_calc_3_0/src/. --exclude build --exclude history.txt --exclude *.dylib --exclude *.gitignor --exclude *.pkg --exclude archive_smart_calc_3_0
	

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

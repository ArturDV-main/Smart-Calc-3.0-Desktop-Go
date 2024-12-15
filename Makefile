OS := $(shell uname -s)
.PHONY: clean pkg

all: lib app open

pkg:
	pkgbuild --root ./build/bin/smartcalc.app --identifier com.yourcompany.myapp --version 1.0 Smartcalc.pkg

open:
ifeq ($(OS), Darwin)
	export DYLD_LIBRARY_PATH="/Users/leftrana/projects/Smart-Calc-3.0-Desktop-Go/pkg/calcadapter" && open /Users/leftrana/projects/Smart-Calc-3.0-Desktop-Go/build/bin/smartcalc.app
else
	./build/bin/smartcalc
endif

install: all

uninstall: clean

dvi:
	cd cpp && make dvi

# lib app
# otool -L ./MyApp.app/Contents/MacOS/MyApp
dist: 
ifeq ($(OS), Darwin)
	rm -rf archive_smart_calc_3_0
	mkdir -p archive_smart_calc_3_0
	mkdir -p archive_smart_calc_3_0/src
	rsync -av ./* ./archive_smart_calc_3_0/src/. --exclude build --exclude history.txt --exclude *.dylib --exclude *.gitignor --exclude *.pkg --exclude archive_smart_calc_3_0
	rsync -av ./build/bin/smartcalc.app ./archive_smart_calc_3_0
	rsync -av ./build/libsmart_calc.dylib ./archive_smart_calc_3_0
	tar cvzf archive_smart_calc_3_0.tgz archive_smart_calc_3_0
	rm -rf archive_smart_calc_3_0/

else
	rm -rf archive_smart_calc_3_0
	mkdir -p archive_smart_calc_3_0
	mkdir -p archive_smart_calc_3_0/src
	rsync -av ./* ./archive_smart_calc_3_0/src/. --exclude build --exclude history.txt --exclude *.dylib --exclude *.gitignor --exclude *.pkg --exclude archive_smart_calc_3_0
	cp ./build/bin/smartcalc ./archive_smart_calc_3_0
	rsync -av ./build/libsmart_calc.so ./archive_smart_calc_3_0
	tar cvzf archive_smart_calc_3_0.tgz archive_smart_calc_3_0
	rm -rf archive_smart_calc_3_0/
endif

app:
	cp ./appicon.png ./build
ifeq ($(OS), Darwin)
	wails build
else
	wails build -tags webkit2_41
endif
	sh ./apptool.sh
	make pkg

tests: lib
	cd ./pkg/calcadapter && go test


clean:
	rm -rf ./build/*

lib:
	cd ./cpp && make lib

dev:
	wails dev -tags webkit2_41

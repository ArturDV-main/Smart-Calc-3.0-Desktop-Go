all:
	wails build
	cp ../pkg/calcadapter/lib
	./build/bin/smartcalc

test:
	cd ./cpp && make
	cd ./pkg/calcadapter && go test


clean:
	rm -rf ./build/*
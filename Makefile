all:
	wails build
	./build/bin/smartcalc

test:
	cd ./cpp && make
	cd ./pkg/calcadapter && go test


clean:
	rm -rf ./build/*
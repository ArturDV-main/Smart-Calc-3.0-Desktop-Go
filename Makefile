all:
	wails build -tags webkit2_41
	./build/bin/smartcalc
test:
	cd ./cpp && make
	cd ./pkg/calcadapter && go test .


clean:
	rm -rf ./build/*
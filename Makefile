all:
	cd ./cpp && make
	cd ./pkg/calcadapter && go test .
	cd ./cmd/smart_calculator && go build . && ./smart_calculator


clean:
	rm -rf ./build/*
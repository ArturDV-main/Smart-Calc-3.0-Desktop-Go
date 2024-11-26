all:
	cd ./cpp && make
	cp ./cmd/* ./build/
	cd ./build && go build . && ./smartcalc
	
.PHONE: build
build:
	go build -o build/demo . && rm -rf build

.PHONE: run
run:
	go build -o build/demo . && ./build/demo

.PHONE: clean
clean:
	rm -rf ./build
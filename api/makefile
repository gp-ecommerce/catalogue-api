GOCMD = /usr/local/go/bin/go
BUILD_PATH = ./build
BINARY = $(BUILD_PATH)/catalogue_api 
TEST_PATH= ./test/...

.PHONY: build test

build:
	mkdir -p build
	$(GOCMD) build -o $(BINARY) main.go

run: build
	$(BINARY)
test:
	$(GOCMD) test -v ${TEST_PATH}

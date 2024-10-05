BINARY_NAME=$(notdir $(shell pwd))

build:
	@go build -o bin/${BINARY_NAME}
# run: build
# 	@./bin/api
run:
	@air
test:
	@go test -v ./...




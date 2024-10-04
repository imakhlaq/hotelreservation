build:
	@go build -o bin/api
# run: build
# 	@./bin/api
run:
	@air
test:
	@go test -v ./...




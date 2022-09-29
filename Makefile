all:  build run

run:
	bin/main

build:
	go build cmd/main.go -o bin/main


lint:
	golangci-lint run ./...

test:
	go test ./...


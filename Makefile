LAST_VERSION = $(git tag -l --sort=-v:refname | head -n 1)

all: build run

run:
	bin/main

build:
	go build -o bin/main cmd/main.go


lint:
	golangci-lint run ./...

test:
	go test ./...

release:
	./release.sh

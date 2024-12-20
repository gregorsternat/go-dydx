format:
	go fmt ./...

lint:
	golangci-lint run ./...

build:
	go build -o bin/ ./...

.PHONY: format lint build
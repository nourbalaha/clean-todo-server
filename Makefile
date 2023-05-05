lint:
	golangci-lint run --fix

test:
	go test -race -v ./...

build:
	go build ./cmd

run-app:
	go run ./cmd

.PHONY: lint test build run-app oapi
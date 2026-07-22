DAY ?= 01

.PHONY: run fmt lint

run:
	go run ./day$(DAY) < day$(DAY)/day$(DAY).input

fmt:
	go fmt ./...

lint:
	go vet ./...

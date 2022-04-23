.PHONY: build
build:
	go build -v ./cmd/app

.PHONY: run
run:
	go run ./cmd/app
	
.DEFAULT_GOAL := build
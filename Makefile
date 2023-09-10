prepare:
	go mod download
.PHONY: prepare

build:
	make prepare && go build .
.PHONY: build

start:
	make prepare && go run main.go api
.PHONY: start
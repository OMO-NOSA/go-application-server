prepare:
	go mod download
.PHONY: prepare

build:
	make prepare && go build .
.PHONY: build

start:
	make prepare && go run main.go api
.PHONY: start

codegen:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen && \
	oapi-codegen --config=.oapi-codegen.yml api/api.yml && \
	go mod tidy
.PHONY: generate-api 
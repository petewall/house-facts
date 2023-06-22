.PHONY: build lint

SRC = $(shell find . -name "*.go" | grep -v "_test\." )

lint: $(SRC)
	golangci-lint run

VERSION := $(or $(VERSION), dev)
LDFLAGS="-X github.com/petewall/house-facts/cmd.version=$(VERSION)"

build: build/house-facts

build/house-facts: $(SRC)
	go build -o build/house-facts -ldflags ${LDFLAGS} ./main.go

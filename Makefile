.PHONY: build build-image lint test test-units

SRC = $(shell find . -name "*.go" | grep -v "_test\." )

lint: $(SRC)
	golangci-lint run

test-units: $(SRC)
	ginkgo -r .

test: lint test-units

VERSION := $(or $(VERSION), dev)
LDFLAGS="-X github.com/petewall/house-facts/cmd.version=$(VERSION)"

build: build/house-facts

build/house-facts: $(SRC)
	go build -o build/house-facts -ldflags ${LDFLAGS} ./main.go

build-image: Dockerfile
	docker build . --tag petewall/house-facts

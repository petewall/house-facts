.PHONY: build build-image lint test test-units

HAS_GINKGO := $(shell command -v ginkgo;)
HAS_GOLANGCI_LINT := $(shell command -v golangci-lint;)
PLATFORM := $(shell uname -s)

SRC = $(shell find . -name "*.go" | grep -v "_test\." )

deps:
	go mod download

lint: $(SRC) deps
ifndef HAS_GOLANGCI_LINT
ifeq ($(PLATFORM), Darwin)
	brew install golangci-lint
endif
ifeq ($(PLATFORM), Linux)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
endif
endif
	golangci-lint run

test-units: $(SRC) deps
ifndef HAS_GINKGO
	go install github.com/onsi/ginkgo/v2/ginkgo
endif
	ginkgo -r .

test: lint test-units

VERSION := $(or $(VERSION), dev)
LDFLAGS="-X github.com/petewall/house-facts/cmd.version=$(VERSION)"

build: build/house-facts

build/house-facts: $(SRC)
	go build -o build/house-facts -ldflags ${LDFLAGS} ./main.go

build-image: Dockerfile
	docker build . --tag petewall/house-facts

set-pipeline: ci/pipeline.yaml
	fly --target wallhouse set-pipeline \
		--pipeline house-facts \
		--config ci/pipeline.yaml
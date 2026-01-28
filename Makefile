NAME := connections
DESC := NYTimes Connections Solver
VERSION := $(shell git describe --tags --always --dirty --match "[0-9]*.[0-9]*.[0-9]*" 2>/dev/null || echo "0.1.0-dev")
GOVERSION := 1.22.0
BUILDVERSION := $(shell go version)
BUILDTIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILDER := $(shell echo "`git config user.name` <`git config user.email`>" 2>/dev/null || echo "unknown")
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.buildTime=$(BUILDTIME)' \
           -X 'main.builder=$(BUILDER)' \
           -X 'main.goversion=$(BUILDVERSION)' \
           -X 'main.name=$(NAME)'

build: test clean target/local

staticcheck:
	staticcheck ./...

lint:
	golangci-lint -v run ./...

test:
	go test ./...

target/local: modules
	mkdir -p target/local/bin && go build -ldflags "$(LDFLAGS)" -o target/local/bin/connections ./cmd/cli

modules:
	go mod tidy

clean-target:
	rm -rf target

clean: clean-target

run:
	@go run ./cmd/cli/main.go

.PHONY: clean lint modules build run staticcheck test clean-target

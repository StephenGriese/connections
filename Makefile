NAME := connections
DESC := NYTimes Connections Puzzle Solver with AI
PREFIX ?= usr/local
VERSION := $(shell git describe --tags --always --dirty --match "[0-9]*.[0-9]*.[0-9]*")
GOVERSION := 1.19.0
BUILDVERSION := $(shell go version)
BUILDTIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILDER := $(shell echo "`git config user.name` <`git config user.email`>")
PROJECT_URL := "https://github.com/StephenGriese/$(NAME)"
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.buildTime=$(BUILDTIME)' \
           -X 'main.builder=$(BUILDER)' \
           -X 'main.goversion=$(BUILDVERSION)' \
           -X 'main.name=$(NAME)'

build: staticcheck lint test clean target/local

init:
	git config core.hooksPath .githooks

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
	@echo "Running Connections Solver..."
	@./target/local/bin/connections

.PHONY: clean lint modules build run

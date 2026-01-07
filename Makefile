# Project variables
BINARY_NAME=platform-cli
BUILD_DIR=bin
MAIN_PATH=./cmd/platform-cli/main.go
MODULE_NAME=github.com/abevz/platform-iac-cli

.PHONY: all build test lint clean help

all: lint test build

## build: Compile the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

## test: Run unit tests with race detection
test:
	@echo "Running tests..."
	go test -race -v ./...

## lint: Run golangci-lint (system binary required)
lint:
	@echo "Running linter..."
	golangci-lint run

## clean: Remove build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out

## init: Initialize infrastructure via compiled binary
init: build
	./$(BUILD_DIR)/$(BINARY_NAME) init --verbose

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

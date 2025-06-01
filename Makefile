# Makefile for find_whats_hidden

BINARY_NAME=find_whats_hidden
GO=go
GOFLAGS=-v
LDFLAGS=-s -w

# Default target
.DEFAULT_GOAL := build

.PHONY: all
all: clean test build

.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	$(GO) build $(GOFLAGS) -ldflags="$(LDFLAGS)" -o $(BINARY_NAME) .

.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test $(GOFLAGS) ./...

.PHONY: clean
clean:
	@echo "Cleaning..."
	$(GO) clean
	rm -f $(BINARY_NAME)

.PHONY: run
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)

.PHONY: install
install: build
	@echo "Installing $(BINARY_NAME)..."
	$(GO) install $(GOFLAGS)

# Cross compilation
.PHONY: build-all
build-all: build-linux build-darwin build-windows

.PHONY: build-linux
build-linux:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 $(GO) build -ldflags="$(LDFLAGS)" -o $(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=386 $(GO) build -ldflags="$(LDFLAGS)" -o $(BINARY_NAME)-linux-386 .

.PHONY: build-darwin
build-darwin:
	@echo "Building for macOS..."
	GOOS=darwin GOARCH=amd64 $(GO) build -ldflags="$(LDFLAGS)" -o $(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 $(GO) build -ldflags="$(LDFLAGS)" -o $(BINARY_NAME)-darwin-arm64 .

.PHONY: build-windows
build-windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 $(GO) build -ldflags="$(LDFLAGS)" -o $(BINARY_NAME)-windows-amd64.exe .
	GOOS=windows GOARCH=386 $(GO) build -ldflags="$(LDFLAGS)" -o $(BINARY_NAME)-windows-386.exe .

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build       - Build the binary"
	@echo "  test        - Run tests"
	@echo "  clean       - Clean build files"
	@echo "  run         - Build and run"
	@echo "  install     - Install the binary"
	@echo "  build-all   - Build for all platforms"
	@echo "  help        - Show this help message"
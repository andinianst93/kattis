.PHONY: help run build test clean install fmt vet lint dev

# Variables
APP_NAME := main
BUILD_DIR := bin
GO := go

# Default target
help:
	@echo "Available commands:"
	@echo "  make run      - Run the application (go run main.go)"
	@echo "  make dev      - Run with hot reload using air"
	@echo "  make build    - Build the application"
	@echo "  make test     - Run tests"
	@echo "  make clean    - Clean build artifacts"
	@echo "  make install  - Install dependencies"
	@echo "  make fmt      - Format code"
	@echo "  make vet      - Run go vet"
	@echo "  make lint     - Run linter"

# Run the application (like go run main.go)
run:
	$(GO) run main.go

# Run with hot reload (requires air)
dev:
	air

# Build the application
build:
	mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(APP_NAME) .

# Build for production
build-prod:
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 $(GO) build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) .

# Run tests
test:
	$(GO) test -v ./...

# Run tests with coverage
test-coverage:
	$(GO) test -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run benchmarks
benchmark:
	$(GO) test -bench=. -benchmem ./...

# Install dependencies
install:
	$(GO) mod download
	$(GO) mod tidy

# Install a new package (usage: make get PKG=github.com/some/package)
get:
	$(GO) get $(PKG)
	$(GO) mod tidy

# Format code
fmt:
	$(GO) fmt ./...

# Run go vet
vet:
	$(GO) vet ./...

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

# Install development tools
tools:
	$(GO) install github.com/cosmtrek/air@latest
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run the built binary
start:
	./$(BUILD_DIR)/$(APP_NAME)

# Build and run
build-run: build start

# Watch and restart on changes (simple version without air)
watch:
	@while true; do \
		make run; \
		inotifywait -qre close_write .; \
	done
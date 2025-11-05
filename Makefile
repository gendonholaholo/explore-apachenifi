.PHONY: all build test lint clean run help install

# Default target
all: test lint build

# Build the binary
build:
	@echo "Building nifiparser..."
	@go build -o bin/nifiparser ./cmd/nifiparser

# Run tests
test:
	@echo "Running tests..."
	@go test -v -race -cover ./...

# Run tests with coverage report
coverage:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run linter
lint:
	@echo "Running golangci-lint..."
	@golangci-lint run

# Run linter with auto-fix
lint-fix:
	@echo "Running golangci-lint with auto-fix..."
	@golangci-lint run --fix

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html

# Run the tool with sample data
run:
	@echo "Running nifiparser with sample data..."
	@go run ./cmd/nifiparser -file testdata/sample_nifi_config.json

# Install the binary to GOPATH/bin
install:
	@echo "Installing nifiparser..."
	@go install ./cmd/nifiparser

# Format code
fmt:
	@echo "Formatting code..."
	@gofmt -w .
	@goimports -w .

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

# Verify dependencies
verify:
	@echo "Verifying dependencies..."
	@go mod verify

# Help target
help:
	@echo "Available targets:"
	@echo "  all        - Run tests, lint, and build (default)"
	@echo "  build      - Build the binary"
	@echo "  test       - Run tests"
	@echo "  coverage   - Generate coverage report"
	@echo "  lint       - Run linter"
	@echo "  lint-fix   - Run linter with auto-fix"
	@echo "  clean      - Clean build artifacts"
	@echo "  run        - Run with sample data"
	@echo "  install    - Install to GOPATH/bin"
	@echo "  fmt        - Format code"
	@echo "  deps       - Download dependencies"
	@echo "  verify     - Verify dependencies"
	@echo "  help       - Show this help message"

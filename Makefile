# Name of the Go binary
BINARY_NAME=tyr

# Go-related variables
GO=go
GOBUILD=$(GO) build
GOTEST=$(GO) test
GOCLEAN=$(GO) clean

# Default target to build the project
all: build

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME)

# Run tests
test:
	$(GOTEST) ./...

# Clean up binary and build artifacts
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Run the app
run: build
	@./$(BINARY_NAME)

.PHONY: all build test clean run
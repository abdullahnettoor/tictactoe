# Variables
BINARY_NAME=tictactoe
GO=go
BUILD_DIR=build
MAIN_FILE=main.go
VERSION=v0.2.0
LDFLAGS=-ldflags "-X github.com/abdullahnettoor/tictactoe/cmd.Version=${VERSION}"

# Colors for terminal output
GREEN=\033[0;32m
NC=\033[0m # No Color
YELLOW=\033[1;33m
RED=\033[0;31m
BLUE=\033[0;34m

# Default target
.DEFAULT_GOAL := help

# Build the application
build:
	@echo "$(YELLOW)Building $(BINARY_NAME)...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@$(GO) build ${LDFLAGS} -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "$(GREEN)Build successful!$(NC)"

# Run the application
run: build
	@echo "$(YELLOW)Running $(BINARY_NAME)...$(NC)"
	@./$(BUILD_DIR)/$(BINARY_NAME) new

# Game Commands
new: build
	@echo "$(BLUE)Starting new game...$(NC)"
	@./$(BUILD_DIR)/$(BINARY_NAME) new

help-game: build
	@echo "$(BLUE)Showing game help...$(NC)"
	@./$(BUILD_DIR)/$(BINARY_NAME) help

version: build
	@echo "$(BLUE)Showing version...$(NC)"
	@./$(BUILD_DIR)/$(BINARY_NAME) version

# Clean build artifacts
clean:
	@echo "$(YELLOW)Cleaning build artifacts...$(NC)"
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "$(GREEN)Cleanup complete!$(NC)"

# Install dependencies
deps:
	@echo "$(YELLOW)Installing dependencies...$(NC)"
	@$(GO) mod download
	@echo "$(GREEN)Dependencies installed successfully!$(NC)"

# Run tests
test:
	@echo "$(YELLOW)Running tests...$(NC)"
	@$(GO) test ./...
	@echo "$(GREEN)Tests completed!$(NC)"

# Format code
fmt:
	@echo "$(YELLOW)Formatting code...$(NC)"
	@$(GO) fmt ./...
	@echo "$(GREEN)Code formatting complete!$(NC)"

# Install the binary globally
install:
	@echo "$(YELLOW)Installing $(BINARY_NAME) globally...$(NC)"
	@$(GO) install
	@echo "$(GREEN)Installation complete! You can now use '$(BINARY_NAME)' command$(NC)"

# Show help
help:
	@echo "$(YELLOW)Available commands:$(NC)"
	@echo ""
	@echo "$(BLUE)Game Commands:$(NC)"
	@echo "  $(GREEN)make new$(NC)        - Start a new game"
	@echo "  $(GREEN)make help-game$(NC)  - Show game help"
	@echo "  $(GREEN)make version$(NC)    - Show game version"
	@echo ""
	@echo "$(BLUE)Development Commands:$(NC)"
	@echo "  $(GREEN)make build$(NC)      - Build the application"
	@echo "  $(GREEN)make run$(NC)        - Build and run the application"
	@echo "  $(GREEN)make clean$(NC)      - Remove build artifacts"
	@echo "  $(GREEN)make deps$(NC)       - Install dependencies"
	@echo "  $(GREEN)make test$(NC)       - Run tests"
	@echo "  $(GREEN)make fmt$(NC)        - Format code"
	@echo "  $(GREEN)make install$(NC)    - Install binary globally"
	@echo "  $(GREEN)make help$(NC)       - Show this help message"

.PHONY: build run clean deps test fmt install help new help-game version

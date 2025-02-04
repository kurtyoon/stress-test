.PHONY: build clean run test

BINARY_NAME=stress-test
BUILD_DIR=build
CMD_DIR=cmd/loadtest

build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)/main.go

run:
	@if [ -z "$(URL)" ] || [ -z "$(RPS)" ] || [ -z "$(DURATION)" ]; then \
		echo "Usage: make run URL=<target_url> RPS=<requests_per_second> DURATION=<duration_in_seconds>"; \
		exit 1; \
	fi
	@$(BUILD_DIR)/$(BINARY_NAME) -url $(URL) -rps $(RPS) -duration $(DURATION)

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

test:
	@echo "Running tests..."
	@go test ./...

install: build
	@echo "Installing..."
	@cp $(BUILD_DIR)/$(BINARY_NAME) $(GOPATH)/bin/

help:
	@echo "Available commands:"
	@echo "  make build              - Build the application"
	@echo "  make run URL=<url> RPS=<rps> DURATION=<duration> - Run the application"
	@echo "  make clean              - Remove build artifacts"
	@echo "  make test               - Run tests"
	@echo "  make install            - Install to GOPATH/bin"
	@echo ""
	@echo "Example:"
	@echo "  make run URL=https://example.com RPS=10 DURATION=5" 
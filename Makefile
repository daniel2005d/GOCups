BINARY_NAME=gocups
BUILD_DIR=bin
LDFLAGS := -ldflags="-s -w" -trimpath

.PHONY: all clean help build-all build-linux build-windows build-darwin

all: build-linux

build-all: build-linux build-windows build-darwin

build-linux:
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .

build-windows:
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)x64.exe .

build-darwin:
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .

clean:
	rm -rf $(BUILD_DIR)

help:
	@echo "Available targets:"
	@echo "  make build-linux    - Build for Linux (amd64)"
	@echo "  make build-windows  - Build for Windows (amd64)"
	@echo "  make build-darwin   - Build for macOS (arm64 Apple Silicon)"
	@echo "  make build-all      - Build for all three platforms"
	@echo "  make clean          - Remove build artifacts"
BINARY := specula
SRC_DIR := src
BUILD_DIR := build
GO := go
GOFLAGS := -ldflags "-s -w"

.PHONY: all clean fmt

all: fmt build

fmt:
	@$(GO) fmt ./$(SRC_DIR)/...

build: $(BUILD_DIR)/$(BINARY)

$(BUILD_DIR)/$(BINARY): $(wildcard $(SRC_DIR)/*.go)
	@mkdir -p $(BUILD_DIR)
	@echo "Building $(BINARY) → $(BUILD_DIR)/$(BINARY)"
	@$(GO) build $(GOFLAGS) -o $@ ./$(SRC_DIR)

clean:
	@echo "Removing build directory"
	@rm -rf $(BUILD_DIR)

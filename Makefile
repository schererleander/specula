BINARY := specula
SRC_DIR := src
BUILD_DIR := build
PREFIX    ?= /usr/local
GO := go
GOFLAGS := -ldflags "-s -w"

.PHONY: all clean fmt build install uninstall

all: fmt build

fmt:
	@$(GO) fmt ./$(SRC_DIR)/...

build: $(BUILD_DIR)/$(BINARY)

$(BUILD_DIR)/$(BINARY): $(wildcard $(SRC_DIR)/*.go)
	@mkdir -p $(BUILD_DIR)
	@echo "Building $(BINARY) → $(BUILD_DIR)/$(BINARY)"
	@$(GO) build $(GOFLAGS) -o $@ ./$(SRC_DIR)

install: build
	@echo "Installing $(BINARY) to $(PREFIX)/bin"
	@mkdir -p $(DESTDIR)$(PREFIX)/bin
	@install -m 0755 $(BUILD_DIR)/$(BINARY) $(DESTDIR)$(PREFIX)/bin/$(BINARY)

uninstall:
	@echo "Removing $(BINARY) from $(PREFIX)/bin"
	@rm -f $(DESTDIR)$(PREFIX)/bin/$(BINARY)

clean:
	@echo "Removing build directory"
	@rm -rf $(BUILD_DIR)

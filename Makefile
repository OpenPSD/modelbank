GO_CMD             ?= go
GO_DEP_CMD         ?= dep
APP_NAME        ?= modelbank

install-dependencies:
	@echo "install dependencies..."
	$(GO_CMD) get ./...

build: install-dependencies
	@echo "build $(APP_NAME)"
	$(GO_CMD) build


.PHONY: help all build run test test-coverage coverage coverage-func clean swag

APP_NAME=server.exe
MAIN_FILE=cmd/server/main.go
GO=go
COVER_PROFILE=coverage.out
COVER_HTML=coverage.html

help: ## Display this help message
	@echo "Available targets:"
	@echo "  make build          - Build the application"
	@echo "  make run            - Run the application"
	@echo "  make test           - Run all tests"
	@echo "  make test-coverage  - Run tests with coverage report"
	@echo "  make coverage       - Alias for test-coverage"
	@echo "  make coverage-func  - Show coverage by function"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make swag           - Generate swagger documentation"

all: build

build:
	$(GO) build -o $(APP_NAME) $(MAIN_FILE)

run:
	$(GO) run $(MAIN_FILE)

test:
	$(GO) test -v -race -coverprofile=$(COVER_PROFILE) -covermode=atomic ./tests/...

test-coverage: test
	$(GO) tool cover -html=$(COVER_PROFILE) -o $(COVER_HTML)

coverage:
	$(GO) tool cover -html=$(COVER_PROFILE) -o $(COVER_HTML)

coverage-func:
	$(GO) tool cover -func=$(COVER_PROFILE)

clean:
	$(GO) clean
	-del /Q $(APP_NAME) 2>nul
	-del /Q $(COVER_PROFILE) 2>nul
	-del /Q $(COVER_HTML) 2>nul

swag:
	swag init -g $(MAIN_FILE)

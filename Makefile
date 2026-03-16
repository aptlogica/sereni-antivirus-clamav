.PHONY: all build run test clean swag

APP_NAME=server.exe
MAIN_FILE=cmd/server/main.go

all: build

build:
	go build -o $(APP_NAME) $(MAIN_FILE)

run:
	go run $(MAIN_FILE)

test:
	mkdir -p coverage
	go test -cover -coverpkg=./internal/... -coverprofile=coverage/coverage.out ./tests

coverage:
	go tool cover -html=coverage/coverage.out -o coverage/coverage.html

coverage-func:
	go tool cover -func=coverage/coverage.out

clean:
	go clean
	rm -f $(APP_NAME)
	rm -rf coverage

swag:
	swag init -g $(MAIN_FILE)

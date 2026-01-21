.PHONY: all build run test clean swag

APP_NAME=server.exe
MAIN_FILE=cmd/server/main.go

all: build

build:
	go build -o $(APP_NAME) $(MAIN_FILE)

run:
	go run $(MAIN_FILE)

test:
	go test -cover -coverpkg=./internal/... -coverprofile=coverage.out ./tests

coverage:
	go tool cover -html=coverage.out

coverage-func:
	go tool cover -func=coverage.out

clean:
	go clean
	rm -f $(APP_NAME)

swag:
	swag init -g $(MAIN_FILE)

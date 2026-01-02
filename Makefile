.PHONY: all build run test clean swag

APP_NAME=server.exe
MAIN_FILE=cmd/server/main.go

all: build

build:
	go build -o $(APP_NAME) $(MAIN_FILE)

run:
	go run $(MAIN_FILE)

test:
	go test ./...

clean:
	go clean
	rm -f $(APP_NAME)

swag:
	swag init -g $(MAIN_FILE)

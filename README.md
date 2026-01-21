# Sereni Antivirus ClamAV

## Overview

Sereni Antivirus is a microservice for scanning files for malware using ClamAV. It provides REST API endpoints to upload and scan single or multiple files concurrently.

The service loads configuration from environment variables, initializes a ClamAV provider, and exposes HTTP endpoints via Gin framework. File scanning is performed asynchronously for multiple files using goroutines.

External dependencies include ClamAV daemon (clamd) for antivirus scanning.

## Features

- Scan single file for malware
- Scan multiple files concurrently
- RESTful API with JSON responses
- Configurable upload size limits
- Swagger/OpenAPI documentation
- Environment-based configuration

## Project Structure

```
.
├── cmd/server/
│   └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go        # Configuration loading
│   ├── handlers/
│   │   └── scan_handler.go  # HTTP request handlers
│   ├── providers/antivirus/
│   │   ├── clamav/
│   │   │   └── clamav.go    # ClamAV provider implementation
│   │   ├── factory.go       # Provider factory
│   │   └── interfaces/
│   │       └── interfaces.go # Provider interfaces
│   ├── routes/
│   │   └── routes.go        # Route setup
│   └── services/
│       └── antivirus_service.go # Business logic
├── tests/                   # Unit tests
├── docs/                    # Generated Swagger docs
├── .env.example             # Environment variables template
└── README.md
```

## Requirements

- Go 1.24.4 or later
- ClamAV daemon (clamd) running and accessible
- Linux/macOS environment

## Configuration

Copy `.env.example` to `.env` and adjust values as needed.

```bash
cp .env.example .env
```

## Running the Application

1. Ensure ClamAV daemon is running:
   ```bash
   sudo systemctl start clamav-daemon  # or equivalent
   ```

2. Set environment variables:
   ```bash
   export HOST=localhost
   export PORT=8080
   export BASE_URL=http://localhost:8080
   export CLAMAV_ADDRESS=127.0.0.1:3310
   ```

3. Run the application:
   ```bash
   go run cmd/server/main.go
   ```

The server will start on `localhost:8080`.

## API Documentation

API documentation is available via Swagger UI at `${BASE_URL}/swagger/index.html`.

### Endpoints

- `POST /scan` - Scan a single file
- `POST /scan-files` - Scan multiple files

### Example Request

```bash
curl -X POST ${BASE_URL}/scan \
  -F "file=@/path/to/file"
```

### Response

```json
{
  "file_name": "example.txt",
  "clean": true,
  "threat": ""
}
```

## Environment Variables

| Variable              | Required | Default          | Description |
|-----------------------|----------|------------------|-------------|
| HOST                  | No      | localhost       | Server host |
| PORT                  | No      | 8080            | Server port |
| BASE_URL              | No      | http://localhost:8080 | Base URL for API |
| ANTIVIRUS_DRIVER      | No      | clamav          | Antivirus provider |
| CLAMAV_ADDRESS        | No      | 127.0.0.1:3310  | ClamAV daemon address |
| CLAMAV_TIMEOUT_SECONDS| No      | 30              | Scan timeout |
| MAX_UPLOAD_SIZE_MB    | No      | 32              | Max upload size in MB |

For production, set `HOST=0.0.0.0` and adjust `CLAMAV_ADDRESS` to the production ClamAV instance. Never commit secrets or production-specific values.

## Common Commands

- Build the application:
  ```bash
  go build -o bin/server cmd/server/main.go
  ```

- Run tests with coverage:
  ```bash
  go test -cover -coverpkg=./internal/... -coverprofile=coverage.out ./tests
  ```

- View coverage summary:
  ```bash
  go tool cover -func=coverage.out
  ```

- View coverage report in browser:
  ```bash
  go tool cover -html=coverage.out
  ```

- Generate Swagger docs:
  ```bash
  go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/server/main.go
  ```

- Lint code:
  ```bash
  go vet ./...

# Build stage
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git make

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy && go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o server.exe ./cmd/server

# Prepare optional config directory (always exists, may be empty)
RUN mkdir -p /app/config && \
    (cp .env* /app/config/ 2>/dev/null || true)

# Runtime stage
FROM alpine:latest

WORKDIR /app

# Install ClamAV and dependencies
RUN apk add --no-cache \
    clamav \
    clamav-daemon \
    clamav-libunrar \
    ca-certificates \
    && mkdir -p /var/lib/clamav \
    && chown -R clamav:clamav /var/lib/clamav

# Copy the binary from builder
COPY --from=builder /app/server.exe .

# Copy config directory (may be empty if no .env files exist)
COPY --from=builder /app/config ./

# Create necessary directories
RUN mkdir -p /tmp/uploads

# Expose the application port (adjust if needed)
EXPOSE 6060

# Start ClamAV daemon and the application
CMD ["/bin/sh", "-c", "freshclam --daemon-notify=/tmp/clamd.sock & clamd & sleep 5 && ./server.exe"]

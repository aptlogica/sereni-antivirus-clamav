<h1 align="center">sereni-antivirus-clamav - Cloud-Native Antivirus Microservice</h1>

<p align="center">Enterprise-grade antivirus microservice and open source antivirus service with ClamAV integration. A comprehensive malware scanning API and backend virus scanner providing advanced file scanning, REST APIs, and seamless integration with modern security infrastructure.</p>

<p align="center">
<a href="LICENSE"><img src="https://img.shields.io/badge/Version-1.0.0-blue?style=for-the-badge" alt="Version"></a>
<a href="https://golang.org"><img src="https://img.shields.io/badge/Go-1.26.2-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Version"></a>
<a href="https://www.clamav.net/"><img src="https://img.shields.io/badge/ClamAV-Powered-4B275F?style=for-the-badge&logo=clamav&logoColor=white" alt="ClamAV"></a>
<a href="https://gin-gonic.com/"><img src="https://img.shields.io/badge/Gin-Framework-008ECF?style=for-the-badge&logo=gin&logoColor=white" alt="Gin"></a>
<a href="https://www.docker.com/"><img src="https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker"></a>
<a href="https://swagger.io/"><img src="https://img.shields.io/badge/Swagger-Documented-85EA2D?style=for-the-badge&logo=swagger&logoColor=black" alt="Swagger"></a>
</p>

<p align="center">
<a href="https://github.com/aptlogica/sereni-antivirus-clamav/actions/workflows/ci.yml"><img src="https://github.com/aptlogica/sereni-antivirus-clamav/actions/workflows/ci.yml/badge.svg" alt="CI"></a>
<a href="https://github.com/aptlogica/sereni-antivirus-clamav/actions/workflows/codeql.yml"><img src="https://github.com/aptlogica/sereni-antivirus-clamav/actions/workflows/codeql.yml/badge.svg" alt="CodeQL"></a>
<a href="https://app.fossa.com/projects/git%2Bgithub.com%2Faptlogica%2Fsereni-antivirus-clamav?ref=badge_shield"><img src="https://app.fossa.com/api/projects/git%2Bgithub.com%2Faptlogica%2Fsereni-antivirus-clamav.svg?type=shield" alt="FOSSA Status"></a>
<a href="https://app.fossa.com/projects/git%2Bgithub.com%2Faptlogica%2Fsereni-antivirus-clamav?ref=badge_shield&issueType=security"><img src="https://app.fossa.com/api/projects/git%2Bgithub.com%2Faptlogica%2Fsereni-antivirus-clamav.svg?type=shield&issueType=security" alt="FOSSA Security"></a>
<a href="https://sonarcloud.io/dashboard?id=aptlogica_sereni-antivirus-clamav"><img src="https://sonarcloud.io/api/project_badges/measure?project=aptlogica_sereni-antivirus-clamav&metric=alert_status" alt="Quality Gate"></a>
<a href="https://sonarcloud.io/dashboard?id=aptlogica_sereni-antivirus-clamav"><img src="https://sonarcloud.io/api/project_badges/measure?project=aptlogica_sereni-antivirus-clamav&metric=coverage" alt="Coverage"></a>
<a href="https://sonarcloud.io/dashboard?id=aptlogica_sereni-antivirus-clamav"><img src="https://sonarcloud.io/api/project_badges/measure?project=aptlogica_sereni-antivirus-clamav&metric=security_rating" alt="Security"></a>
</p>

<p align="center">
<a href="LICENSE"><img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg" alt="License: Apache 2.0"></a>
</p>


## Overview

**Sereni Antivirus ClamAV**  is an open-source antivirus and malware scanning solution built on ClamAV, designed to detect and prevent viruses, ransomware, and other threats across files, applications, and servers. It allows seamless integration into APIs, workflows, and cloud environments, enabling secure file uploads and real-time threat detection with a scalable and cost-effective approach.

## Key Features

- **Advanced ClamAV Integration**: Latest malware definitions with real-time updates
- **Comprehensive REST API**: RESTful endpoints for file scanning and status monitoring
- **Multi-Tenant Support**: Isolated scanning contexts for different organizations
- **Secure File Processing**: Encrypted file handling with automatic cleanup
- **Performance Optimization**: Concurrent scanning with intelligent resource management
- **File Upload Security**: Complete antivirus file scanning with secure file upload capabilities
- **Cloud-Native Security**: Kubernetes deployment with security best practices

## Architecture
- Go 1.26.2, idiomatic design
- Modular, testable codebase

## Installation
```sh
go get github.com/aptlogica/sereni-antivirus-clamav
```

## Configuration
See `.env.example` for environment variables and configuration options.

## Documentation

### Deployment & Operations
- **[PRODUCTION.md](PRODUCTION.md)** - Production deployment guide with environment configuration, resource constraints, health checks, and persistent volumes
- **[DEPLOYMENT.md](DEPLOYMENT.md)** - Cloud deployment guides for Docker Compose, Kubernetes, AWS ECS, Google Cloud Run, and Azure

These documents cover:
- ✅ Complete environmental variable reference
- ✅ Resource sizing and constraints (CPU, memory, disk)
- ✅ Health check mechanisms and monitoring
- ✅ Persistent volume setup for virus signatures (avoids re-downloading)
- ✅ Deployment examples for all major platforms

## Quick Start

```go
package main

import (
    "context"
    "log"
    "os"
    
    "github.com/aptlogica/sereni-antivirus-clamav/pkg/client"
    "github.com/aptlogica/sereni-antivirus-clamav/pkg/config"
)

func main() {
    // Initialize configuration
    cfg := config.New()
    cfg.ClamAVHost = "localhost"
    cfg.ClamAVPort = 3310
    cfg.MaxFileSize = "100MB"
    
    // Create antivirus client
    client, err := client.New(cfg)
    if err != nil {
        log.Fatal("Failed to create client:", err)
    }
    defer client.Close()
    
    // Scan a file
    file, err := os.Open("document.pdf")
    if err != nil {
        log.Fatal("Failed to open file:", err)
    }
    defer file.Close()
    
    ctx := context.Background()
    result, err := client.ScanFile(ctx, file)
    if err != nil {
        log.Fatal("Failed to scan file:", err)
    }
    
    if result.IsClean {
        log.Println("File is clean")
    } else {
        log.Printf("Threat detected: %s", result.ThreatName)
    }
}
```

## Development

### Local Setup
```bash
# Clone the repository
git clone https://github.com/aptlogica/sereni-antivirus-clamav.git
cd sereni-antivirus-clamav

# Install dependencies
go mod download

# Install ClamAV (Ubuntu/Debian)
sudo apt-get update
sudo apt-get install clamav clamav-daemon

# Update virus definitions
sudo freshclam

# Set up environment
cp .env.example .env
# Configure ClamAV connection in .env

# Start ClamAV daemon
sudo systemctl start clamav-daemon

# Start development server
go run ./cmd/server
```

### Environment Configuration
```bash
CLAMAV_HOST=localhost
CLAMAV_PORT=3310
MAX_FILE_SIZE=100MB
SCAN_TIMEOUT=30s
TEMP_DIR=/tmp/scans
PORT=8080
LOG_LEVEL=debug
```

### Docker Development
```bash
# Start ClamAV with Docker
docker run -d \
  --name clamav \
  -p 3310:3310 \
  clamav/clamav:stable

# Wait for virus definitions to update
docker logs -f clamav

# Run the service
go run ./cmd/server
```

### Testing Malware Detection
```bash
# Download EICAR test file (safe malware test)
wget https://secure.eicar.org/eicar.com.txt

# Test scanning via API
curl -X POST -F "file=@eicar.com.txt" http://localhost:8080/scan
```

## Testing
- Run `go test ./...` to execute unit tests

## Security
See [SECURITY.md](SECURITY.md) for reporting vulnerabilities.

## License
Apache License 2.0. Copyright (c) 2026 Aptlogica Technologies.
Licensed under the Apache License, Version 2.0. See the [LICENSE](LICENSE) file for details.

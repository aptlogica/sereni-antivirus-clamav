# Sereni Antivirus - ClamAV Microservice

> A production-ready, scalable REST API microservice for real-time malware scanning and threat detection using ClamAV. Deploy as a standalone service or integrate into any application ecosystem.

[![Go Version](https://img.shields.io/badge/Go-1.24.4+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat-square&logo=docker)](https://www.docker.com/)

## 📋 Table of Contents

- [Overview](#overview)
- [Key Features](#key-features)
- [Use Cases](#use-cases)
- [Quick Start](#quick-start)
- [Installation](#installation)
- [Configuration](#configuration)
- [API Documentation](#api-documentation)
- [Integration Guide](#integration-guide)
- [Docker Deployment](#docker-deployment)
- [Architecture](#architecture)
- [Development](#development)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)

## Overview

**Sereni Antivirus** is a high-performance microservice that provides real-time malware detection and file scanning capabilities. It leverages ClamAV (Clam AntiVirus), a leading open-source antivirus engine, to scan files for viruses, trojans, and malicious software.

The service is designed with scalability and flexibility in mind:

- **Microservice Architecture**: Deploy independently or as part of a larger system
- **Language-Agnostic**: Can be integrated with applications written in any language or framework (Node.js, Python, Java, .NET, etc.)
- **Concurrent Scanning**: Process multiple files simultaneously using efficient worker pools
- **RESTful API**: Simple HTTP/JSON interface for easy integration
- **Production-Ready**: Docker support, health checks, comprehensive logging, and error handling

### Why Use Sereni Antivirus?

- **Zero Configuration Complexity**: Environment-based configuration - no complex setup files
- **Battle-Tested Foundation**: Built on ClamAV, trusted by millions globally for threat detection
- **High Performance**: Concurrent file scanning with automatic worker pool optimization
- **Easy Integration**: Simple REST API that works with any tech stack
- **Docker-Native**: One-command deployment with docker-compose
- **Complete Documentation**: Swagger UI, API docs, and integration guides included

## Key Features

✅ **Single File Scanning**
- Upload and scan individual files for malware
- Real-time threat detection with detailed scan results
- Support for unlimited file types

✅ **Batch File Processing**
- Scan multiple files concurrently
- Automatic worker pool based on CPU cores
- Efficient resource utilization

✅ **RESTful API**
- Clean, intuitive HTTP/JSON interface
- Multipart form-data support for file uploads
- Structured JSON responses with threat intelligence

✅ **High Performance**
- Non-blocking concurrent scanning using goroutines
- Automatic CPU-aware worker pool sizing
- Configurable request size limits (up to 500MB+)

✅ **Configuration & Control**
- Environment-based configuration
- Customizable upload size limits
- ClamAV timeout configuration
- CORS support for cross-origin requests

✅ **Developer Friendly**
- Swagger/OpenAPI documentation with interactive UI
- Comprehensive error messages
- Health check endpoints
- Detailed logging

✅ **Malware Intelligence**
- Detection of viruses, trojans, worms, and rootkits
- Signature-based detection with regular updates
- Behavioral threat analysis
- Detailed scanning reports

✅ **Enterprise Features**
- Docker and Docker Compose support
- Health checks and readiness probes
- Resource limits and timeout handling
- CORS policy configuration

## Use Cases

### 1. **File Upload Security**
Protect your web application's file upload functionality by automatically scanning all user-uploaded files before storing them.

```
User Application → Sereni Antivirus → ClamAV → Clean/Threat Verdict
```

### 2. **Email Security & Attachment Scanning**
Integrate with email servers to scan attachments for threats before delivery.

### 3. **Document Management Systems**
Ensure all documents in your DMS are clean before archival or storage.

### 4. **API Proxy/Gateway**
Deploy as a middleware service to add antivirus scanning to any API endpoint.

### 5. **Batch Processing**
Scan entire directories or downloaded files in batch operations.

### 6. **Cloud Storage Protection**
Monitor and scan files uploaded to cloud storage platforms (S3, GCS, etc.).

### 7. **Containerized Environments**
Deploy in Kubernetes or Docker Swarm for resilient threat detection.

## Quick Start

### Prerequisites
- **Docker & Docker Compose** (recommended) or **Go 1.24.4+**
- **2GB+ RAM**, **1+ CPU cores**
- **Internet connection** (for ClamAV virus database updates)

### 30-Second Setup (Docker)

```bash
# Clone the repository
git clone https://github.com/aptlogica/sereni-antivirus-clamav.git
cd sereni-antivirus-clamav

# Start services
docker-compose up -d

# Verify installation
curl http://localhost:8084/health
```

The service will be available at `http://localhost:8084` with API documentation at `http://localhost:8084/swagger/index.html`.

## Installation

### Option 1: Docker Compose (Recommended)

```bash
# 1. Clone repository
git clone https://github.com/aptlogica/sereni-antivirus-clamav.git
cd sereni-antivirus-clamav

# 2. Copy environment template
cp .env.example .env

# 3. Start all services
docker-compose up -d

# 4. Wait for ClamAV to initialize (first startup takes 2-3 minutes)
docker-compose logs -f clamav

# 5. Test the service
curl -X POST http://localhost:8084/scan \
  -F "file=@/path/to/testfile"
```

### Option 2: Manual Setup (Linux/macOS)

#### Step 1: Install ClamAV

**Ubuntu/Debian:**
```bash
sudo apt-get update
sudo apt-get install clamav clamav-daemon

# Start ClamAV daemon
sudo systemctl start clamav-daemon
sudo systemctl enable clamav-daemon
```

**macOS (via Homebrew):**
```bash
brew install clamav

# Update virus database
freshclam

# Start ClamAV
clamd
```

**Other Systems**: [ClamAV Installation Guide](https://docs.clamav.net/manual/Installing.html)

#### Step 2: Set Up Go Project

```bash
# 1. Clone and navigate
git clone https://github.com/aptlogica/sereni-antivirus-clamav.git
cd sereni-antivirus-clamav

# 2. Copy environment file
cp .env.example .env

# 3. Edit .env with your settings
nano .env

# 4. Download dependencies
go mod download

# 5. Run application
go run cmd/server/main.go
```

### Option 3: Build Docker Image Manually

```bash
# Build image
docker build -t sereni-antivirus:latest .

# Run container
docker run -d \
  --name sereni-antivirus \
  -p 8084:8084 \
  -e HOST=0.0.0.0 \
  -e PORT=8084 \
  -e CLAMAV_ADDRESS=clamav-host:3310 \
  sereni-antivirus:latest
```

## Configuration

### Environment Variables

Create `.env` file in project root to customize service behavior:

```dotenv
# === Server Configuration ===
HOST=0.0.0.0                           # Bind address (use 0.0.0.0 for Docker)
PORT=8084                              # Service port
BASE_URL=http://localhost:8084         # Base URL for Swagger documentation

# === CORS Configuration ===
ALLOWED_ORIGINS=*                      # CORS allowed origins (* = all)

# === Antivirus Configuration ===
ANTIVIRUS_DRIVER=clamav                # Antivirus backend (currently: clamav)

# === ClamAV Configuration ===
CLAMAV_ADDRESS=127.0.0.1:3310         # ClamAV daemon address
CLAMAV_TIMEOUT_SECONDS=30              # Scan timeout in seconds

# === Upload Configuration ===
MAX_UPLOAD_SIZE_MB=32                  # Maximum file upload size in MB
```

### Default Configuration

If `.env` file is not present, these defaults will be used:
- HOST: localhost
- PORT: 8084
- CLAMAV_ADDRESS: 127.0.0.1:3310
- MAX_UPLOAD_SIZE_MB: 32MB

### Configuration in Docker

For Docker Compose, edit `.env` file before running `docker-compose up`.

For Kubernetes, use ConfigMaps and Secrets:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: sereni-antivirus-config
data:
  HOST: "0.0.0.0"
  PORT: "8084"
  CLAMAV_ADDRESS: "clamav-service:3310"
  MAX_UPLOAD_SIZE_MB: "64"
```

## API Documentation

### Interactive API Docs

Once the service is running, visit: **http://localhost:8084/swagger/index.html**

Full Swagger/OpenAPI documentation is available with the running service.

### Endpoints

#### 1. Health Check
```http
GET /health
```

**Response:**
```json
{
  "status": "ok"
}
```

#### 2. Scan Single File
```http
POST /scan
Content-Type: multipart/form-data

file: <binary file data>
```

**Response:**
```json
{
  "filename": "document.pdf",
  "size": 1024000,
  "clean": true,
  "threat": "",
  "timestamp": "2024-03-11T10:30:00Z"
}
```

#### 3. Scan Multiple Files
```http
POST /scan-files
Content-Type: multipart/form-data

files: <file1>, <file2>, <file3>...
```

**Response:**
```json
[
  {
    "filename": "file1.pdf",
    "size": 512000,
    "clean": true,
    "threat": "",
    "timestamp": "2024-03-11T10:30:00Z"
  },
  {
    "filename": "file2.exe",
    "size": 256000,
    "clean": false,
    "threat": "Win.Trojan.Generic-12345",
    "timestamp": "2024-03-11T10:30:01Z"
  }
]
```

### API Response Format

| Field | Type | Description |
|-------|------|-------------|
| `filename` | string | Name of the scanned file |
| `size` | integer | File size in bytes |
| `clean` | boolean | True if no threats detected |
| `threat` | string | Threat name if detected (empty if clean) |
| `timestamp` | string | RFC3339 scan timestamp |

### Error Responses

**400 - Bad Request**
```json
{
  "error": "No file uploaded"
}
```

**413 - Payload Too Large**
```json
{
  "error": "Request body too large"
}
```

**500 - Server Error**
```json
{
  "error": "ClamAV connection failed"
}
```

## Integration Guide

### Node.js / Express.js

```javascript
const express = require('express');
const multer = require('multer');
const axios = require('axios');

const app = express();
const upload = multer({ storage: multer.memoryStorage() });

const ANTIVIRUS_URL = 'http://localhost:8084';

// Scan single file endpoint
app.post('/upload', upload.single('file'), async (req, res) => {
  try {
    const formData = new FormData();
    formData.append('file', req.file.buffer, req.file.originalname);

    const response = await axios.post(`${ANTIVIRUS_URL}/scan`, formData, {
      headers: formData.getHeaders()
    });

    if (response.data.clean) {
      // Save file to database
      res.json({ success: true, message: 'File is clean' });
    } else {
      res.status(400).json({ 
        error: 'Malware detected', 
        threat: response.data.threat 
      });
    }
  } catch (error) {
    res.status(500).json({ error: 'Scan failed' });
  }
});

app.listen(3000, () => console.log('Server running on port 3000'));
```

### Python / Flask

```python
from flask import Flask, request, jsonify
import requests

app = Flask(__name__)
ANTIVIRUS_URL = 'http://localhost:8084'

@app.route('/upload', methods=['POST'])
def upload_file():
    file = request.files['file']
    
    # Prepare multipart form
    files = {'file': (file.filename, file.stream)}
    
    # Send to antivirus service
    response = requests.post(f'{ANTIVIRUS_URL}/scan', files=files)
    scan_result = response.json()
    
    if scan_result['clean']:
        # Process file
        return jsonify({'success': True, 'message': 'File is clean'})
    else:
        return jsonify({
            'error': 'Malware detected',
            'threat': scan_result['threat']
        }), 400

if __name__ == '__main__':
    app.run(debug=True)
```

### Java / Spring Boot

```java
@RestController
@RequestMapping("/upload")
public class FileUploadController {
    
    private final RestTemplate restTemplate = new RestTemplate();
    private static final String ANTIVIRUS_URL = "http://localhost:8084/scan";
    
    @PostMapping
    public ResponseEntity<?> uploadFile(@RequestParam("file") MultipartFile file) {
        try {
            // Create multipart request
            MultiValueMap<String, Object> body = new LinkedMultiValueMap<>();
            body.add("file", file.getResource());
            
            HttpHeaders headers = new HttpHeaders();
            headers.setContentType(MediaType.MULTIPART_FORM_DATA);
            
            HttpEntity<MultiValueMap<String, Object>> requestEntity = 
                new HttpEntity<>(body, headers);
            
            // Call antivirus service
            ResponseEntity<ScanResult> response = restTemplate.exchange(
                ANTIVIRUS_URL,
                HttpMethod.POST,
                requestEntity,
                ScanResult.class
            );
            
            if (response.getBody().isClean()) {
                return ResponseEntity.ok("File is clean");
            } else {
                return ResponseEntity.status(400)
                    .body("Threat detected: " + response.getBody().getThreat());
            }
        } catch (Exception e) {
            return ResponseEntity.status(500).body("Scan failed: " + e.getMessage());
        }
    }
}

class ScanResult {
    private String filename;
    private boolean clean;
    private String threat;
    // Getters and setters...
}
```

### PHP / Laravel

```php
<?php
namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class FileUploadController extends Controller
{
    public function upload(Request $request)
    {
        $file = $request->file('file');
        
        try {
            $response = Http::attach(
                'file',
                $file->get(),
                $file->getClientOriginalName()
            )
            ->post('http://localhost:8084/scan');
            
            $result = $response->json();
            
            if ($result['clean']) {
                // Save file
                return response()->json(['success' => true]);
            } else {
                return response()->json([
                    'error' => 'Malware detected',
                    'threat' => $result['threat']
                ], 400);
            }
        } catch (\Exception $e) {
            return response()->json(['error' => 'Scan failed'], 500);
        }
    }
}
```

### .NET / C#

```csharp
using System;
using System.Net.Http;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Http;

[ApiController]
[Route("api/[controller]")]
public class FileUploadController : ControllerBase
{
    private readonly HttpClient _httpClient;
    private const string AntivirusUrl = "http://localhost:8084/scan";

    public FileUploadController(HttpClient httpClient)
    {
        _httpClient = httpClient;
    }

    [HttpPost("upload")]
    public async Task<IActionResult> UploadFile(IFormFile file)
    {
        try
        {
            using var content = new MultipartFormDataContent();
            using var stream = file.OpenReadStream();
            content.Add(new StreamContent(stream), "file", file.FileName);

            var response = await _httpClient.PostAsync(AntivirusUrl, content);
            var result = await response.Content.ReadAsAsync<ScanResult>();

            if (result.Clean)
            {
                return Ok(new { message = "File is clean" });
            }
            else
            {
                return BadRequest(new { 
                    error = "Malware detected", 
                    threat = result.Threat 
                });
            }
        }
        catch (Exception ex)
        {
            return StatusCode(500, new { error = "Scan failed: " + ex.Message });
        }
    }
}

public class ScanResult
{
    public string Filename { get; set; }
    public int Size { get; set; }
    public bool Clean { get; set; }
    public string Threat { get; set; }
    public DateTime Timestamp { get; set; }
}
```

## Docker Deployment

### Using Docker Compose (Recommended)

1. **Start Services:**
   ```bash
   docker-compose up -d
   ```

2. **Monitor Startup:**
   ```bash
   docker-compose logs -f clamav
   # Wait for: "Listening on [::]:3310"
   ```

3. **Process Management:**
   ```bash
   # Stop services
   docker-compose down

   # View logs
   docker-compose logs sereni-antivirus

   # Rebuild after code changes
   docker-compose up -d --build
   ```

### Docker Compose Configuration

See [docker-compose.yml](docker-compose.yml) for detailed setup including:
- ClamAV service with automatic virus database updates
- Antivirus service with health checks
- Volume mounting for scanning local directories
- Resource limits and environment configuration

### Kubernetes Deployment

Example deployment manifest:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: sereni-antivirus
spec:
  replicas: 2
  selector:
    matchLabels:
      app: sereni-antivirus
  template:
    metadata:
      labels:
        app: sereni-antivirus
    spec:
      containers:
      - name: antivirus
        image: sereni-antivirus:latest
        ports:
        - containerPort: 8084
        env:
        - name: HOST
          value: "0.0.0.0"
        - name: PORT
          value: "8084"
        - name: CLAMAV_ADDRESS
          value: "clamav-service:3310"
        - name: MAX_UPLOAD_SIZE_MB
          value: "64"
        resources:
          requests:
            memory: "512Mi"
            cpu: "500m"
          limits:
            memory: "1Gi"
            cpu: "1000m"
        livenessProbe:
          httpGet:
            path: /health
            port: 8084
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /health
            port: 8084
          initialDelaySeconds: 10
          periodSeconds: 5

---
apiVersion: v1
kind: Service
metadata:
  name: sereni-antivirus
spec:
  selector:
    app: sereni-antivirus
  type: ClusterIP
  ports:
  - protocol: TCP
    port: 8084
    targetPort: 8084
```

## Architecture

### System Design

```
┌─────────────────────────────────────────────────────────────┐
│                     Client Applications                      │
│          (Web, Mobile, Desktop, CLI, Other Services)         │
└────────────────────┬────────────────────────────────────────┘
                     │ HTTP/REST
                     │
┌────────────────────▼────────────────────────────────────────┐
│              Sereni Antivirus Service                        │
│  ┌──────────────  Gin Web Framework  ──────────────┐       │
│  │  HTTP Handlers │ Routes │ CORS │ Error Handler └│       │
│  └──────────────────────┬─────────────────────────────┘       │
│                         │                                     │
│              ┌──────────▼──────────┐                          │
│              │  Config Manager     │                          │
│              │ (Env Variables)     │                          │
│              └─────────────────────┘                          │
│                         │                                     │
│              ┌──────────▼──────────┐                          │
│              │  Antivirus Service  │                          │
│              │ (Business Logic)    │                          │
│              └──────────┬──────────┘                          │
│                         │                                     │
│         ┌───────────────┴───────────────┐                    │
│         │                               │                    │
│    ┌────▼────┐                  ┌──────▼─────┐              │
│    │  Single │                  │  Multi     │              │
│    │  File   │                  │  File      │              │
│    │  Scan   │                  │  Scan      │              │
│    └────┬────┘                  └──────┬─────┘              │
│         │                              │                    │
│         └──────────────┬───────────────┘                    │
│                        │                                    │
│              ┌─────────▼─────────┐                          │
│              │  ClamAV Provider  │                          │
│              │ (Adapter Pattern) │                          │
│              └─────────┬─────────┘                          │
└────────────────────────┼──────────────────────────────────┘
                         │ TCP
                         │
        ┌────────────────▼────────────────┐
        │   ClamAV Daemon (clamd)         │
        │  Virus Database Updates         │
        │  (Signature-based Detection)    │
        └─────────────────────────────────┘
```

### Component Responsibilities

| Component | Responsibility |
|-----------|-----------------|
| **Handler** | HTTP request/response management |
| **Service** | Business logic & synchronization |
| **Provider** | ClamAV integration & adaptation |
| **Config** | Environment-based configuration |
| **ClamAV** | Actual malware detection engine |

## Development

### Project Structure

```
.
├── cmd/server/
│   └── main.go                  # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── handlers/
│   │   └── scan_handler.go      # HTTP handlers
│   ├── providers/antivirus/
│   │   ├── clamav/
│   │   │   └── clamav.go        # ClamAV implementation
│   │   ├── factory.go           # Provider factory
│   │   └── interfaces/
│   │       └── interfaces.go    # Interfaces
│   ├── routes/
│   │   └── routes.go            # Route definitions
│   └── services/
│       └── antivirus_service.go # Service layer
├── tests/                       # Unit tests
├── docs/                        # Swagger documentation
├── docker-compose.yml           # Docker Compose setup
├── Dockerfile                   # Docker image build
├── Makefile                     # Build automation
├── go.mod                       # Go module dependencies
├── go.sum                       # Dependency checksums
└── README.md                    # This file
```

### Setup Development Environment

```bash
# Clone repository
git clone https://github.com/aptlogica/sereni-antivirus-clamav.git
cd sereni-antivirus-clamav

# Install Go (https://golang.org/doc/install)
# For macOS:
brew install go

# Verify installation
go version

# Download dependencies
go mod download

# Build the project
go build -o antivirus-service cmd/server/main.go

# Run tests
go test ./...

# Run with live reload (requires air: go install github.com/cosmtrek/air@latest)
air
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test -run TestScanFile ./tests

# Verbose output
go test -v ./...
```

### Building & Releasing

```bash
# Build release binary
go build -o antivirus-service \
  -ldflags="-X main.Version=1.0.0" \
  cmd/server/main.go

# Build docker image
docker build -t sereni-antivirus:1.0.0 .

# Push to registry
docker push your-registry/sereni-antivirus:1.0.0
```

## Troubleshooting

### Common Issues

#### 1. ClamAV Connection Failed

**Error:** `Failed to connect to ClamAV daemon`

**Solutions:**
```bash
# Check if ClamAV is running
docker-compose ps clamav
# or
ps aux | grep clamd

# Verify ClamAV address
telnet 127.0.0.1 3310

# Check ClamAV logs
docker-compose logs clamav

# Restart ClamAV
docker-compose restart clamav
```

#### 2. Port Already in Use

**Error:** `bind: address already in use`

**Solutions:**
```bash
# Find and kill process using port 8084
lsof -i :8084
kill -9 <PID>

# Or change port in .env
echo "PORT=8085" >> .env
```

#### 3. Request Body Too Large

**Error:** `413 Payload Too Large`

**Solutions:**
```bash
# Increase max upload size in .env
MAX_UPLOAD_SIZE_MB=128

# Or modify docker-compose.yml ClamAV settings
# MaxFileSize=500M
# StreamMaxLength=500M
```

#### 4. Slow Scanning Performance

**Check:**
```bash
# Monitor system resources
docker stats sereni-antivirus

# Check ClamAV virus database size
docker exec clamav du -sh /var/lib/clamav/

# Update virus database
docker exec clamav freshclam
```

#### 5. Service Not Starting

**Debugging:**
```bash
# Check logs
docker-compose logs --tail=50 sereni-antivirus

# Check health
curl http://localhost:8084/health

# Verify environment
docker-compose config
```

### Health Checks

```bash
# Basic health check
curl -i http://localhost:8084/health

# Test single file scan
curl -X POST http://localhost:8084/scan \
  -F "file=@testfile.txt" \
  -H "Accept: application/json"

# Test with eicar test file (safe test virus signature)
curl -X POST http://localhost:8084/scan \
  -F "file=@eicar.com" \
  --output - | jq .
```

### Logs and Debugging

```bash
# View service logs
docker-compose logs sereni-antivirus

# Follow logs in real-time
docker-compose logs -f sereni-antivirus

# View ClamAV logs
docker-compose logs -f clamav

# Enable debug mode (add to .env)
# DEBUG=true (if implemented)

# Check network connectivity
docker exec sereni-antivirus curl -v clamav:3310
```

## Contributing

We welcome contributions! Please follow these guidelines:

### Getting Started

1. Fork the repository
2. Create feature branch: `git checkout -b feature/amazing-feature`
3. Follow code style conventions
4. Write tests for new functionality
5. Update documentation

### Code Standards

- Follow Go idioms and best practices
- Run `go fmt` before committing
- Use meaningful variable names
- Add comments for exported functions
- Write unit tests for new features

### Commit Guidelines

```bash
# Good commit message
git commit -m "feat: add support for batch scanning with progress reporting"

# Format: type(scope): description
# Types: feat, fix, docs, style, refactor, test, chore
```

### Pull Request Process

1. Update CHANGELOG.md with changes
2. Update README if adding new features
3. Ensure all tests pass: `go test ./...`
4. Request code review from maintainers
5. Address feedback promptly

## Project Status & Roadmap

### Current Features (v1.0)
- ✅ Single file scanning
- ✅ Batch file scanning
- ✅ RESTful API
- ✅ Docker deployment
- ✅ Swagger documentation
- ✅ CORS support

### Planned Features
- 🔄 Progress tracking for batch operations
- 🔄 File quarantine system
- 🔄 Multiple antivirus engines support (YARA, etc.)
- 🔄 Advanced threat intelligence integration
- 🔄 Web UI dashboard
- 🔄 Metrics and monitoring (Prometheus)
- 🔄 Database integration for scan history

## FAQ

**Q: Can I use Sereni Antivirus in production?**
A: Yes! It's designed for production use. It includes health checks, error handling, and Docker support.

**Q: What's the maximum file size I can scan?**
A: Default is 32MB, configurable up to 500MB+ via `MAX_UPLOAD_SIZE_MB`.

**Q: Does it work with any programming language?**
A: Yes! Since it's a REST API, it works with any language. See Integration Guide above.

**Q: How often is the virus database updated?**
A: By default every 2 hours (12 times daily). Configurable via docker-compose.yml.

**Q: Can I use custom antivirus engines?**
A: Yes! The architecture supports multiple providers via the factory pattern. Implement the `interfaces.go` interface.

**Q: Is there a web UI?**
A: Currently, there's Swagger UI for API docs. A full dashboard is planned for future releases.

**Q: How do I scale this for high throughput?**
A: Deploy multiple instances behind a load balancer and use a shared ClamAV service.

## License

This project is licensed under the MIT License - see [LICENSE](LICENSE) file for details.

## Support & Community

- **Report Issues**: GitHub Issues
- **Documentation**: [Docs](docs/)
- **ClamAV Documentation**: [ClamAV Docs](https://docs.clamav.net/)

## Acknowledgments

- Built with [Gin Web Framework](https://github.com/gin-gonic/gin)
- Powered by [ClamAV](https://www.clamav.net/)
- Go ClamAV Client: [go-clamd](https://github.com/dutchcoders/go-clamd)

---

**Made with ❤️ for secure file operations**

### API Response

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
| PORT                  | No      | 8084            | Server port |
| BASE_URL              | No      | http://localhost:8084 | Base URL for API |
| ALLOWED_ORIGINS       | No      | *               | CORS allowed origins (comma-separated or *) |
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

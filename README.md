# Sereni Antivirus - ClamAV Microservice

> A production-ready, scalable REST API microservice for real-time malware scanning and threat detection using ClamAV. Deploy as a standalone service or integrate into any application ecosystem.

[![Version](https://img.shields.io/badge/Version-1.0.0-blue.svg)](https://github.com/aptlogica/sereni-antivirus-clamav)
[![Go Version](https://img.shields.io/badge/Go-1.24.4+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat-square&logo=docker)](https://www.docker.com/)
[![ClamAV](https://img.shields.io/badge/ClamAV-Powered-red.svg)](https://www.clamav.net/)

## 📋 Table of Contents

- [Overview](#overview)
- [Key Features](#key-features)
- [Why Choose Sereni Antivirus?](#why-choose-sereni-antivirus)
- [Use Cases](#use-cases)
- [Quick Start](#quick-start)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Integration Guide](#integration-guide)
- [Architecture](#architecture)
- [Development](#development)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [FAQ](#faq)
- [License](#license)

## Overview

**Sereni Antivirus** is a high-performance, language-agnostic microservice that provides real-time malware detection and file scanning capabilities. It leverages ClamAV (Clam AntiVirus), a leading open-source antivirus engine, to scan files for viruses, trojans, worms, and malicious software. Whether you're building a file-sharing platform, document management system, or email protection service, Sereni Antivirus provides enterprise-grade threat detection out of the box.

The service is designed with scalability, flexibility, and ease of integration in mind:

- **Microservice Architecture**: Deploy independently or as part of a larger system without coupling
- **Language-Agnostic**: Integrates seamlessly with applications in any language or framework (Node.js, Python, Java, .NET, Go, PHP, Ruby, etc.)
- **Concurrent Scanning**: Process multiple files simultaneously with efficient goroutine-based worker pools
- **RESTful API**: Simple, intuitive HTTP/JSON interface for quick integration
- **Production-Ready**: Docker support, health checks, comprehensive logging, error handling, and resource management

### Why Choose Sereni Antivirus?

- **Zero Configuration Complexity**: Sensible defaults with environment-based configuration - deploy in minutes
- **Battle-Tested Foundation**: Built on ClamAV, trusted by millions globally for threat detection
- **High Performance**: Concurrent file scanning with automatic CPU-aware worker optimization
- **Easy Integration**: Clean REST API that requires minimal code changes to existing systems
- **Docker-Native**: Single docker-compose command for complete setup with ClamAV
- **Transparent & Open**: MIT-licensed open-source project with full source code visibility
- **Industry-Standard Detection**: Leverages regularly-updated ClamAV virus signatures
- **Production Monitoring**: Built-in health checks and detailed logging for production environments

## Key Features

✅ **Single File Scanning**
- Upload and scan individual files for malware in real-time
- Comprehensive threat detection with detailed results
- Support for unlimited file types and formats
- Configurable scan timeouts and resource limits

✅ **Batch File Processing**
- Efficiently scan multiple files concurrently in one request
- Automatic worker pool optimization based on CPU cores
- Handles hundreds of files without degradation
- Reduces Round-Trip Time (RTT) for bulk operations

✅ **RESTful API**
- Clean, intuitive HTTP/JSON interface
- Multipart form-data support for file uploads
- Structured JSON responses with full threat details
- Standard HTTP status codes for easy error handling

✅ **High Performance & Scalability**
- Non-blocking concurrent scanning using Go goroutines
- Automatic CPU-aware worker pool sizing
- Configurable request size limits (up to 500MB+)
- Minimal memory footprint with efficient resource utilization
- Horizontal scaling support via Docker

✅ **Flexible Configuration & Control**
- Environment-based configuration - no conf file parsing
- Customizable upload size limits per request
- ClamAV timeout and recursion settings
- CORS support for cross-origin requests from browsers
- Host and port customization for any deployment

✅ **Developer-Friendly Experience**
- Interactive Swagger/OpenAPI documentation with try-it-out
- Comprehensive error messages with actionable guidance
- Example curl commands and integration code samples
- Detailed logging for debugging and monitoring
- Health check endpoints for monitoring systems

✅ **Advanced Malware Intelligence**
- Detection of viruses, trojans, worms, rootkits, and spyware
- Signature-based detection with regularly updated databases
- Behavioral threat analysis and heuristic detection
- Detailed scanning reports with threat classifications
- File metadata preservation and analysis

✅ **Enterprise-Ready Features**
- Docker and Docker Compose support with pre-built images
- Health checks and readiness probes for orchestration
- Resource limits and timeout handling for stability
- CORS policy configuration for security
- Structured logging for centralized log aggregation
- Ready for Kubernetes and cloud-native deployments

## Why Choose Sereni Antivirus?

Sereni Antivirus stands out from alternatives through:

1. **Minimal Setup Time**: From installation to production in under 5 minutes with docker-compose
2. **Language Agnostic Integration**: Works seamlessly with any tech stack without SDK dependencies
3. **No Proprietary Dependencies**: Built entirely on open-source ClamAV engine
4. **Cost Effective**: Free and open-source, no licensing fees or usage limits
5. **Enterprise Support Ready**: Structured logging, monitoring integration, and comprehensive documentation

## Use Cases

### 1. **Web Application File Upload Security**
Protect your web application's file upload functionality by automatically scanning all user-uploaded files before storing them in cloud storage or database.

```
Client Upload → Web App → Sereni Antivirus → ClamAV → Clean/Threat Verdict → Storage Decision
```

**Why it matters**: Prevent malware distribution through user uploads, ensure compliance with security policies.

### 2. **Email Security Gateway**
Scan email attachments in real-time before delivery to end users. Integrate with email servers to check attachments against malware signatures.

```
Incoming Email → Email Gateway → Sereni Antivirus → Deliver/Quarantine Decision
```

**Why it matters**: Block malware-laden emails before they reach user inboxes, reduce security incidents.

### 3. **Document Management System**
Scan documents before indexing and archiving in your document management or DMS system. Ensure all stored documents are verified clean.

```
Document Upload → DMS → Sereni Antivirus → Archive/Reject → Search Index
```

**Why it matters**: Maintain a trusted document repository, prevent sharing of infected files, audit trail of scans.

### 4. **API & Microservice Network**
Use as a shared microservice across your microservices architecture. Multiple applications call Sereni for scanning without duplicating logic.

```
Service A, Service B, Service C → Sereni Antivirus (Shared) → Central Threat Intelligence
```

**Why it matters**: Avoid code duplication, centralized security policy, easier maintenance and updates.

### 5. **Cloud Storage Integration**
Scan files as they're uploaded to cloud storage (S3, Google Cloud Storage, Azure Blob). Quarantine suspicious files automatically.

```
Cloud Storage Upload → Webhook Trigger → Sereni Antivirus → Quarantine/Allow
```

**Why it matters**: Protect cloud data from infected uploads, automated compliance checking.

### 6. **Content Delivery Network (CDN)**
Scan and catalog user-generated content before CDN distribution. Ensure all distributed content is malware-free.

```
User Content → CDN Ingestion → Sereni Antivirus → CDN Distribution
```

**Why it matters**: Prevent distribution of infected files to users, maintain CDN safety reputation.

## Quick Start

### Prerequisites
- **Docker** (20.10+) - Container runtime
- **Docker Compose** (2.0+) - Multi-container orchestration
- OR **Go** (1.24.4+) - If running without containers
- **4GB RAM** minimum - For ClamAV virus database
- **2GB Disk Space** - For ClamAV signature database

### 30-Second Setup

```bash
# 1. Clone the repository
git clone https://github.com/aptlogica/sereni-antivirus-clamav.git
cd sereni-antivirus-clamav

# 2. Copy environment file
cp .env.example .env

# 3. Start with Docker Compose (recommended - auto-starts ClamAV)
docker-compose up -d

# 4. Verify the service is running
curl -X POST http://localhost:8084/scan \
  -F "file=@/etc/hostname" \
  -H "Content-Type: multipart/form-data"
```

The service is now available at `http://localhost:8084`

**What just happened:**
- Docker pulled the latest ClamAV and built the Sereni application
- ClamAV daemon started and initialized virus signatures (may take 30-60 seconds)
- Sereni Antivirus service bound to port 8084
- Your first test scan verified everything is working

**Next steps:** 
- Try the interactive API docs at http://localhost:8084/swagger/index.html
- See [Installation](#installation) for alternative setup methods
- Check [Usage](#usage) for practical examples
- Visit [Integration Guide](#integration-guide) for your specific language/framework

## Installation

### Option 1: Docker Compose (Recommended)

Best for: Quick setup, development, and production deployments. One command starts both ClamAV and Sereni.

```bash
# Step 1: Clone repository
git clone https://github.com/aptlogica/sereni-antivirus-clamav.git
cd sereni-antivirus-clamav

# Step 2: Copy environment template
cp .env.example .env

# Step 3: (Optional) Customize configuration
# Edit .env for port, upload size, timeouts, etc.
nano .env

# Step 4: Start services (pulls images, starts containers)
docker-compose up -d

# Step 5: Wait for ClamAV to initialize (check logs)
docker-compose logs -f clamav

# Step 6: Verify installation
docker-compose ps
```

**Result:** Both ClamAV and Sereni services running, health checks passing, ready for requests.

### Option 2: Manual Setup (Advanced Users)

Best for: Development, custom networking, or integration with existing ClamAV installations.

**Prerequisites:**
- Go 1.24.4 or higher installed
- ClamAV daemon running (separately or in another container)
- Network connectivity to ClamAV (default: localhost:3310)

```bash
# Step 1: Clone repository
git clone https://github.com/aptlogica/sereni-antivirus-clamav.git
cd sereni-antivirus-clamav

# Step 2: Copy environment template
cp .env.example .env

# Step 3: Configure environment
# Edit .env to point to your ClamAV instance:
# CLAMAV_ADDRESS=your-clamav-host:3310
nano .env

# Step 4: Install dependencies
go mod download

# Step 5: Run the application
go run cmd/server/main.go

# Step 6: Test the service
curl -X POST http://localhost:8084/scan \
  -F "file=@test.txt"
```

**Result:** Sereni running on port 8084, connected to ClamAV, ready for file scanning.

### Option 3: Docker Only (Without Docker Compose)

Best for: Kubernetes deployments, custom orchestration, or when ClamAV is hosted separately.

```bash
# Step 1: Clone and build
git clone https://github.com/aptlogica/sereni-antivirus-clamav.git
cd sereni-antivirus-clamav

# Step 2: Build Docker image
docker build -t sereni-antivirus:latest .

# Step 3: Create .env file for configuration
cat > .env << EOF
HOST=0.0.0.0
PORT=8084
CLAMAV_ADDRESS=clamav-host:3310
CLAMAV_TIMEOUT_SECONDS=30
MAX_UPLOAD_SIZE_MB=500
EOF

# Step 4: Run container
docker run -d \
  --name sereni-antivirus \
  --env-file .env \
  -p 8084:8084 \
  --network=host \
  sereni-antivirus:latest

# Step 5: Verify
curl http://localhost:8084/scan
```

**Result:** Sereni container running, ready to connect to external ClamAV instance.

## Configuration

### Environment Variables

Create `.env` file in your project root. All variables are optional with sensible defaults:

```dotenv
# === Server Configuration ===
HOST=localhost                      # Server bind address
PORT=8084                           # Server port (change for multiple instances)
BASE_URL=http://localhost:8084      # Base URL for API (used in docs)

# === Antivirus Configuration ===
ANTIVIRUS_DRIVER=clamav             # Antivirus engine (only clamav supported)

# === ClamAV Configuration ===
CLAMAV_ADDRESS=127.0.0.1:3310       # ClamAV daemon address and port
CLAMAV_TIMEOUT_SECONDS=30           # Scan timeout in seconds

# === Upload Configuration ===
MAX_UPLOAD_SIZE_MB=32               # Max single file size in MB
                                     # For 500MB+, ensure ClamAV MaxFileSize is updated

# === CORS Configuration (Optional) ===
ALLOWED_ORIGINS=*                   # Comma-separated list of origins (*=all)
```

### Default Values

| Variable | Default | Description |
|----------|---------|-------------|
| `HOST` | `localhost` | Bind address (use `0.0.0.0` for Docker) |
| `PORT` | `8084` | Server port |
| `CLAMAV_ADDRESS` | `127.0.0.1:3310` | ClamAV TCP/IP address |
| `CLAMAV_TIMEOUT_SECONDS` | `30` | Max scan time per file |
| `MAX_UPLOAD_SIZE_MB` | `32` | Max upload size |
| `ANTIVIRUS_DRIVER` | `clamav` | Scanning engine |

### Configuration Examples

**For Development:**
```dotenv
HOST=localhost
PORT=8084
CLAMAV_ADDRESS=localhost:3310
CLAMAV_TIMEOUT_SECONDS=30
MAX_UPLOAD_SIZE_MB=512
ALLOWED_ORIGINS=localhost:3000,localhost:3001
```

**For Production:**
```dotenv
HOST=0.0.0.0
PORT=8084
CLAMAV_ADDRESS=clamav-service:3310
CLAMAV_TIMEOUT_SECONDS=20
MAX_UPLOAD_SIZE_MB=500
ALLOWED_ORIGINS=app.example.com,api.example.com
```

**For Kubernetes:**
```dotenv
HOST=0.0.0.0
PORT=8084
CLAMAV_ADDRESS=${CLAMAV_SERVICE_HOST}:${CLAMAV_SERVICE_PORT}
CLAMAV_TIMEOUT_SECONDS=30
MAX_UPLOAD_SIZE_MB=500
```

## Usage

### Basic Usage - Scan Single File

```bash
# Scan a file from command line
curl -X POST http://localhost:8084/scan \
  -F "file=@document.pdf"
```

### Example 1: Scan Single File with Detailed Response

```bash
# Upload and scan a single file
curl -X POST http://localhost:8084/scan \
  -F "file=@virus-sample.exe" \
  -H "Content-Type: multipart/form-data"
```

**Output (Clean File):**
```json
{
  "filename": "virus-sample.exe",
  "clean": true,
  "threat": "",
  "description": "No threat detected",
  "scan_time_ms": 145
}
```

**Output (Infected File):**
```json
{
  "filename": "virus-sample.exe",
  "clean": false,
  "threat": "Win.Malware.123-456",
  "description": "Detected Win.Malware.123-456",
  "scan_time_ms": 342
}
```

### Example 2: Batch Scan Multiple Files

```bash
# Scan multiple files in one request
curl -X POST http://localhost:8084/scan-files \
  -F "files=@document1.pdf" \
  -F "files=@document2.docx" \
  -F "files=@image.jpg" \
  -H "Content-Type: multipart/form-data"
```

**Output:**
```json
[
  {
    "filename": "document1.pdf",
    "clean": true,
    "threat": "",
    "description": "No threat detected",
    "scan_time_ms": 234
  },
  {
    "filename": "document2.docx",
    "clean": false,
    "threat": "W97M.Malware.A",
    "description": "Detected W97M.Malware.A",
    "scan_time_ms": 567
  },
  {
    "filename": "image.jpg",
    "clean": true,
    "threat": "",
    "description": "No threat detected",
    "scan_time_ms": 89
  }
]
```

### Example 3: Large File Upload

```bash
# Upload and scan a large file (respects MAX_UPLOAD_SIZE_MB setting)
curl -X POST http://localhost:8084/scan \
  --form "file=@large-archive-500mb.zip" \
  --max-time 120  # 2 minute timeout for large file
```

### Example 4: Using with File Processing Pipeline

```bash
# Scan file before processing/storage
#!/bin/bash

FILE=$1
SCAN_ENDPOINT="http://localhost:8084/scan"

# Scan the file
RESULT=$(curl -s -X POST $SCAN_ENDPOINT \
  -F "file=@$FILE" \
  -w "\n%{http_code}")

HTTP_CODE=$(echo "$RESULT" | tail -n1)
RESPONSE=$(echo "$RESULT" | head -n-1)

# Check if scan was successful
if [ "$HTTP_CODE" -ne 200 ]; then
  echo "Scan failed with status $HTTP_CODE"
  exit 1
fi

# Parse response and decide
CLEAN=$(echo "$RESPONSE" | grep -o '"clean":[^,}]*' | grep -o 'true\|false')

if [ "$CLEAN" = "true" ]; then
  echo "File is clean, processing..."
  # Continue with file processing
else
  echo "File is infected! Quarantining..."
  # Move file to quarantine folder
  mv "$FILE" "/quarantine/$FILE"
fi
```

## API Documentation

### Interactive API Docs

Once the service is running, visit the interactive Swagger UI:

**http://localhost:8084/swagger/index.html**

This provides:
- Full endpoint documentation
- Request/response schemas
- Try-it-out functionality to test endpoints
- Example responses
- Authentication details (if any)

### Endpoints

#### POST /scan
```http
POST /scan HTTP/1.1
Host: localhost:8084
Content-Type: multipart/form-data
```

**Description:** Upload and scan a single file for malware

**Parameters:**
- `file` (required, multipart/form-data): File to scan

**Request (curl):**
```bash
curl -X POST http://localhost:8084/scan \
  -F "file=@sample.pdf" \
  -H "Content-Type: multipart/form-data"
```

**Response (Success - 200):**
```json
{
  "filename": "sample.pdf",
  "clean": true,
  "threat": "",
  "description": "No threat detected",
  "scan_time_ms": 245
}
```

**Response (Infected - 200):**
```json
{
  "filename": "malware.exe",
  "clean": false,
  "threat": "Win.Malware.Generic",
  "description": "Detected Win.Malware.Generic",
  "scan_time_ms": 567
}
```

**Response (Error - 400):**
```json
{
  "error": "File size exceeds maximum allowed size"
}
```

#### POST /scan-files
```http
POST /scan-files HTTP/1.1
Host: localhost:8084
Content-Type: multipart/form-data
```

**Description:** Upload and scan multiple files concurrently

**Parameters:**
- `files` (required, multipart/form-data[], multiple): Files to scan

**Request (curl):**
```bash
curl -X POST http://localhost:8084/scan-files \
  -F "files=@file1.pdf" \
  -F "files=@file2.docx" \
  -F "files=@file3.zip"
```

**Response (Success - 200):**
```json
[
  {
    "filename": "file1.pdf",
    "clean": true,
    "threat": "",
    "description": "No threat detected",
    "scan_time_ms": 234
  },
  {
    "filename": "file2.docx",
    "clean": true,
    "threat": "",
    "description": "No threat detected",
    "scan_time_ms": 156
  },
  {
    "filename": "file3.zip",
    "clean": false,
    "threat": "PK.Malware.A",
    "description": "Detected PK.Malware.A",
    "scan_time_ms": 892
  }
]
```

### Error Codes

| Code | Name | Description | Solution |
|------|------|-------------|----------|
| 200 | OK | Request successful, file(s) scanned | Check `clean` field in response |
| 400 | Bad Request | Invalid request (missing file, invalid params) | Verify file is attached and format is correct |
| 413 | Payload Too Large | File exceeds MAX_UPLOAD_SIZE_MB | Reduce file size or increase MAX_UPLOAD_SIZE_MB |
| 500 | Internal Server Error | ClamAV unreachable or scan error | Check ClamAV service, review logs |
| 504 | Gateway Timeout | Scan took longer than CLAMAV_TIMEOUT_SECONDS | Increase timeout or check ClamAV performance |

## Integration Guide

### JavaScript / Node.js + Express

```javascript
const express = require('express');
const axios = require('axios');
const FormData = require('form-data');
const fs = require('fs');

const app = express();
const SERENI_URL = 'http://localhost:8084';

// Middleware to handle file uploads
const multer = require('multer');
const upload = multer({ dest: 'uploads/' });

// Scan single file endpoint
app.post('/upload', upload.single('file'), async (req, res) => {
  try {
    // Create form data for Sereni
    const formData = new FormData();
    formData.append('file', fs.createReadStream(req.file.path));

    // Call Sereni antivirus service
    const response = await axios.post(`${SERENI_URL}/scan`, formData, {
      headers: formData.getHeaders(),
      timeout: 60000 // 60 second timeout
    });

    const scanResult = response.data;

    // Clean up temp file
    fs.unlinkSync(req.file.path);

    if (scanResult.clean) {
      res.json({ 
        success: true, 
        message: 'File is safe', 
        file: req.file.originalname 
      });
    } else {
      res.status(400).json({ 
        success: false, 
        error: `Malware detected: ${scanResult.threat}` 
      });
    }
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

app.listen(3000, () => console.log('Server running on port 3000'));
```

### Python / Flask

```python
from flask import Flask, request, jsonify
import requests
import os

app = Flask(__name__)
SERENI_URL = 'http://localhost:8084'

@app.route('/upload', methods=['POST'])
def upload_and_scan():
    try:
        # Check if file is in request
        if 'file' not in request.files:
            return jsonify({'error': 'No file provided'}), 400

        file = request.files['file']
        
        # Prepare multipart request to Sereni
        files = {'file': (file.filename, file.stream, file.content_type)}
        
        # Call Sereni antivirus service
        response = requests.post(
            f'{SERENI_URL}/scan',
            files=files,
            timeout=60
        )

        scan_result = response.json()

        if response.status_code == 200:
            if scan_result.get('clean'):
                return jsonify({
                    'success': True,
                    'message': 'File is safe',
                    'filename': file.filename
                })
            else:
                return jsonify({
                    'success': False,
                    'error': f"Malware detected: {scan_result.get('threat')}"
                }), 400
        else:
            return jsonify({'error': 'Scan failed'}), response.status_code

    except Exception as e:
        return jsonify({'error': str(e)}), 500

if __name__ == '__main__':
    app.run(debug=True, port=5000)
```

### Java / Spring Boot

```java
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;
import org.springframework.web.client.RestTemplate;
import org.springframework.http.*;
import java.util.*;

@RestController
@RequestMapping("/api")
public class AntivirusController {
    
    private final RestTemplate restTemplate = new RestTemplate();
    private static final String SERENI_URL = "http://localhost:8084";
    
    @PostMapping("/upload")
    public ResponseEntity<?> uploadAndScan(@RequestParam("file") MultipartFile file) {
        try {
            // Create multipart request to Sereni
            HttpHeaders headers = new HttpHeaders();
            headers.setContentType(MediaType.MULTIPART_FORM_DATA);
            
            // Prepare the request body with file
            LinkedMultiValueMap<String, Object> body = new LinkedMultiValueMap<>();
            body.add("file", file.getResource());
            
            HttpEntity<LinkedMultiValueMap<String, Object>> requestEntity = 
                new HttpEntity<>(body, headers);
            
            // Call Sereni
            ResponseEntity<Map> response = restTemplate.exchange(
                SERENI_URL + "/scan",
                HttpMethod.POST,
                requestEntity,
                Map.class
            );
            
            Map<String, Object> scanResult = response.getBody();
            Boolean isClean = (Boolean) scanResult.get("clean");
            
            if (isClean) {
                return ResponseEntity.ok(Map.of(
                    "success", true,
                    "message", "File is safe",
                    "filename", file.getOriginalFilename()
                ));
            } else {
                return ResponseEntity.badRequest().body(Map.of(
                    "success", false,
                    "error", "Malware detected: " + scanResult.get("threat")
                ));
            }
        } catch (Exception e) {
            return ResponseEntity.status(500).body(Map.of("error", e.getMessage()));
        }
    }
}
```

### PHP / Laravel

```php
<?php
namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class AntivirusController extends Controller
{
    private const SERENI_URL = 'http://localhost:8084';

    public function uploadAndScan(Request $request)
    {
        try {
            // Validate file upload
            $request->validate(['file' => 'required|file']);
            
            $file = $request->file('file');
            
            // Call Sereni antivirus service
            $response = Http::attach(
                'file',
                fopen($file->getRealPath(), 'r'),
                $file->getClientOriginalName()
            )->post(self::SERENI_URL . '/scan');

            if ($response->successful()) {
                $scanResult = $response->json();
                
                if ($scanResult['clean']) {
                    return response()->json([
                        'success' => true,
                        'message' => 'File is safe',
                        'filename' => $file->getClientOriginalName()
                    ]);
                } else {
                    return response()->json([
                        'success' => false,
                        'error' => 'Malware detected: ' . $scanResult['threat']
                    ], 400);
                }
            } else {
                return response()->json(['error' => 'Scan failed'], 500);
            }
        } catch (\Exception $e) {
            return response()->json(['error' => $e->getMessage()], 500);
        }
    }
}
```

### C# / ASP.NET Core

```csharp
using System;
using System.IO;
using System.Net.Http;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;

[ApiController]
[Route("api/[controller]")]
public class AntivirusController : ControllerBase
{
    private readonly HttpClient _httpClient;
    private const string SERENI_URL = "http://localhost:8084";

    public AntivirusController(HttpClient httpClient)
    {
        _httpClient = httpClient;
    }

    [HttpPost("upload")]
    public async Task<IActionResult> UploadAndScan([FromForm] IFormFile file)
    {
        try
        {
            if (file == null || file.Length == 0)
                return BadRequest("No file provided");

            // Create multipart content
            using (var content = new MultipartFormDataContent())
            {
                using (var fileStream = file.OpenReadStream())
                {
                    var fileContent = new StreamContent(fileStream);
                    fileContent.Headers.ContentType = 
                        System.Net.Http.Headers.MediaTypeHeaderValue.Parse(file.ContentType);
                    
                    content.Add(fileContent, "file", file.FileName);

                    // Call Sereni antivirus service
                    var response = await _httpClient.PostAsync(
                        SERENI_URL + "/scan",
                        content
                    );

                    if (response.IsSuccessStatusCode)
                    {
                        var jsonResponse = await response.Content.ReadAsStringAsync();
                        dynamic scanResult = JsonConvert.DeserializeObject(jsonResponse);
                        
                        if (scanResult.clean)
                        {
                            return Ok(new 
                            { 
                                success = true, 
                                message = "File is safe",
                                filename = file.FileName
                            });
                        }
                        else
                        {
                            return BadRequest(new 
                            { 
                                success = false, 
                                error = $"Malware detected: {scanResult.threat}" 
                            });
                        }
                    }
                    else
                    {
                        return StatusCode(500, "Scan failed");
                    }
                }
            }
        }
        catch (Exception ex)
        {
            return StatusCode(500, new { error = ex.Message });
        }
    }
}
```

## Architecture

### System Architecture

```
┌─────────────────────────────────────────────────────────────┐
│              Client Applications                            │
│   (Web, Mobile, Desktop, CLI, APIs, Microservices)         │
└────────────────────┬────────────────────────────────────────┘
                     │ HTTP/REST (any language)
                     │
┌────────────────────▼────────────────────────────────────────┐
│           Sereni Antivirus API Service                      │
│  ┌────────────────────────────────────────────────────┐    │
│  │  HTTP Handlers: /scan, /scan-files, /swagger       │    │
│  │  CORS Middleware │ Request Validation              │    │
│  └────────────────┬─────────────────────────────────┘    │
│                   │                                         │
│  ┌────────────────▼──────────────────────────────────┐    │
│  │  Antivirus Service Layer                          │    │
│  │  - File processing logic                          │    │
│  │  - Worker pool management                         │    │
│  │  - Error handling                                 │    │
│  └────────────────┬──────────────────────────────────┘    │
│                   │                                         │
│  ┌────────────────▼──────────────────────────────────┐    │
│  │  ClamAV Provider (Interface)                      │    │
│  │  - Abstraction layer for antivirus engines        │    │
│  │  - Extensible for future engines                 │    │
│  └────────────────┬──────────────────────────────────┘    │
└────────────────────┼────────────────────────────────────────┘
                     │ TCP/IP Port 3310
    ┌────────────────▼──────────────────────┐
    │                                       │
    ▼                                       ▼
┌──────────────────────┐         ┌──────────────────────┐
│   ClamAV Daemon      │         │  Docker Container    │
│   (Virus Scanning)   │         │  (Long-term storage) │
│                      │         │                      │
│ - Signature Database │         │ - Virus definitions  │
│ - Scan Engine        │         │ - ClamAV config      │
│ - Heuristics         │         │ - Health checks      │
└──────────────────────┘         └──────────────────────┘
```

### Component Responsibilities

| Component | Responsibility |
|-----------|-----------------|
| **HTTP Handlers** | Receive file uploads, validate requests, format responses |
| **Route Layer** | Map HTTP endpoints to handlers, setup CORS, add middleware |
| **Antivirus Service** | Orchestrate file scanning, manage worker pools, aggregate results |
| **ClamAV Provider** | Interface with ClamAV daemon via TCP/IP protocol |
| **ClamAV Daemon** | Perform actual malware scanning using signature database |

### Design Patterns

- **Provider Pattern**: ClamAV provider abstracts the antivirus implementation, allowing future support for multiple engines
- **Service Layer**: AntivirusService decouples HTTP handlers from scanning logic
- **Worker Pool**: Goroutine-based workers for concurrent file processing without goroutine explosion
- **Factory Pattern**: Factory method creates appropriate antivirus provider based on configuration

## Development

### Project Structure

```
sereni-antivirus-clamav/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point, initialization
├── internal/
│   ├── config/
│   │   └── config.go            # Environment configuration loading
│   ├── handlers/
│   │   └── scan_handler.go      # HTTP request handlers for /scan endpoints
│   ├── providers/
│   │   ├── antivirus/
│   │   │   ├── clamav/
│   │   │   │   └── clamav.go    # ClamAV integration logic
│   │   │   └── factory.go       # Provider factory pattern
│   ├── routes/
│   │   └── routes.go            # Router setup and middleware
│   └── services/
│       └── antivirus_service.go # Business logic for scanning
├── tests/
│   └── ...                      # Unit and integration tests
├── docs/
│   └── ...                      # Generated Swagger documentation
├── Dockerfile                   # Container image definition
├── docker-compose.yml           # Multi-container orchestration
├── Makefile                     # Build automation
├── go.mod / go.sum             # Dependency management
├── .env.example                # Environment template
└── README.md                    # This file
```

### Development Setup

```bash
# 1. Clone repository
git clone https://github.com/aptlogica/sereni-antivirus-clamav.git
cd sereni-antivirus-clamav

# 2. Install Go dependencies
go mod download

# 3. Copy environment file
cp .env.example .env

# 4. Edit configuration if needed (optional)
nano .env

# 5. Start ClamAV separately (if not using docker-compose)
# Use existing ClamAV installation or run in Docker

# 6. Run application
go run cmd/server/main.go

# Or build binary first
go build -o antivirus cmd/server/main.go
./antivirus
```

### Running Tests

```bash
# Run all tests with coverage
make test
# or
go test -cover -coverpkg=./internal/... -coverprofile=coverage.out ./tests

# View coverage in browser
make coverage
# or
go tool cover -html=coverage.out

# View coverage as text
make coverage-func
# or  
go tool cover -func=coverage.out

# Run specific test
go test -v ./tests -run TestScanFile

# Run tests with verbose output
go test -v ./tests
```

### Building for Production

```bash
# Build binary for current OS
make build
# or
go build -o server.exe cmd/server/main.go

# Build Docker image
docker build -t sereni-antivirus:latest .

# Tag for registry
docker tag sereni-antivirus:latest myregistry.azurecr.io/sereni-antivirus:1.0.0

# Push to registry
docker push myregistry.azurecr.io/sereni-antivirus:1.0.0

# Generate Swagger documentation
make swag
# or
swag init -g cmd/server/main.go -o docs
```

### Code Standards

- Follow Go idioms and package conventions
- Use meaningful variable names and add comments for complex logic
- Write unit tests for new features (aim for >80% coverage)
- Run `gofmt` or use `go fmt ./...` before committing
- Add documentation comments to exported functions

## Troubleshooting

### Common Issues

#### 1. "Connection refused" to ClamAV

**Error Message:**
```
failed to initialize antivirus provider: connection refused
```

**Root Cause:** ClamAV daemon is not running or not accessible at CLAMAV_ADDRESS

**Solution:**
```bash
# Step 1: Verify ClamAV is running
docker-compose ps clamav
# Should show "healthy" status

# Step 2: Check logs for ClamAV
docker-compose logs clamav
# Look for "clamd accepting connections"

# Step 3: If using external ClamAV, verify network connectivity
nc -zv clamav-host 3310
# Should show "succeeded" if reachable

# Step 4: Update .env with correct address
echo "CLAMAV_ADDRESS=clamav-host:3310" >> .env

# Step 5: Restart Sereni
docker-compose restart antivirus-service
```

#### 2. "413 Payload Too Large" when uploading

**Error Message:**
```
413 Payload Too Large: File size exceeds maximum allowed size
```

**Solution:**
```bash
# Step 1: Check current limit
grep MAX_UPLOAD_SIZE_MB .env

# Step 2: Increase limit in .env
echo "MAX_UPLOAD_SIZE_MB=500" >> .env

# Step 3: Also update ClamAV limits in docker-compose.yml
# Change MaxFileSize, MaxScanSize, StreamMaxLength

# Step 4: Restart services
docker-compose restart

# Step 5: Test with larger file
curl -X POST http://localhost:8084/scan -F "file=@large-file.bin"
```

#### 3. Scan takes too long (timeout)

**Error Message:**
```
504 Gateway Timeout: Scan took longer than configured timeout
```

**Solution:**
```bash
# Step 1: Check current timeout
grep CLAMAV_TIMEOUT_SECONDS .env

# Step 2: Increase timeout
echo "CLAMAV_TIMEOUT_SECONDS=60" >> .env

# Step 3: Check ClamAV performance
docker-compose logs clamav | tail -20

# Step 4: Check system resources
docker stats

# Step 5: Restart and test
docker-compose restart
# Try scanning again with increased timeout curl --max-time 120
```

#### 4. ClamAV signature database is outdated

**Error Message:**
```
Warning: Signatures are X days old
```

**Solution:**
```bash
# Step 1: Check ClamAV logs
docker-compose logs clamav | grep -i signature

# Step 2: Manually update signatures
docker exec clamav freshclam

# Step 3: Verify update
docker exec clamav clamscan --version

# Step 4: If auto-update failing, check environment
docker-compose exec clamav env | grep FRESHCLAM

# Step 5: Restart ClamAV daemon
docker-compose restart clamav
```

#### 5. High memory usage by ClamAV

**Error Message:**
```
Container exceeds memory limit, OOMKilled
```

**Solution:**
```bash
# Step 1: Check current usage
docker stats --no-stream clamav

# Step 2: Reduce ClamAV memory usage in docker-compose.yml
# Reduce MaxRecursion, MaxFiles, or disable certain features

# Step 3: Increase Docker memory allocation
# Edit docker-compose.yml:
services:
  clamav:
    deploy:
      resources:
        limits:
          memory: 4G

# Step 4: Restart
docker-compose up -d
```

### Health Checks

```bash
# Basic health check - is service responding
curl http://localhost:8084/scan

# Detailed status - check ClamAV connectivity
curl -X POST http://localhost:8084/scan \
  -F "file=@/etc/hostname" \
  -w "\nStatus: %{http_code}\n"

# Check Docker service status
docker-compose ps

# Check service logs for errors
docker-compose logs antivirus-service | tail -20

# Verify ClamAV connection
docker exec antivirus-service \
  /app/antivirus-service -test-connection
```

### Debugging

```bash
# View full application logs
docker-compose logs -f antivirus-service

# View with timestamps and tail last 50 lines
docker-compose logs --timestamps -f --tail=50 antivirus-service

# View ClamAV debug logs  
docker-compose logs -f clamav

# Check network connectivity between services
docker exec antivirus-service netstat -an | grep 3310

# Inspect running processes
docker exec antivirus-service ps aux

# Test ClamAV directly from Sereni container
docker exec antivirus-service clamdscan --version
```

## Contributing

We welcome contributions! Here's how to contribute:

### Getting Started

1. Fork the repository on GitHub
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/sereni-antivirus-clamav.git`
3. Create feature branch: `git checkout -b feature/amazing-feature`
4. Make your changes following code standards (see below)
5. Write tests for new functionality
6. Ensure all tests pass: `make test`
7. Commit with clear, descriptive messages (see format below)
8. Push to your branch: `git push origin feature/amazing-feature`
9. Create Pull Request with detailed description

### Code Standards

- **Go Idioms**: Follow effective Go practices and standard library conventions
- **Formatting**: Run `go fmt ./...` before committing - use `gofmt` or IDE integration
- **Naming**: Use clear, concise names - `handler` not `h`, `scanResult` not `sr`
- **Comments**: Add comments for public types and functions, explain "why" not "what"
- **Error Handling**: Don't ignore errors, handle them explicitly with meaningful messages
- **Tests**: Write unit tests for new features, aim for >80% code coverage
- **Dependencies**: Keep dependencies minimal and current

### Testing Requirements

- Write unit tests for all new features
- Ensure all tests pass: `make test`
- Maintain at least 80% code coverage
- Test edge cases, error conditions, and boundary values
- Add integration tests for cross-component interactions

### Commit Message Format

```
type(scope): description

Longer explanation if needed. Wrap at 72 characters.

Fixes #123
```

**Types:** `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`

**Examples:**
- `feat(api): add support for custom timeout configuration`
- `fix(scan): handle files with empty extensions`
- `docs(readme): add Kubernetes deployment examples`
- `test(service): add tests for concurrent scanning`

### Pull Request Checklist

- [ ] Tests pass: `make test`
- [ ] Code follows standards: `go fmt && gofmt`
- [ ] Documentation updated for new features
- [ ] No unnecessary changes or dependencies
- [ ] Commit messages are clear and descriptive
- [ ] PR addresses an issue or feature request
- [ ] No hardcoded values or secrets

## FAQ

**Q: Can I use Sereni Antivirus in production?**
A: Absolutely! It's designed for production use with health checks, error handling, and Docker support. See [Architecture](#architecture) for deployment patterns and [Configuration](#configuration) for production settings.

**Q: What are the system requirements?**
A: Minimum 4GB RAM for ClamAV database, Go 1.24.4+ if building from source, Docker 20.10+ for containerized deployment. See [Quick Start](#quick-start) for details.

**Q: Does it work with any programming language?**
A: Yes! Since Sereni Antivirus is a RESTful API microservice, it works with any language or framework (Node.js, Python, Java, .NET, PHP, Ruby, Go, Rust, etc.). See [Integration Guide](#integration-guide) for examples.

**Q: How do I scale this for high traffic?**
A: Sereni is stateless and horizontally scalable - run multiple instances behind a load balancer. ClamAV can also run in a separate container scaled independently. See [Architecture](#architecture) for scaling patterns.

**Q: What are the performance limitations?**
A: Performance depends on file size, complexity, and ClamAV updates. Typical speeds: 10-100 MB/sec. For optimization, configure `MAX_UPLOAD_SIZE_MB`, `CLAMAV_TIMEOUT_SECONDS`, and server resources. See [Configuration](#configuration).

**Q: Is this truly open source?**
A: Yes! Licensed under the MIT License - full source code, no proprietary components. See [License](#license) for details. Contributions welcome!

**Q: How do I report security issues?**
A: Please email security@aptlogica.com with details. Do not open public issues for security vulnerabilities. See [Contributing](#contributing) for responsible disclosure.

**Q: Can I use a different antivirus engine?**
A: Currently supports ClamAV only. The provider pattern allows easy addition of new engines. Open an issue or PR for engine requests.

**Q: What about scanning files larger than 500MB?**
A: ClamAV and Sereni support files up to several GB. Update `MaxFileSize` in docker-compose.yml and `MAX_UPLOAD_SIZE_MB` in .env. Client timeout may need adjustment.

**Q: Does it scan inside archives (ZIP, RAR, 7Z)?**
A: Yes! ClamAV extracts and scans archive contents based on `MaxRecursion` and `MaxFiles` settings in docker-compose.yml.

**Q: How often are virus signatures updated?**
A: By default, ClamAV checks for updates every 2 hours (12 times daily). Adjust `FRESHCLAM_CHECKS` in docker-compose.yml for different schedules.

## License

This project is licensed under the MIT License - see [LICENSE](LICENSE) file for details.

**Summary:**
✅ Free for private use  
✅ Free for commercial use  
✅ Modify and distribute  
✅ Include in proprietary products  
❌ Hold liable for damages

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

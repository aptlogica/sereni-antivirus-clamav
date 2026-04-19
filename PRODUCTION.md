# Production Deployment Guide

This document provides essential guidance for deploying and operating the Sereni Antivirus ClamAV service in production environments.

## Table of Contents

- [Environmental Configuration](#environmental-configuration)
- [Resource Constraints & Sizing](#resource-constraints--sizing)
- [Health Check Mechanisms](#health-check-mechanisms)
- [Persistent Volumes](#persistent-volumes)
- [Deployment Best Practices](#deployment-best-practices)
- [Monitoring & Troubleshooting](#monitoring--troubleshooting)

## Environmental Configuration

### Application Configuration

Configure the antivirus service using environment variables:

| Variable | Default | Description | Example |
|----------|---------|-------------|---------|
| `HOST` | `localhost` | Service bind address | `0.0.0.0` |
| `PORT` | `8084` | Service port | `8084` |
| `ALLOWED_ORIGINS` | `*` | CORS allowed origins (comma-separated) | `https://example.com,https://api.example.com` |
| `ANTIVIRUS_DRIVER` | `clamav` | Antivirus engine driver | `clamav` |
| `CLAMAV_ADDRESS` | `127.0.0.1:3310` | ClamAV daemon address | `clamav:3310` |
| `CLAMAV_TIMEOUT_SECONDS` | `30` | Scan timeout in seconds | `60` |
| `MAX_UPLOAD_SIZE_MB` | `32` | Maximum file upload size in MB | `256` |

### ClamAV Configuration

ClamAV daemon configuration is managed through environment variables in the `clamav` service:

| Variable | Default | Description |
|----------|---------|-------------|
| `CLAMD_STARTUP_TIMEOUT` | `180` | Clamd startup timeout in seconds |
| `FRESHCLAM_CHECKS` | `12` | Freshclam update checks per day (1 = every 24h) |
| `MaxScanSize` | `500M` | Maximum size of file to scan |
| `MaxFileSize` | `500M` | Maximum file size allowed |
| `StreamMaxLength` | `500M` | Maximum stream length to scan |
| `MaxRecursion` | `20` | Maximum recursion depth in archives |
| `MaxFiles` | `10000` | Maximum number of files to scan |

### Example .env Configuration

```bash
# Application Service
HOST=0.0.0.0
PORT=8084
ALLOWED_ORIGINS=https://example.com,https://api.example.com

# Antivirus Configuration
ANTIVIRUS_DRIVER=clamav
CLAMAV_ADDRESS=clamav:3310
CLAMAV_TIMEOUT_SECONDS=60
MAX_UPLOAD_SIZE_MB=256

# ClamAV Configuration (passed via docker-compose)
# CLAMD_STARTUP_TIMEOUT=180
# FRESHCLAM_CHECKS=12
```

## Resource Constraints & Sizing

### Minimum Requirements

For small-scale deployments (< 100 scans/hour):

**ClamAV Container:**
- CPU: 0.5 cores
- Memory: 512 MB
- Disk: 500 MB (virus database + 1 GB buffer)

**Antivirus Service Container:**
- CPU: 0.25 cores
- Memory: 256 MB

### Recommended Requirements

For production deployments (100-500 scans/hour):

**ClamAV Container:**
- CPU: 1-2 cores
- Memory: 1-2 GB
- Disk: 2-3 GB (virus database + working space)

**Antivirus Service Container:**
- CPU: 1 core
- Memory: 512 MB

### Large-Scale Requirements

For high-throughput deployments (500+ scans/hour):

**ClamAV Container:**
- CPU: 2-4 cores
- Memory: 2-4 GB
- Disk: 4-5 GB (database + cache + working space)

**Antivirus Service Container:**
- CPU: 2-4 cores
- Memory: 1-2 GB

### Kubernetes Resource Requests/Limits Example

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: clamav-pod
spec:
  containers:
  - name: clamav
    image: clamav/clamav:1.4.5
    resources:
      requests:
        cpu: 500m
        memory: 512Mi
      limits:
        cpu: 2000m
        memory: 2Gi
    volumeMounts:
    - name: virus-db
      mountPath: /var/lib/clamav
      
  - name: antivirus-service
    image: sereni-antivirus:latest
    resources:
      requests:
        cpu: 250m
        memory: 256Mi
      limits:
        cpu: 1000m
        memory: 512Mi
    env:
    - name: CLAMAV_ADDRESS
      value: "localhost:3310"
    - name: MAX_UPLOAD_SIZE_MB
      value: "256"
      
  volumes:
  - name: virus-db
    persistentVolumeClaim:
      claimName: clamav-db-pvc
```

## Health Check Mechanisms

### ClamAV Service Health

The docker-compose configuration includes a health check for ClamAV:

```yaml
healthcheck:
  test: ["CMD", "clamdscan", "--version"]
  interval: 60s
  timeout: 10s
  retries: 5
  start_period: 30s
```

This check:
- Runs every 60 seconds
- Verifies ClamAV daemon is responsive
- Waits 30 seconds before starting checks (initialization period)
- Marks unhealthy after 5 consecutive failures

### Antivirus Service Health

Implement application-level health checks:

**HTTP Health Endpoint:** `GET /health`

Expected response when healthy:
```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "status": "healthy",
  "clamav_connected": true,
  "uptime_seconds": 3600
}
```

**Kubernetes Probe Configuration:**

```yaml
livenessProbe:
  httpGet:
    path: /health
    port: 8084
  initialDelaySeconds: 30
  periodSeconds: 10
  timeoutSeconds: 5
  failureThreshold: 3

readinessProbe:
  httpGet:
    path: /health
    port: 8084
  initialDelaySeconds: 10
  periodSeconds: 5
  timeoutSeconds: 3
  failureThreshold: 2
```

### Monitoring Health

Common health check patterns:

1. **Load Balancer Health Checks:**
   ```bash
   curl -f http://localhost:8084/health || exit 1
   ```

2. **Docker Compose Health Monitoring:**
   ```bash
   docker-compose ps  # View health status
   docker-compose logs clamav  # View ClamAV logs
   ```

3. **Kubernetes Health Status:**
   ```bash
   kubectl get pods -o wide  # View pod status
   kubectl describe pod <pod-name>  # Detailed health info
   ```

### Handling Unhealthy States

**ClamAV Recovery:**
- Container auto-restart (via `restart: unless-stopped`)
- Automatic virus database refresh (freshclam service)
- Graceful shutdown with persistent database

**Service Recovery:**
- Connection retry logic with exponential backoff
- Request timeout handling (CLAMAV_TIMEOUT_SECONDS)
- Automatic reconnection on ClamAV restart

## Persistent Volumes

### Why Persistent Volumes Are Essential

The ClamAV virus database is large (~500 MB) and requires frequent updates via the `freshclam` daemon. Without persistent volumes:

- Database is re-downloaded on every container restart (5-10 minutes per restart)
- Increased bandwidth consumption
- Service downtime during database initialization
- Poor scalability in orchestrated environments

### Volume Configuration

**Docker Compose Setup:**

```yaml
services:
  clamav:
    image: clamav/clamav:1.4.5
    volumes:
      - clamav-db:/var/lib/clamav  # Persistent virus database
      - ./data:/data                # Optional: scan-related files

volumes:
  clamav-db:
    driver: local
```

**Kubernetes Persistent Volume:**

```yaml
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: clamav-db-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 3Gi
  storageClassName: standard

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: clamav
spec:
  template:
    spec:
      containers:
      - name: clamav
        image: clamav/clamav:1.4.5
        volumeMounts:
        - name: virus-db
          mountPath: /var/lib/clamav
      volumes:
      - name: virus-db
        persistentVolumeClaim:
          claimName: clamav-db-pvc
```

### Volume Maintenance

**Checking Database Status:**

```bash
# Inside ClamAV container
docker exec clamav freshclam -V

# View database files
docker exec clamav ls -lh /var/lib/clamav/
```

**Manual Database Update:**

```bash
# Force immediate database update
docker exec clamav freshclam -v

# Monitor update progress
docker logs -f clamav
```

**Volume Cleanup (when needed):**

```bash
# Docker Compose
docker volume rm clamav-db
docker-compose up -d  # Recreate and initialize

# Kubernetes
kubectl delete pvc clamav-db-pvc
kubectl apply -f deployment.yaml  # Recreate
```

### Backup & Recovery

**Backup ClamAV Database:**

```bash
# Docker
docker run --rm -v clamav-db:/db -v $(pwd):/backup \
  alpine tar czf /backup/clamav-db.tar.gz -C /db .

# Kubernetes
kubectl cp clamav-pod:/var/lib/clamav ./clamav-db-backup
```

**Restore ClamAV Database:**

```bash
# Docker
docker volume rm clamav-db
docker run --rm -v clamav-db:/db -v $(pwd):/backup \
  alpine tar xzf /backup/clamav-db.tar.gz -C /db

# Kubernetes
kubectl cp ./clamav-db-backup clamav-pod:/var/lib/clamav
```

## Deployment Best Practices

### Security

1. **Network Isolation:**
   - Run ClamAV on internal network only
   - Use service-to-service authentication
   - Restrict API access with authentication

2. **Image Security:**
   - Use specific ClamAV versions (not `latest`)
   - Scan container images for vulnerabilities
   - Keep base images updated

3. **Container Security:**
   - Run as non-root user
   - Use read-only filesystems where possible
   - Implement network policies

### Performance Optimization

1. **Concurrency:**
   - Set appropriate `MaxFiles` for ClamAV (default: 10000)
   - Configure connection pooling in antivirus service
   - Use load balancing for multiple instances

2. **Database Updates:**
   - Adjust `FRESHCLAM_CHECKS` based on threat levels
   - Schedule updates during low-traffic periods
   - Monitor update success rates

3. **File Handling:**
   - Implement file size limits (`MAX_UPLOAD_SIZE_MB`)
   - Use temporary directories with cleanup
   - Stream large files instead of loading in memory

### Backup & Disaster Recovery

1. **Database Backups:**
   - Backup ClamAV database daily
   - Test restore procedures monthly
   - Store backups in separate locations

2. **Configuration Backups:**
   - Version control `.env` files (use .env.example)
   - Backup docker-compose.yml configurations
   - Document any manual configurations

3. **Recovery Procedures:**
   - Document recovery steps
   - Maintain runbooks for common failures
   - Test failover procedures

## Monitoring & Troubleshooting

### Key Metrics to Monitor

1. **Service Availability:**
   - Container uptime
   - Health check pass/fail rate
   - API response times

2. **Scan Performance:**
   - Scan completion time (p50, p95, p99)
   - Throughput (scans/minute)
   - Concurrent scan count

3. **Resource Usage:**
   - CPU utilization
   - Memory usage
   - Disk I/O (database updates)

4. **Database Health:**
   - Last update timestamp
   - Database size
   - Update success rate

### Common Issues & Solutions

| Issue | Symptom | Solution |
|-------|---------|----------|
| **ClamAV timeout** | "deadline exceeded" errors | Increase `CLAMAV_TIMEOUT_SECONDS` or `MaxScanSize` |
| **Database not updating** | Old signatures in logs | Check freshclam logs: `docker logs clamav` |
| **Disk space low** | Scans fail with "no space" | Increase persistent volume size |
| **Service not responding** | Connection refused errors | Verify `CLAMAV_ADDRESS` and network connectivity |
| **Memory exhaustion** | Container killed/OOMKilled | Increase container memory limits |

### Debugging Commands

```bash
# Check service status
docker-compose ps

# View service logs
docker-compose logs antivirus-service
docker-compose logs clamav

# Test ClamAV connectivity
docker exec antivirus-service telnet clamav 3310

# Check database status
docker exec clamav clamscan --version
docker exec clamav ls -lh /var/lib/clamav/

# Manual scan
docker exec clamav clamscan /data --recursive
```

### Log Analysis

Monitor logs for:
- `scanning` - Normal scan operations
- `FOUND` - Detected threats
- `freshclam` - Database updates
- `ERROR` - Service errors
- `timeout` - Performance issues

Example log patterns to alert on:
```
ERROR - Critical errors
FOUND - Potential threats
timeout - Performance degradation
failed update - Database update failures
```

---

**Last Updated:** 2026-04-18  
**Version:** 1.0  
**Maintainer:** DevOps Team

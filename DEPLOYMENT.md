# Deployment Guide

Complete instructions for deploying Sereni Antivirus ClamAV across different environments.

## Table of Contents

- [Quick Start](#quick-start)
- [Docker Compose Deployment](#docker-compose-deployment)
- [Kubernetes Deployment](#kubernetes-deployment)
- [Cloud Platforms](#cloud-platforms)
- [Post-Deployment Validation](#post-deployment-validation)

## Quick Start

### Prerequisites

- Docker & Docker Compose 3.8+
- 2 GB RAM minimum
- 1 GB disk space for ClamAV database
- Network access to ClamAV port (3310) and API port (8084)

### Local Development

```bash
# Clone repository
git clone https://github.com/aptlogica/sereni-antivirus-clamav.git
cd sereni-antivirus-clamav

# Create environment file
cp .env.example .env

# Start services
docker-compose up -d

# Verify services
docker-compose ps

# View logs
docker-compose logs -f antivirus-service

# Test health
curl http://localhost:8084/health
```

### Stop Services

```bash
docker-compose down

# Stop and remove volumes (fresh start)
docker-compose down -v
```

## Docker Compose Deployment

### Full Configuration Example

Create `.env` file with:

```bash
# Application Service
HOST=0.0.0.0
PORT=8084
ALLOWED_ORIGINS=https://api.example.com

# Antivirus Configuration
ANTIVIRUS_DRIVER=clamav
CLAMAV_ADDRESS=clamav:3310
CLAMAV_TIMEOUT_SECONDS=60
MAX_UPLOAD_SIZE_MB=256

# Volume Configuration (optional, defaults to .clamav-db in project root)
# CLAMAV_DB_PATH=/mnt/clamav-db
```

### Production docker-compose.yml Override

Create `docker-compose.prod.yml`:

```yaml
version: '3.8'

services:
  clamav:
    restart: always
    logging:
      driver: json-file
      options:
        max-size: "10m"
        max-file: "3"
    deploy:
      resources:
        limits:
          cpus: '2'
          memory: 2G
        reservations:
          cpus: '1'
          memory: 1G
    volumes:
      - /mnt/production/clamav-db:/var/lib/clamav

  antivirus-service:
    restart: always
    logging:
      driver: json-file
      options:
        max-size: "10m"
        max-file: "3"
    deploy:
      resources:
        limits:
          cpus: '2'
          memory: 512M
        reservations:
          cpus: '500m'
          memory: 256M
```

### Deploy Production

```bash
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d

# Monitor logs
docker-compose -f docker-compose.yml -f docker-compose.prod.yml logs -f

# Check status
docker-compose -f docker-compose.yml -f docker-compose.prod.yml ps
```

### Scaling with Docker Compose

For multiple service instances:

```yaml
version: '3.8'

services:
  clamav:
    image: clamav/clamav:1.4.5
    volumes:
      - clamav-db:/var/lib/clamav

  antivirus-service:
    build: .
    deploy:
      replicas: 3
    depends_on:
      - clamav
    environment:
      CLAMAV_ADDRESS: clamav:3310

volumes:
  clamav-db:
    driver: local
```

## Kubernetes Deployment

### Prerequisites

- Kubernetes 1.20+ cluster
- kubectl configured
- Persistent storage class available

### Namespace Setup

```bash
kubectl create namespace antivirus
kubectl label namespace antivirus env=production
```

### ConfigMap & Secrets

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: antivirus-config
  namespace: antivirus
data:
  HOST: "0.0.0.0"
  PORT: "8084"
  ANTIVIRUS_DRIVER: "clamav"
  CLAMAV_TIMEOUT_SECONDS: "60"
  MAX_UPLOAD_SIZE_MB: "256"

---
apiVersion: v1
kind: Secret
metadata:
  name: antivirus-secrets
  namespace: antivirus
type: Opaque
stringData:
  ALLOWED_ORIGINS: "https://api.example.com"
```

### Persistent Volume Claim

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: clamav-db-pvc
  namespace: antivirus
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: standard
  resources:
    requests:
      storage: 3Gi
```

### ClamAV Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: clamav
  namespace: antivirus
  labels:
    app: clamav
spec:
  replicas: 1
  selector:
    matchLabels:
      app: clamav
  template:
    metadata:
      labels:
        app: clamav
    spec:
      containers:
      - name: clamav
        image: clamav/clamav:1.4.5
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 3310
          name: clamd
        env:
        - name: CLAMD_STARTUP_TIMEOUT
          value: "180"
        - name: FRESHCLAM_CHECKS
          value: "12"
        - name: MaxScanSize
          value: "500M"
        - name: MaxFileSize
          value: "500M"
        - name: StreamMaxLength
          value: "500M"
        volumeMounts:
        - name: virus-db
          mountPath: /var/lib/clamav
        resources:
          requests:
            cpu: 500m
            memory: 512Mi
          limits:
            cpu: 2000m
            memory: 2Gi
        livenessProbe:
          exec:
            command:
            - /bin/sh
            - -c
            - clamdscan --version
          initialDelaySeconds: 30
          periodSeconds: 60
          timeoutSeconds: 10
          failureThreshold: 5
        readinessProbe:
          exec:
            command:
            - /bin/sh
            - -c
            - clamdscan --version
          initialDelaySeconds: 10
          periodSeconds: 10
          timeoutSeconds: 5
      volumes:
      - name: virus-db
        persistentVolumeClaim:
          claimName: clamav-db-pvc

---
apiVersion: v1
kind: Service
metadata:
  name: clamav
  namespace: antivirus
spec:
  selector:
    app: clamav
  ports:
  - port: 3310
    targetPort: 3310
    name: clamd
  type: ClusterIP
```

### Antivirus Service Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: antivirus-service
  namespace: antivirus
  labels:
    app: antivirus-service
spec:
  replicas: 3
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
  selector:
    matchLabels:
      app: antivirus-service
  template:
    metadata:
      labels:
        app: antivirus-service
    spec:
      containers:
      - name: antivirus-service
        image: sereni-antivirus:latest
        imagePullPolicy: Always
        ports:
        - containerPort: 8084
          name: api
        envFrom:
        - configMapRef:
            name: antivirus-config
        - secretRef:
            name: antivirus-secrets
        env:
        - name: CLAMAV_ADDRESS
          value: "clamav:3310"
        resources:
          requests:
            cpu: 250m
            memory: 256Mi
          limits:
            cpu: 1000m
            memory: 512Mi
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
        securityContext:
          readOnlyRootFilesystem: true
          runAsNonRoot: true
          runAsUser: 65534
          allowPrivilegeEscalation: false
          capabilities:
            drop:
            - ALL

---
apiVersion: v1
kind: Service
metadata:
  name: antivirus-service
  namespace: antivirus
  labels:
    app: antivirus-service
spec:
  selector:
    app: antivirus-service
  ports:
  - port: 8084
    targetPort: 8084
    name: api
  type: LoadBalancer

---
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: antivirus-service-hpa
  namespace: antivirus
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: antivirus-service
  minReplicas: 2
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
```

### Deploy to Kubernetes

```bash
# Create namespace and secrets
kubectl create namespace antivirus
kubectl apply -f configmap-secret.yaml

# Deploy ClamAV and persistence
kubectl apply -f pvc.yaml
kubectl apply -f clamav-deployment.yaml

# Deploy antivirus service
kubectl apply -f service-deployment.yaml

# Verify deployment
kubectl get all -n antivirus
kubectl get pvc -n antivirus

# Check pod logs
kubectl logs -n antivirus deployment/clamav -f
kubectl logs -n antivirus deployment/antivirus-service -f
```

## Cloud Platforms

### AWS ECS

```json
{
  "family": "sereni-antivirus",
  "networkMode": "awsvpc",
  "requiresCompatibilities": ["FARGATE"],
  "cpu": "1024",
  "memory": "2048",
  "containerDefinitions": [
    {
      "name": "clamav",
      "image": "clamav/clamav:1.4.5",
      "portMappings": [
        {
          "containerPort": 3310,
          "protocol": "tcp"
        }
      ],
      "mountPoints": [
        {
          "sourceVolume": "clamav-db",
          "containerPath": "/var/lib/clamav",
          "readOnly": false
        }
      ],
      "memory": 1024,
      "essential": true
    },
    {
      "name": "antivirus-service",
      "image": "<account-id>.dkr.ecr.<region>.amazonaws.com/sereni-antivirus:latest",
      "portMappings": [
        {
          "containerPort": 8084,
          "protocol": "tcp"
        }
      ],
      "environment": [
        {
          "name": "CLAMAV_ADDRESS",
          "value": "localhost:3310"
        }
      ],
      "memory": 1024,
      "essential": true
    }
  ],
  "volumes": [
    {
      "name": "clamav-db",
      "efsVolumeConfiguration": {
        "fileSystemId": "fs-xxxxx",
        "transitEncryption": "ENABLED"
      }
    }
  ]
}
```

### Google Cloud Run

```bash
# Build and push image
gcloud builds submit --tag gcr.io/PROJECT_ID/sereni-antivirus

# Deploy
gcloud run deploy sereni-antivirus \
  --image gcr.io/PROJECT_ID/sereni-antivirus \
  --memory 512Mi \
  --cpu 1 \
  --set-env-vars CLAMAV_ADDRESS=clamav:3310 \
  --port 8084
```

### Azure Container Instances

```yaml
apiVersion: '2021-09-01'
location: eastus
name: sereni-antivirus-group
properties:
  containers:
  - name: clamav
    properties:
      image: clamav/clamav:1.4.5
      resources:
        requests:
          cpu: 0.5
          memoryInGb: 0.5
      ports:
      - port: 3310
  - name: antivirus-service
    properties:
      image: sereni-antivirus:latest
      resources:
        requests:
          cpu: 0.5
          memoryInGb: 0.5
      ports:
      - port: 8084
      environmentVariables:
      - name: CLAMAV_ADDRESS
        value: localhost:3310
  osType: Linux
  ipAddress:
    type: Public
    ports:
    - port: 8084
      protocol: tcp
```

## Post-Deployment Validation

### Health Checks

```bash
# Check service health
curl -v http://localhost:8084/health

# Expected response
{
  "status": "healthy",
  "clamav_connected": true,
  "uptime_seconds": 3600
}
```

### ClamAV Connectivity

```bash
# Test ClamAV daemon
telnet localhost 3310

# Or use clamd command
clamdscan --version
```

### Test File Scanning

```bash
# Create test file
echo "This is a test" > test.txt

# Upload and scan
curl -F "file=@test.txt" http://localhost:8084/scan
```

### Performance Baseline

```bash
# Measure scan time
time curl -F "file=@largefile.zip" http://localhost:8084/scan

# Check concurrent requests
ab -n 100 -c 10 -p file.txt http://localhost:8084/scan
```

### Log Monitoring

```bash
# Docker Compose
docker-compose logs --tail=50 -f

# Kubernetes
kubectl logs -n antivirus -l app=antivirus-service -f

# AWS ECS
aws logs tail /ecs/sereni-antivirus --follow

# Azure
az container logs --name sereni-antivirus --resource-group mygroup
```

---

**Last Updated:** 2026-04-18  
**Version:** 1.0

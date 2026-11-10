# Deployment Documentation

## Prerequisites

- AlmaLinux 9.8 (or compatible Linux distribution)
- Go 1.26.7
- Ollama API access at the configured endpoint
- Sufficient permissions to bind to the desired port (default 8080)

## Installation

### 1. Clone the Repository

```bash
git clone <repository-url>
cd gentic
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Configure Environment Variables

Copy `.env.example` to `.env` and configure the required values:

```bash
cp .env.example .env
```

Edit `.env` with your configuration:

```
# Ollama (Inferência de IA)
OLLAMA_SERVER_ADDRESS=https://api.denck.sbs
OLLAMA_MODEL=muse-glimmer:30b-q4_K_M-dflash
OLLAMA_TIMEOUT=300

# VMware
VMWARE_VCENTER_URL=https://vcenter.example.com
VMWARE_USERNAME=admin
VMWARE_PASSWORD=senha123
VMWARE_INSECURE=false

# Graylog (Monitoring)
GRAYLOG_URL=https://graylog.example.com
GRAYLOG_TOKEN=seu-token-aqui
GRAYLOG_TIMEOUT=30

# Cloud (CloudStack)
CLOUDSTACK_URL=https://cloudstack.example.com
CLOUDSTACK_API_KEY=sua-chave-api
CLOUDSTACK_SECRET_KEY=suave-secret-key
CLOUDSTACK_TIMEOUT=30
```

### 4. Build the Project

```bash
go build -o infrastructure-ai-agent ./cmd/server/
```

### 5. Run the Server

```bash
./infrastructure-ai-agent
```

Or using Go directly:

```bash
go run ./cmd/server/
```

The server will start on `http://localhost:8080` by default.

## Running in Production

### Systemd Service

Create a systemd service file at `/etc/systemd/system/infrastructure-ai.service`:

```ini
[Unit]
Description=Infrastructure AI Agents Platform
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/bin/infrastructure-ai-agent
Restart=on-failure
RestartSecert=60
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
```

Enable and start the service:

```bash
systemctl enable infrastructure-ai
systemctl start infrastructure-ai
```

### Configuration Directories

- `/etc/infrastructure-ai/` - Configuration directory
- `/var/log/infrastructure-ai/` - Log directory
- `/var/lib/infrastructure-ai/` - Data directory

### Docker-Free Deployment

This platform does not require Docker or Podman. It runs natively on AlmaLinux 9.8.

## Post-Deployment Verification

```bash
# Check health endpoint
curl http://localhost:8080/health

# List available agents
curl http://localhost:8080/api/v1/agents

# Test VMware agent
curl -X POST http://localhost:8080/api/v1/agents/vmware/check

# Test Linux agent
curl -X POST http://localhost:8080/api/v1/agents/linux/system
```

## Uninstallation

```bash
systemctl stop infrastructure-ai
systemctl disable infrastructure-ai
rm /etc/systemd/system/infrastructure-ai.service
rm /usr/local/bin/infrastructure-ai-agent
```

## Backup and Recovery

- Make regular backups of the `.env` configuration file
- Backup log files from the configured log directory
- The platform does not persist state between restarts (stateless design)
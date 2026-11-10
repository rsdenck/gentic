#!/bin/bash
# deployment script for Infrastructure AI Agents
# Usage: ./deploy.sh [environment]
# Environments: local, production, docker

set -e

echo "=== Infrastructure AI Agents - Deployment Script ==="

# Determine environment
ENVIRONMENT="${1:-local}"
echo "Ambiente: $ENVIRONMENT"

# Load appropriate .env file
if [ "$ENVIRONMENT" = "production" ]; then
    if [ -f .env.production ]; then
        export $(cat .env.production | grep -v '^#' | xargs)
        echo "✅ Using .env.production"
    else
        echo "⚠️  .env.production not found, using .env"
        export $(cat .env | grep -v '^#' | xargs)
    fi
elif [ "$ENVIRONMENT" = "docker" ]; then
    echo "✅ Using docker-compose environment"
else
    export $(cat .env | grep -v '^#' | xargs)
    echo "✅ Using .env (local)"
fi

# Step 1: Build the project
echo ""
echo "=== Step 1: Building project ==="
go build -o infrastructure-ai-agents ./cmd/server/
echo "✅ Build concluído"

# Step 2: Create necessary directories
echo ""
echo "=== Step 2: Creating directories ==="
mkdir -p /var/log/gentic
mkdir -p /etc/gentic
echo "✅ Diretórios criados"

# Step 3: Set up systemd service (for local deployment)
if [ "$ENVIRONMENT" = "local" ] || [ "$ENVIRONMENT" = "production" ]; then
    echo ""
    echo "=== Step 3: Setting up systemd service ==="
    
    # Copy service file
    sudo cp /root/projects/gentic/infra-ai.service /etc/systemd/system/gentic.service 2>/dev/null || {
        # Create service file if it doesn't exist
        sudo cat > /etc/systemd/system/gentic.service << 'EOF'
[Unit]
Description=Infrastructure AI Agents Server
After=network.target
Documentation=https://github.com/rsdenck/gentic

[Service]
Type=simple
User=root
WorkingDirectory=/root/projects/gentic
ExecStart=/root/projects/gentic/infrastructure-ai-agents
Restart=always
RestartSec=5
EnvironmentFile=/etc/gentic/.env

[Install]
WantedBy=multi-user.target
EOF
    }
    
    sudo systemctl daemon-reload
    sudo systemctl enable gentic 2>/dev/null || echo "⚠️  Could not enable service (may need sudo)"
    echo "✅ Service configurado"
fi

# Step 4: Start the server
echo ""
echo "=== Step 4: Starting server ==="

if [ "$ENVIRONMENT" = "docker" ]; then
    echo "🐳 Starting with Docker Compose..."
    if [ -f docker-compose.yml ]; then
        docker-compose up -d 2>/dev/null || docker compose up -d 2>/dev/null || echo "⚠️  Docker Compose not available"
    else
        echo "⚠️  docker-compose.yml not found"
    fi
elif [ "$ENVIRONMENT" = "docker" ] && command -v docker &>/dev/null; then
    echo "🐳 Building and running Docker..."
    docker build -t gentic:latest . 2>/dev/null || echo "⚠️  Docker build failed"
    docker run -d -p 8080:8080 --name gentic-server -v "$(pwd)/.env":/app/.env:ro gentic:latest 2>/dev/null || echo "⚠️  Docker run failed"
else
    echo "▶️ Starting server in background..."
    nohup ./infrastructure-ai-agents > /tmp/gentic-server.log 2>&1 &
    echo "PID: $!"
    sleep 3
fi

# Step 5: Verify deployment
echo ""
echo "=== Step 5: Verifying deployment ==="

# Wait for server to start
sleep 3

# Health check
HEALTH=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/health)
if [ "$HEALTH" = "200" ]; then
    echo "✅ Server is running (HTTP 200)"
else
    echo "⚠️  Server health check returned: $HEALTH"
    echo "   Check logs: tail -20 /tmp/gentic-server.log 2>/dev/null || journalctl -u gentic -n 20"
fi

# API test
AGENTS=$(curl -s http://localhost:8080/api/v1/agents | head -c 100)
echo "📡 Agents disponíveis: ${AGENTS:0:50}..."

# Step 6: Summary
echo ""
echo "=========================================="
echo "=== DEPLOYMENT CONCLUÍDO ==="
echo "=========================================="
echo "🌐 Server: http://localhost:8080"
echo "🔍 Health: http://localhost:8080/health"
echo "📋 Agents: http://localhost:8080/api/v1/agents"
echo "📄 Logs: tail -f /tmp/gentic-server.log"
echo "=========================================="
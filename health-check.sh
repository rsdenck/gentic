#!/bin/bash
# health-check.sh - Verificação de saúde do servidor

SERVER="http://localhost:8080/health"
TIMEOUT=5

echo "=== Health Check do Infrastructure AI Agents ==="
echo "Verificando: $SERVER"
echo "Timestamp: $(date '+%Y-%m-%d %H:%M:%S')"
echo ""

# Verificar status HTTP
HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" --max-time $TIMEOUT "$SERVER")

# Verificar tempo de resposta
RESPONSE_TIME=$(curl -s -o /dev/null -w "%{time_total}" --max-time $TIMEOUT "$SERVER")
RESPONSE_TIME_FORMATED=$(printf "%.2f" "$RESPONSE_TIME")

# Verificar conteúdo do response
RESPONSE_BODY=$(curl -s --max-time $TIMEOUT "$SERVER")

echo "Resultado:"
echo "  Status HTTP: $HTTP_CODE"
echo "  Tempo de resposta: ${RESPONSE_TIME_FORMATED}s"
echo ""

# Avaliar saúde
if [ "$HTTP_CODE" = "200" ]; then
    echo "✅ Server SAUDÁVEL"
    echo "   Response: $RESPONSE_BODY"
    
    # Verificar se os agents estão respondendo
    AGENTS=$(curl -s http://localhost:8080/api/v1/agents | head -1)
    echo "   Agents: $AGENTS"
    
    exit 0
else
    echo "❌ Server NÃO SAUDÁVEL"
    echo "   O server pode estar reiniciando ou com problemas"
    
    # Tentar reiniciar se for um serviço systemd
    if systemctl status gentic >/dev/null 2>&1; then
        echo "   Tentando reiniciar service gentic..."
        systemctl restart gentic 2>/dev/null
    fi
    
    exit 1
fi
EOF

chmod +x /root/projects/gentic/health-check.sh
echo "✅ health-check.sh criado"
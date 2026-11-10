#!/bin/bash
# update.sh - Atualização do Infrastructure AI Agents

set -e

echo "=== Atualizando Infrastructure AI Agents ==="

# 1. Pull das mudanças mais recentes
echo ""
echo "=== 1. Pull do Git ==="
git pull origin master
echo "✅ Pull concluído"

# 2. Atualizar dependências
echo ""
echo "=== 2. Atualizando dependências ==="
go mod tidy
echo "✅ Dependências atualizadas"

# 3. Build do projeto
echo ""
echo "=== 3. Build do projeto ==="
go build -o infrastructure-ai-agents ./cmd/server/
echo "✅ Build concluído"

# 4. Verificar saúde após build
echo ""
echo "=== 4. Verificando saúde ==="
HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" --max-time 5 http://localhost:8080/health 2>/dev/null || echo "000")
if [ "$HTTP_CODE" = "200" ]; then
    echo "✅ Server ainda responding após build"
else
    echo "⚠️  Server não responding, reiniciando..."
    pkill -f "infrastructure-ai-agents" 2>/dev/null || true
    sleep 2
    nohup ./infrastructure-ai-agents > /tmp/gentic-server.log 2>&1 &
    sleep 3
fi

# 5. Atualizar tags (se houver novas releases)
echo ""
echo "=== 5. Verificando releases ==="
LATEST_TAG=$(git describe --tags --abbrev=0 2>/dev/null || echo "none")
echo "Última tag: $LATEST_TAG"

echo ""
echo "=== Atualização Concluída ==="
echo "📋 Próximos passos:"
echo "   - testar: go test ./..."
echo "   - health: curl http://localhost:8080/health"
echo "   - agents: curl http://localhost:8080/api/v1/agents"
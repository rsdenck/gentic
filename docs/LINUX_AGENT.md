# Linux Agent Documentation

## Overview

The Linux Agent is specialized in Linux system administration and diagnostics.

## Responsibilities

- Consultar CPU, memória e disco
- Consultar processos e serviços systemd
- Consultar logs do sistema
- Analisar processos com consumo elevado
- Diagnosticar problemas de recursos
- Identificar serviços indisponíveis

## Configuration

Required environment variables:

| Variable | Description |
|----------|-------------|
| (no agent-specific variables) | System metrics are collected via AI-generated commands |

## Security

- **Read-only policy**: Cannot modify system configuration, install packages, or change services
- All operations are read-only and for diagnostic purposes only
- No arbitrary command execution - tools use validated, restricted commands
- Audit logging of all operations

## Available Tools

- `get_system_info` - Collects system information (hostname, uptime, CPU, memory, disk)
- `get_cpu_usage` - Obtains CPU usage with per-core percentages
- `get_memory_usage` - Obtains memory usage information
- `get_disk_usage` - Obtains disk usage with partition info
- `get_top_processes` - Lists top processes by CPU or memory consumption
- `get_services` - Lists system services and their status
- `get_logs` - Retrieves system logs from recent period

## Workflows

### System Diagnosis

1. Recebe uma solicitação de diagnóstico do sistema
2. Coleta informações do sistema (hostname, uptime, recursos)
3. Consulta uso de CPU, memória e disco
4. Lista processos de maior consumo
5. Verifica status dos serviços
6. Retrieves system logs
7. Gera relatório estruturado com evidências e recomendações

### Ferramentas Autorizadas

Somente as ferramentas listadas acima podem ser executadas. O modelo não pode executar comandos arbitrários recebidos diretamente.

## Constraints

- Nenhum agente pode modificar configurações de sistema, instalar pacotes ou alterar serviços
- A execução de comandos tem lista de operações permitidas, timeout e registro de auditoria
- A execução de comandos deve ter lista de operações permitidas, timeout e registro de auditoria
- Não classificar alertas normais, flapping ou eventos esperados como incidentes sem evidências
- Todo conclusão deve distinguir: Evidência encontrada, Hipótese, Causa confirmada, Recomendação
- Não registrar tokens, senhas ou dados secretos nos logs de auditoria
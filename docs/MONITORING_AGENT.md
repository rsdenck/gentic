# Monitoring Agent Documentation

## Overview

The Monitoring Agent is specialized in monitoring and observability using Graylog 7.x.

## Responsibilities

- Pesquisar logs por período
- Pesquisar por hostname
- Pesquisar por aplicação
- Pesquisar por nível de severidade
- Consultar mensagens relacionadas a incidentes
- Correlação temporal
- Agrupamento de erros
- Geração de relatórios técnicos

## Configuration

Required environment variables:

| Variable | Description |
|----------|-------------|
| GRAYLOG_URL | Graylog server URL (required) |
| GRAYLOG_TOKEN | Graylog authentication token (required) |
| GRAYLOG_TIMEOUT | Request timeout in seconds | |

## Security

- **Read-only policy**: Cannot alter Graylog configuration, delete messages, or modify streams
- All operations are read-only and for analysis purposes only
- No alert classification without evidence
- Audit logging of all queries and analyses

## Available Tools

- `search_logs` - Search Graylog events by time period
- `search_by_hostname` - Search events by hostname
- `search_by_application` - Search events by application name
- `search_by_severity` - Search events by severity level
- `consult_incidents` - Consult incidents in Graylog by time period
- `correlate_events` - Correlate events across time and criteria
- `analyze_incident` - Perform structured incident analysis

## Workflows

### Análise de Incidente

1. Recebe uma solicitação de análise de incidente
2. Consulta incidentes no Graylog no período especificado
3. Correlaciona eventos relacionados
4. Utiliza o modelo para analisar os dados reunidos
5. Distinguir: Evidência encontrada, Hipótese, Causa confirmada, Recomendação
6. Gera relatório técnico estruturado
7. Registra execução na auditoria

### Correlação de Eventos

1. Recebe IDs de eventos para correlacionar
2. Busca eventos no Graylog
3. Analisa padrões temporais
4. Identifica causas potenciais
5. Gera relatório de correlação

## Constraints

- Não classificar alertas normais, flapping ou eventos esperados como incidentes sem evidências
- Toda conclusão deve distinguir: Evidência encontrada, Hipótese, Causa confirmada, Recomendação
- Não inventar evidências ou resultados de consultas
- Não classificar alertas normais, flapping ou eventos esperados como incidentes sem evidências
- Toda conclusão deve distinguir: Evidência encontrada, Hipótese, Causa confirmada, Recomendação
- Não registrar tokens, senhas ou dados secretos nos logs de auditoria
- Toda conclusão deve distinguir: Evidência encontrada, Hipótese, Causa confirmada, Recomendação
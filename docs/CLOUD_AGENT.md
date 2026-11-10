# Cloud Agent Documentation

## Overview

The Cloud Agent is specialized in Apache CloudStack private cloud management.

## Responsibilities

- Gestão de instâncias, volumes e redes
- Consultar templates e ofertas de serviço
- Monitorar recursos e quotas
- Consultar domínios e contas
- Estado dos recursos

## Configuration

Required environment variables:

| Variable | Description |
|----------|-------------|
| CLOUDSTACK_URL | CloudStack server URL (required) |
| CLOUDSTACK_API_KEY | CloudStack API key (required) |
| CLOUDSTACK_SECRET_KEY | CloudStack secret key (required) |
| CLOUDSTACK_TIMEOUT | Request timeout in seconds | |

## Security

- **Read-only policy**: Cannot create, modify, or delete resources
- All operations are read-only for diagnostic and monitoring purposes
- No resource creation, deletion, or configuration changes
- Audit logging of all queries

## Available Tools

- `list_instances` - Lists CloudStack instances
- `list_volumes` - Lists CloudStack volumes
- `list_networks` - Lists CloudStack networks
- `list_templates` - Lists CloudStack templates
- `list_service_offerings` - Lists CloudStack service offerings
- `list_resources` - Lists CloudStack domains and accounts
- `check_resource_state` - Checks the state of a CloudStack resource

## Workflows

### Diagnóstico de Infraestrutura CloudStack

1. Recebe uma solicitação de diagnóstico da infraestrutura
2. Lista instâncias, volumes e redes
3. Consulta ofertas de serviço e templates
4. Verifica recursos e quotas
5. Analisa estado dos recursos
6. Gera diagnóstico estruturado
7. Registra execução na auditoria

### Consultas de Recursos

1. Recebe consulta sobre recursos específicos
2. Verifica estado e detalhes do recurso
3. Consulta quotas e domínios
4. Gera relatório de status

## Constraints

- Não criar, modificar ou deletar recursos no CloudStack
- Todas as operações são somente leitura para fins de diagnóstico e monitoramento
- Não inventar dados quando a API não estiver disponível
- Consultar quotas e domínios de forma controlada
- Não registrar tokens, senhas ou dados secretos nos logs de auditoria
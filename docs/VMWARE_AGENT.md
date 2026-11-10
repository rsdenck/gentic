# VMware Agent Documentation

## Overview

The VMware Agent is specialized in managing and diagnosing VMware vSphere infrastructure, including vCenter and ESXi hosts.

## Responsibilities

- Consultar vCenter e ESXi
- Listar máquinas virtuais
- Consultar estado das VMs
- Consultar CPU e memória
- Consultar datastores
- Consultar snapshots
- Consultar hosts e clusters
- Analisar problemas de capacidade
- Gerar diagnósticos técnicos

## Configuration

Required environment variables:

| Variable | Description |
|----------|-------------|
| VMWARE_VCENTER_URL | vCenter server URL (required) |
| VMWARE_USERNAME | vCenter username (required) |
| VMWARE_PASSWORD | vCenter password (required) |
| VMWARE_INSECURE | Enable insecure TLS connections (default: false) |

## Security

- **Read-only policy**: Cannot restart machines, power off VMs, alter networks, or delete files
- All operations are read-only and for diagnostic purposes only
- API key authentication optional (via X-API-Key header)

## Available Tools

- `check_connection` - Verifies vCenter/ESXi connectivity and returns infrastructure info
- `list_vms` - Lists all virtual machines in vCenter
- `get_vm_status` - Consults the state of a specific VM
- `get_vm_resources` - Obtains CPU and memory info for a VM
- `analyze_capacity` - Analyzes capacity issues in the infrastructure
- `generate_diagnostics` - Generates technical diagnostics for VMware issues

## Workflows

### Diagnosis de VM

1. Recebe uma solicitação com problema descrito
2. Lista máquinas virtuais afetadas
3. Consulta estado e recursos das VMs
4. Analisa problemas de capacidade
5. Gera diagnóstico técnico estruturado
6. Registra execução na auditoria

### Ferramentas Autorizadas

Somente as ferramentas listadas acima podem ser executadas. O modelo não pode inventar novos comandos ou operações.

## Constraints

- Nenhum agente pode reiniciar máquinas, desligar VMs, alterar redes, apagar arquivos ou modificar configurações
- Todas as respostas devem ser baseadas em evidências reais consultadas
- Nenhum token, senha ou dado secreto deve ser registrado nos logs de auditoria
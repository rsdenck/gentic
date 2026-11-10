# Infrastructure AI Agents - Architecture

## Overview

The Infrastructure AI Agents platform consists of a modular Go service with a shared AI core and four independent agents, each with its own responsibilities, tools, and workflows.

## Architecture Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                        HTTP / REST API                          │
│                           ^                                       │
│                       (Port 8080)                                 │
└────────────────────┬───────────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────────────┐
│                       API Central                                 │
│  - JSON requests/responses                                      │
│  - Standardized error handling                                  │
│  - Request validation                                           │
│  - Timeout and context cancellation                             │
│  - Structured logging                                           │
│  - OpenAPI documentation                                        │
│  - RBAC and authentication                                      │
└────────────────────┬───────────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────────────┐
│                    SHARED AI CORE                                 │
│  - Firebase Genkit Go                                           │
│  - Ollama plugin (remote endpoint)                              │
│  - Model configuration via env vars                             │
│  - Timeout handling                                             │
│  - Connectivity validation                                      │
│  - Structured logs                                              │
│  - Health checks                                                │
└────────────────────┬───────────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────────────┐
│                    FOUR INDEPENDENT AGENTS                      │
│                                                                 │
│ 1. VMware Agent  │ 2. Linux Agent  │ 3. Monitoring Agent │
│  - vCenter/ESXi  │  - CPU, Memory      │  - Graylog 7.x      │
│  - VM management │  - Processes       │  - Log search       │
│  - Datastores    │  - Services        │  - Incident analysis│
│  - Snapshots     │  - Filesystem      │  - Pattern detection│
│  - Clusters      │  - Disk usage      │  - Reports          │
│                                                                 │
│ 4. Cloud Agent   │                                                 │
│  - Apache CloudStack                                 │
│  - Instances, Volumes                         │
│  - Networks, Templates                        │
│  - Service Offerings, Quotas                  │
└────────────────────┬───────────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────────────┐
│                    TOOLS & WORKFLOWS                            │
│                                                                 │
│  - Modular tools per agent                                      │
│  - Genkit workflows for diagnostics                             │
│  - Audit logging                                                │
│  - Read-only security policy                                    │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

## Core Components

### 1. AI Core (internal/ai/)
- Firebase Genkit Go initialization
- Ollama plugin configuration
- Remote endpoint: https://api.denck.sbs
- Model configuration via environment variables
- Timeout: configurable (default 300s)
- Connectivity validation
- Package-level singleton instance

### 2. API Central (internal/api/)
- REST endpoints for all four agents
- Health check endpoint
- Agents list endpoint
- Chat/query endpoints per agent
- JSON validation and standardized responses
- Read-only security policy

### 3. Agents (internal/agents/)
Each agent has:
- Agent struct with core reference
- Agent-specific tools (internal/tools/)
- Responsibilities and constraints
- Available tools list

### 4. Security
- Read-only policy by default
- API key authentication (optional)
- Audit logging of all requests
- No credential exposure in logs
- Permission separation per agent

## Environment Configuration

Required environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| OLLAMA_SERVER_ADDRESS | Ollama server URL | https://api.denck.sbs |
| OLLAMA_MODEL | AI model name | muse-glimmer:30b-q4_K_M-dflash |
| OLLAMA_TIMEOUT | Request timeout (seconds) | 300 |
| VMWARE_VCENTER_URL | vCenter URL | (required for VMware agent) |
| VMWARE_USERNAME | vCenter username | (required for VMware agent) |
| VMWARE_PASSWORD | vCenter password | (required for VMware agent) |
| VMWARE_INSECURE | Insecure TLS flag | false |
| GRAYLOG_URL | Graylog server URL | (required for monitoring agent) |
| GRAYLOG_TOKEN | Graylog authentication token | (required for monitoring agent) |
| CLOUDSTACK_URL | CloudStack server URL | (required for cloud agent) |
| CLOUDSTACK_API_KEY | CloudStack API key | (required for cloud agent) |
| CLOUDSTACK_SECRET_KEY | CloudStack secret key | (required for cloud agent) |
| APP_ENV | Application environment | development |

## Security Policy

- **Default: READ-ONLY** - No agent can perform destructive operations
- No VM restart, shutdown, or network modification
- No file deletion or configuration changes
- Audit logging of all requests and operations
- API key optional (required for production deployments)
- Input validation on all API endpoints
- Timeout protection on all operations
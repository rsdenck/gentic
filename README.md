# Infrastructure AI Agents

Platform for infrastructure monitoring and management using AI agents.

## Overview

The Infrastructure AI Agents platform consists of four independent agents that share a common AI core but have independent instructions, tools, and responsibilities:

1. **VMware Agent** - VMware vSphere infrastructure management
2. **Linux Agent** - Linux system administration and diagnostics
3. **Monitoring Agent** - Observability and log analysis (Graylog)
4. **Cloud Agent** - Apache CloudStack private cloud management

## Architecture

The platform is built as a modular Go service with:
- Shared AI core using Firebase Genkit with Ollama plugin
- Agent-specific implementations with validated tools
- REST API for agent communication
- Security-first design (read-only operations)
- Structured logging and audit trails

## Agents

- **VMware Agent**: vCenter/ESXi integration, VM management, datastore analysis
- **Linux Agent**: Process monitoring, service management, filesystem diagnostics
- **Monitoring Agent**: Graylog search, incident analysis, pattern detection
- **Cloud Agent**: CloudStack resource management, quota monitoring

## Getting Started

### Prerequisites

- Go 1.26.7
- Ollama API access at the configured endpoint

### Installation

1. Clone the repository
2. Copy `.env.example` to `.env` and configure values
3. Run: `go mod tidy`
4. Run: `go run ./cmd/server/`

### Configuration

See `.env.example` for all available environment variables.

### API

The platform provides a REST API at `http://localhost:8080` with endpoints for each agent.
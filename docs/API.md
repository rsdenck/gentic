# API Documentation

## Base URL

`http://localhost:8080`

## Authentication

Optional API key authentication via `X-API-Key` header. If configured, the key must match the value set in the `API_KEY` environment variable.

## Security Policy

- **Default: Read-only mode** - Write operations (POST, PUT, DELETE, PATCH) are forbidden unless explicitly allowed
- All operations are logged for audit purposes
- No credentials or sensitive data should be exposed in responses

## Endpoints

### Health Check

`GET /health`

Check the health of the platform.

**Response:**
```json
{
  "success": true,
  "data": {
    "status": "healthy|degraded|unknown",
    "uptime": "duration string",
    "version": "1.0.0",
    "environment": "development"
  }
}
```

### List Agents

`GET /api/v1/agents`

List all available agents with their status.

**Response:**
```json
{
  "success": true,
  "data": {
    "agents": [
      {"name": "vmware", "description": "VMware vSphere infrastructure agent", "status": "ready"},
      {"name": "linux", "description": "Linux system administration agent", "status": "ready"},
      {"name": "monitoring", "description": "Monitoring and observability agent", "status": "ready"},
      {"name": "cloud", "description": "Apache CloudStack private cloud agent", "status": "ready"}
    ]
  }
}
```

### VMware Agent

#### Chat

`POST /api/v1/agents/vmware/chat`

Send a prompt to the VMware agent.

**Request:**
```json
{"prompt": "Listar máquinas virtuais no vCenter"}
```

**Response:**
```json
{
  "success": true,
  "data": "Result from AI generation"
}
```

#### Check Connection

`POST /api/v1/agents/vmware/check`

Check vCenter connection.

**Response:**
```json
{
  "success": true,
  "data": {
    "status": "connected",
    "vcenter": "https://vcenter.example.com",
    "datastores": ["datastore1", "datastore2"],
    "hosts": ["esxi-host-1"],
    "clusters": ["cluster-1"],
    "virtualMachines": ["vm-ubuntu", "vm-centos"]
  }
}
```

#### List VMs

`POST /api/v1/agents/vmware/vms`

List virtual machines.

**Response:**
```json
{
  "success": true,
  "data": ["vm-ubuntu", "vm-centos", "vm-windows"]
}
```

#### VM Resources

`POST /api/v1/agents/vmware/resources`

Get CPU and memory for a VM.

**Request:**
```json
{"vm_name": "vm-ubuntu"}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "vm": "vm-ubuntu",
    "cpu_usage": "parsing required",
    "memory_usage": "requires detailed query"
  }
}
```

#### Capacity Analysis

`POST /api/v1/agents/vmware/capacity`

Analyze capacity issues.

**Response:**
```json
{
  "success": true,
  "data": {"analysis": "Result from AI generation"}
}
```

#### Generate Diagnostics

`POST /api/v1/agents/vmware/diagnostics`

Generate technical diagnostics.

**Request:**
```json
{"issue": "VM is slow"}
```

**Response:**
```json
{
  "success": true,
  "data": "Diagnostic text"
}
```

### Linux Agent

#### Chat

`POST /api/v1/agents/linux/chat`

Send a prompt to the Linux agent.

**Response:** Same structure as VMware chat

#### System Info

`POST /api/v1/agents/linux/system`

Get system information.

**Response:**
```json
{
  "success": true,
  "data": {
    "hostname": "server01",
    "uptime": "...",
    "cpu_model": "...",
    "cpu_count": 8,
    "memory_total": 16177260544,
    "memory_free": 4294967296,
    "disk_total": 500107862016,
    "disk_free": 300087113728,
    "load_average": "...",
    "platform": "linux",
    "platform_id": "linux",
    "kernel_version": "..."
  }
}
```

#### CPU Usage

`POST /api/v1/agents/linux/cpu`

Get CPU usage.

**Response:**
```json
{
  "success": true,
  "data": {
    "total_usage": 45.2,
    "idle_usage": 54.8,
    "core_count": 8
  }
}
```

#### Memory Usage

`POST /api/v1/agents/linux/memory`

Get memory usage.

**Response:**
```json
{
  "success": true,
  "data": {
    "total": 16177260544,
    "free": 4294967296,
    "used": 11882293248,
    "usage_percentage": 73.5
  }
}
```

#### Disk Usage

`POST /api/v1/agents/linux/disk`

Get disk usage.

**Response:**
```json
{
  "success": true,
  "data": {
    "total": 500107862016,
    "free": 300087113728,
    "used": 200020748288,
    "usage_percentage": 60.0
  }
}
```

#### Top Processes

`POST /api/v1/agents/linux/processes`

Get top processes.

**Request:**
```json
{"by": "cpu", "limit": 10}
```

**Response:**
```json
{
  "success": true,
  "data": [...]
}
```

#### Services

`POST /api/v1/agents/linux/services`

List services.

**Response:**
```json
{
  "success": true,
  "data": [...]
}
```

#### Logs

`POST /api/v1/agents/linux/logs`

Get logs.

**Request:**
```json
{"since": "1h", "limit": 100}
```

**Response:**
```json
{
  "success": true,
  "data": [...]
}
```

### Monitoring Agent

#### Chat

`POST /api/v1/agents/monitoring/chat`

Send a prompt to the monitoring agent.

**Response:** Same structure as VMware chat

#### Search Logs

`POST /api/v1/agents/monitoring/events`

Search Graylog events.

**Request:**
```json
{"start_time": "2024-01-01T00:00:00Z", "end_time": "2024-12-31T23:59:59Z"}
```

**Response:**
```json
{
  "success": true,
  "data": [...]
}
```

#### Search Incidents

`POST /api/v1/agents/monitoring/incidents`

Search incidents.

**Request:**
```json
{"start_time": "2024-01-01T00:00:00Z", "end_time": "2024-12-31T23:59:59Z"}
```

**Response:**
```json
{
  "success": true,
  "data": [...]
}
```

#### Analyze Incident

`POST /api/v1/agents/monitoring/analyze`

Analyze an incident.

**Request:**
```json
{"incident_id": "INC-001"}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "evidences": ["..."],
    "hypotheses": ["..."],
    "cause_confirmed": false,
    "recommendations": ["..."],
    "severity": "info"
  }
}
```

### Cloud Agent

#### Chat

`POST /api/v1/agents/cloud/chat`

Send a prompt to the cloud agent.

**Response:** Same structure as VMware chat

#### List Instances

`POST /api/v1/agents/cloud/instances`

List instances.

**Response:**
```json
{
  "success": true,
  "data": [...]
}
```

#### List Volumes

`POST /api/v1/agents/cloud/volumes`

List volumes.

**Response:**
```json
{
  "success": true,
  "data": [...]
}
```

#### List Networks

`POST /api/v1/agents/networks`

List networks.

**Response:**
```json
{
  "success": true,
  "data": [...]
}
```

#### List Templates

`POST /api/v1/agents/templates`

List templates.

**Response:**
```json
{
  "success": true,
  "data": [...]
}
```

#### List Resources

`POST /api/v1/agents/resources`

List resources (domains, accounts).

**Response:**
```json
{
  "success": true,
  "data": [...]
}
```

#### Check Resource State

`POST /api/v1/agents/state`

Check resource state.

**Request:**
```json
{"resource_type": "instance", "resource_id": "test-id"}
```

**Response:**
```json
{
  "success": true,
  "data": "running"
}
```

### Standard Error Response

```json
{
  "success": false,
  "error": {
    "code": "error_code",
    "message": "Error description"
  }
}
```

### Standard Success Response

```json
{
  "success": true,
  "data": {...}
}
```
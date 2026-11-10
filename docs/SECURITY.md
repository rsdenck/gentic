# Security Documentation

## Security Policy

### Default: Read-Only Mode

The Infrastructure AI Agents platform implements a read-only security policy by default. This means:

- **No destructive operations**: Agents cannot restart machines, power off VMs, alter networks, delete files, or modify configurations
- **Read-only operations only**: All agent interactions are limited to querying and reporting
- **Audit logging**: Every request and operation is logged for audit purposes
- **API key optional**: Authentication via X-API-Key header is optional for development but recommended for production

### Security Implementations

#### 1. Read-Only Policy Enforcement

The API middleware enforces read-only mode by:

- Intercepting all HTTP methods
- Blocking POST, PUT, DELETE, and PATCH requests in read-only mode
- Allowing GET requests and the agents list endpoint
- Optional API key validation via `X-API-Key` header

#### 2. Audit Logging

All requests are logged with the following information:

- Timestamp
- HTTP method
- Request path
- Client IP address
- User agent
- Success/failure status
- Error code (if applicable)

#### 3. Permission Structure

- **Separation of permissions by agent**: Each agent has its own set of authorized tools
- **No cross-agent tool execution**: The VMware agent cannot execute Linux tools and vice versa
- **Tool validation**: All tool arguments are validated before execution

#### 4. Credential Protection

- No tokens, passwords, or secret keys should be logged
- Environment variables should not be exposed in logs or error messages
- The Ollama API key is handled by the plugin, not passed in requests

#### 5. Input Validation

- All JSON payloads are validated before processing
- Required fields must be present
- String arguments are sanitized
- Type checking on all parameters

### Security Configuration

| Setting | Description | Default |
|---------|-------------|---------|
| `READ_ONLY_MODE` | Enforce read-only policy | true |
| `API_KEY` | Required API key for authentication | (empty - development mode) |
| `ALLOWED_AGENTS` | Comma-separated list of allowed agents | all four agents |

### Security Errors

| Error Code | Meaning | HTTP Status |
|------------|---------|-------------|
| `read_only_policy` | Write operation attempted in read-only mode | 403 Forbidden |
| `authentication_error` | Invalid or missing API key | 401 Unauthorized |
| `invalid_request` | Malformed JSON or missing required fields | 400 Bad Request |
| `agent_error` | Error executing agent tools | 500 Internal Server Request |
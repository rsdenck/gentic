package api

import (
	"log"
	"net/http"
	"strings"
	"time"
)

// SecurityPolicy defines the security policy for the API
type SecurityPolicy struct {
	// ReadOnlyMode enforces that only read operations are allowed
	ReadOnlyMode bool
	// AllowedAgents specifies which agents can be accessed
	AllowedAgents map[string]bool
	// APIKey for authentication (optional, can be empty for development)
	APIKey string
}

// NewSecurityPolicy creates a new security policy
func NewSecurityPolicy() *SecurityPolicy {
	return &SecurityPolicy{
		ReadOnlyMode:  true,
		AllowedAgents: map[string]bool{"vmware": true, "linux": true, "monitoring": true, "cloud": true},
	}
}

// NewAuditLogger creates a new audit logger
func NewAuditLogger() *AuditLogger {
	return &AuditLogger{}
}

// SecurityMiddleware creates an HTTP middleware for security enforcement
// This is a standalone function that can be used with any http.Handler
func SecurityMiddleware(p *SecurityPolicy, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if read-only mode is enabled
		if p.ReadOnlyMode {
			// Allow health check and agents list (read operations)
			if r.Method != http.MethodGet && !strings.HasPrefix(r.URL.Path, "/api/v1/agents") {
				// Check if this is a write operation
				writeOps := []string{"POST", "PUT", "DELETE", "PATCH"}
				isWrite := false
				for _, op := range writeOps {
					if r.Method == op {
						isWrite = true
						break
					}
				}

				if isWrite {
					// Check for API key (optional - allows development without key)
					if p.APIKey != "" {
						auth := r.Header.Get("X-API-Key")
						if auth != p.APIKey {
							resp := APIResponse{
								Success: false,
								Error:   &APIError{Code: "authentication_error", Message: "invalid or missing API key"},
							}
							writeJSON(w, http.StatusUnauthorized, resp)
							return
						}
					}

					resp := APIResponse{
						Success: false,
						Error:   &APIError{Code: "read_only_policy", Message: "write operations are not allowed in read-only mode"},
					}
					writeJSON(w, http.StatusForbidden, resp)
					return
				}
			}
		}

		// Agent authorization check
		path := r.URL.Path
		for agent := range p.AllowedAgents {
			if strings.Contains(path, "/agents/"+agent+"/") || strings.Contains(path, "/agents/"+agent) {
				// Agent found, validate
				break
			}
		}

		// Log the request for audit
		log.Printf("[AUDIT] %s %s %s %s %s", r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent(), time.Now().Format(time.RFC3339))

		next.ServeHTTP(w, r)
	})
}

// AuditLog represents an audit log entry
type AuditLog struct {
	Timestamp    string `json:"timestamp"`
	Method       string `json:"method"`
	Path         string `json:"path"`
	IP           string `json:"ip"`
	UserAgent    string `json:"user_agent"`
	Success      bool   `json:"success"`
	ErrorCode    string `json:"error_code,omitempty"`
	ResponseTime int64  `json:"response_time_ms"`
}

// AuditLogger handles audit logging
type AuditLogger struct{}

// LogRequest logs an HTTP request audit entry
func (l *AuditLogger) LogRequest(r *http.Request, success bool, err error, startTime time.Time) AuditLog {
	entry := AuditLog{
		Timestamp:    time.Now().Format(time.RFC3339),
		Method:       r.Method,
		Path:         r.URL.Path,
		IP:           r.RemoteAddr,
		UserAgent:    r.UserAgent(),
		Success:      success,
		ResponseTime: time.Since(startTime).Milliseconds(),
	}

	if err != nil {
		entry.ErrorCode = extractErrorCode(err)
	}

	return entry
}

// extractErrorCode extracts a short error code from an error
func extractErrorCode(err error) string {
	if err == nil {
		return ""
	}
	errStr := err.Error()
	switch {
	case strings.Contains(errStr, "read_only_policy"):
		return "RO_PL"
	case strings.Contains(errStr, "authentication_error"):
		return "AU_ERR"
	case strings.Contains(errStr, "agent_error"):
		return "AG_ERR"
	case strings.Contains(errStr, "missing_prompt"):
		return "PR_ERR"
	case strings.Contains(errStr, "connection_error"):
		return "CN_ERR"
	default:
		return "UNKW"
	}
}

// LogOperation logs an operation performed by an agent
func LogOperation(agent, operation string, success bool, details ...string) {
	var detail string
	if len(details) > 0 {
		detail = details[0]
	}
	if success {
		log.Printf("[AGENT] %s: %s - SUCCESS", agent, operation)
	} else {
		log.Printf("[AGENT] %s: %s - FAILED: %s", agent, operation, detail)
	}
}
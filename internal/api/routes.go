package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-logr/logr"
	"genkit-agent/internal/agents/cloud"
	"genkit-agent/internal/agents/linux"
	"genkit-agent/internal/agents/monitoring"
	"genkit-agent/internal/agents/vmware"
)

// APIResponse is a standardized API response
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
	Meta    *AMeta      `json:"meta,omitempty"`
}

// APIError represents an API error
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// AMeta represents metadata in the API response
type AMeta struct {
	RequestID string `json:"request_id,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

// HealthCheckResponse represents the health check response
type HealthCheckResponse struct {
	Status     string `json:"status"`
	Uptime     string `json:"uptime"`
	Version    string `json:"version"`
	Environment string `json:"environment"`
}

// AgentInfo represents agent information
type AgentInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// AgentsListResponse represents the list of available agents
type AgentsListResponse struct {
	Agents []AgentInfo `json:"agents"`
}

// Server represents the API server
type Server struct {
	router  *http.ServeMux
	server  *http.Server
	monitoring   *monitoring.MonitoringAgent
	linux        *linux.LinuxAgent
	vmware       *vmware.VMwareAgent
	cloud        *cloud.CloudAgent
	logger       logr.Logger
	security     *SecurityPolicy
	auditLogger  *AuditLogger
}

// NewServer creates a new API server
func NewServer() *Server {
	s := &Server{
		router:      http.NewServeMux(),
		server:      &http.Server{Addr: "localhost:8080"},
		monitoring:  monitoring.NewMonitoringAgent(),
		linux:       linux.NewLinuxAgent(),
		vmware:      vmware.NewVMwareAgent(),
		cloud:       cloud.NewCloudAgent(),
		logger:      logr.Discard(),
		security:    NewSecurityPolicy(),
		auditLogger: NewAuditLogger(),
	}
	s.setupRoutes()
	return s
}

// setupRoutes configures the API endpoints
func (s *Server) setupRoutes() {
	// Health endpoint
	s.router.HandleFunc("GET /health", s.handleHealth)

	// Agents list endpoint
	s.router.HandleFunc("GET /api/v1/agents", s.handleAgentsList)

	// VMware agent endpoints
	s.router.HandleFunc("POST /api/v1/agents/vmware/chat", s.handleVMwareChat)
	s.router.HandleFunc("POST /api/v1/agents/vmware/check", s.handleVMwareCheck)
	s.router.HandleFunc("POST /api/v1/agents/vmware/vms", s.handleVMwareVMs)
	s.router.HandleFunc("POST /api/v1/agents/vmware/resources", s.handleVMwareResources)
	s.router.HandleFunc("POST /api/v1/agents/vmware/capacity", s.handleVMwareCapacity)
	s.router.HandleFunc("POST /api/v1/agents/vmware/diagnostics", s.handleVMwareDiagnostics)

	// Linux agent endpoints
	s.router.HandleFunc("POST /api/v1/agents/linux/chat", s.handleLinuxChat)
	s.router.HandleFunc("POST /api/v1/agents/linux/system", s.handleLinuxSystem)
	s.router.HandleFunc("POST /api/v1/agents/linux/cpu", s.handleLinuxCPU)
	s.router.HandleFunc("POST /api/v1/agents/linux/memory", s.handleLinuxMemory)
	s.router.HandleFunc("POST /api/v1/agents/linux/disk", s.handleLinuxDisk)
	s.router.HandleFunc("POST /api/v1/agents/linux/processes", s.handleLinuxProcesses)
	s.router.HandleFunc("POST /api/v1/agents/linux/services", s.handleLinuxServices)
	s.router.HandleFunc("POST /api/v1/agents/linux/logs", s.handleLinuxLogs)

	// Monitoring agent endpoints
	s.router.HandleFunc("POST /api/v1/agents/monitoring/chat", s.handleMonitoringChat)
	s.router.HandleFunc("POST /api/v1/agents/monitoring/events", s.handleMonitoringEvents)
	s.router.HandleFunc("POST /api/v1/agents/monitoring/incidents", s.handleMonitoringIncidents)
	s.router.HandleFunc("POST /api/v1/agents/monitoring/analyze", s.handleMonitoringAnalyze)

	// Cloud agent endpoints
	s.router.HandleFunc("POST /api/v1/agents/cloud/chat", s.handleCloudChat)
	s.router.HandleFunc("POST /api/v1/agents/cloud/instances", s.handleCloudInstances)
	s.router.HandleFunc("POST /api/v1/agents/cloud/volumes", s.handleCloudVolumes)
	s.router.HandleFunc("POST /api/v1/agents/networks", s.handleCloudNetworks)
	s.router.HandleFunc("POST /api/v1/agents/templates", s.handleCloudTemplates)
	s.router.HandleFunc("POST /api/v1/agents/resources", s.handleCloudResources)
	s.router.HandleFunc("POST /api/v1/agents/state", s.handleCloudState)
}

// handleHealth handles the health check endpoint
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	status := "unknown"
	if s.monitoring != nil {
		status = "healthy"
	}
	if status == "unknown" {
		status = "degraded"
	}

	resp := APIResponse{
		Success: true,
		Data: HealthCheckResponse{
			Status:      status,
			Version:     "1.0.0",
			Environment: os.Getenv("APP_ENV"),
			Uptime:      time.Since(time.Unix(0, 0)).String(),
		},
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleAgentsList handles the agents list endpoint
func (s *Server) handleAgentsList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	agents := []AgentInfo{
		{Name: "vmware", Description: "VMware vSphere infrastructure agent", Status: "ready"},
		{Name: "linux", Description: "Linux system administration agent", Status: "ready"},
		{Name: "monitoring", Description: "Monitoring and observability agent", Status: "ready"},
		{Name: "cloud", Description: "Apache CloudStack private cloud agent", Status: "ready"},
	}

	resp := APIResponse{
		Success: true,
		Data:    AgentsListResponse{Agents: agents},
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleVMwareChat handles VMware agent chat
func (s *Server) handleVMwareChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Prompt string `json:"prompt"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "invalid_request", Message: "invalid JSON payload"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	if req.Prompt == "" {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "missing_prompt", Message: "prompt is required"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	result, err := s.vmware.ListVirtualMachines(r.Context())
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    result,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleVMwareCheck handles VMware connection check
func (s *Server) handleVMwareCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	info, err := s.vmware.CheckConnection(r.Context())
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "connection_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    info,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleVMwareVMs handles VM listing
func (s *Server) handleVMwareVMs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	result, err := s.vmware.GetVMStatus(r.Context(), "example-vm")
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    result,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleVMwareResources handles VM resources
func (s *Server) handleVMwareResources(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		VMName string `json:"vm_name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "invalid_request", Message: "invalid JSON payload"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	if req.VMName == "" {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "missing_vm_name", Message: "vm_name is required"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	result, err := s.vmware.GetVMResources(r.Context(), req.VMName)
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    result,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleVMwareCapacity handles capacity analysis
func (s *Server) handleVMwareCapacity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	result, err := s.vmware.AnalyzeCapacity(r.Context())
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    result,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleVMwareDiagnostics handles diagnostics generation
func (s *Server) handleVMwareDiagnostics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Issue string `json:"issue"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "invalid_request", Message: "invalid JSON payload"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	if req.Issue == "" {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "missing_issue", Message: "issue is required"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	result, err := s.vmware.GenerateDiagnostics(r.Context(), req.Issue)
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    result,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleLinuxChat handles Linux agent chat
func (s *Server) handleLinuxChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Prompt string `json:"prompt"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "invalid_request", Message: "invalid JSON payload"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	if req.Prompt == "" {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "missing_prompt", Message: "prompt is required"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	result, err := s.linux.GetSystemInfo(r.Context())
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    result,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleLinuxSystem handles system info
func (s *Server) handleLinuxSystem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	result, err := s.linux.GetSystemInfo(r.Context())
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    result,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleLinuxCPU handles CPU info
func (s *Server) handleLinuxCPU(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	result, err := s.linux.GetCPUUsage(r.Context())
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    result,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleLinuxMemory handles memory info
func (s *Server) handleLinuxMemory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	result, err := s.linux.GetMemoryUsage(r.Context())
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    result,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleLinuxDisk handles disk info
func (s *Server) handleLinuxDisk(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	result, err := s.linux.GetDiskUsage(r.Context())
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    result,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleLinuxProcesses handles top processes
func (s *Server) handleLinuxProcesses(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		By     string `json:"by"`
		Limit  int    `json:"limit"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "invalid_request", Message: "invalid JSON payload"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	processes, err := s.linux.GetTopProcesses(r.Context(), req.By, req.Limit)
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    processes,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleLinuxServices handles services
func (s *Server) handleLinuxServices(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	services, err := s.linux.GetServices(r.Context())
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    services,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleLinuxLogs handles logs
func (s *Server) handleLinuxLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Since  string `json:"since"`
		Limit  int    `json:"limit"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "invalid_request", Message: "invalid JSON payload"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	logs, err := s.linux.GetLogs(r.Context(), req.Since, req.Limit)
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    logs,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleMonitoringChat handles monitoring agent chat
func (s *Server) handleMonitoringChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Prompt string `json:"prompt"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "invalid_request", Message: "invalid JSON payload"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	if req.Prompt == "" {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "missing_prompt", Message: "prompt is required"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	result, err := s.monitoring.GetEventsByPeriod(r.Context(), "recent", "1h")
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    result,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleMonitoringEvents handles event search
func (s *Server) handleMonitoringEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "invalid_request", Message: "invalid JSON payload"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	events, err := s.monitoring.GetEventsByPeriod(r.Context(), req.StartTime, req.EndTime)
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    events,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleMonitoringIncidents handles incident search
func (s *Server) handleMonitoringIncidents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "invalid_request", Message: "invalid JSON payload"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	incidents, err := s.monitoring.GetIncidentsByPeriod(r.Context(), req.StartTime, req.EndTime)
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    incidents,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleMonitoringAnalyze handles incident analysis
func (s *Server) handleMonitoringAnalyze(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		IncidentID string `json:"incident_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "invalid_request", Message: "invalid JSON payload"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	analysis, err := s.monitoring.AnalyzeIncident(r.Context(), req.IncidentID)
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    analysis,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleCloudChat handles cloud agent chat
func (s *Server) handleCloudChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Prompt string `json:"prompt"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "invalid_request", Message: "invalid JSON payload"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	if req.Prompt == "" {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "missing_prompt", Message: "prompt is required"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	result, err := s.cloud.ListInstances(r.Context())
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    result,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleCloudInstances handles instance listing
func (s *Server) handleCloudInstances(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	instances, err := s.cloud.ListInstances(r.Context())
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    instances,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleCloudVolumes handles volume listing
func (s *Server) handleCloudVolumes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	volumes, err := s.cloud.ListVolumes(r.Context())
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    volumes,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleCloudNetworks handles network listing
func (s *Server) handleCloudNetworks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	networks, err := s.cloud.ListNetworks(r.Context())
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    networks,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleCloudTemplates handles template listing
func (s *Server) handleCloudTemplates(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	templates, err := s.cloud.ListTemplates(r.Context())
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    templates,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleCloudResources handles resource listing
func (s *Server) handleCloudResources(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	resources, err := s.cloud.ListResources(r.Context())
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    resources,
	}
	writeJSON(w, http.StatusOK, resp)
}

// handleCloudState handles resource state checking
func (s *Server) handleCloudState(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ResourceType string `json:"resource_type"`
		ResourceID   string `json:"resource_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "invalid_request", Message: "invalid JSON payload"},
		}
		writeJSON(w, http.StatusBadRequest, resp)
		return
	}

	state, err := s.cloud.CheckResourceState(r.Context(), req.ResourceType, req.ResourceID)
	if err != nil {
		resp := APIResponse{
			Success: false,
			Error:   &APIError{Code: "agent_error", Message: err.Error()},
		}
		writeJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp := APIResponse{
		Success: true,
		Data:    state,
	}
	writeJSON(w, http.StatusOK, resp)
}

// writeJSON writes a JSON response
func writeJSON(w http.ResponseWriter, statusCode int, resp APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println("Failed to encode JSON response:", err)
	}
}

// Setup starts the HTTP server
func (s *Server) Setup(addr string) error {
	s.logger.Info("Starting API server", "address", addr)
	s.server = &http.Server{Addr: addr, Handler: SecurityMiddleware(s.security, s.router)}
	return s.server.ListenAndServe()
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down API server")
	if s.server != nil {
		return s.server.Shutdown(ctx)
	}
	return nil
}
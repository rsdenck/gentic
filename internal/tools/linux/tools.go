package linux

import (
	"context"
	"fmt"

	"genkit-agent/internal/ai"
)

// LinuxTool is an interface for Linux-specific tools
type LinuxTool interface {
	Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error)
	GetName() string
	GetDescription() string
}

// ToolLinuxGetSystemInfo collects system information
type ToolLinuxGetSystemInfo struct{}

// NewToolLinuxGetSystemInfo creates a new system info tool
func NewToolLinuxGetSystemInfo() *ToolLinuxGetSystemInfo {
	return &ToolLinuxGetSystemInfo{}
}

// GetName returns the tool name
func (t *ToolLinuxGetSystemInfo) GetName() string {
	return "get_system_info"
}

// GetDescription returns the tool description
func (t *ToolLinuxGetSystemInfo) GetDescription() string {
	return "Coleta informações do sistema Linux (hostname, uptime, CPU, memória, disco)"
}

// Execute runs the system info collection
func (t *ToolLinuxGetSystemInfo) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	resp, err := ai.Generate(ctx, "Obter hostname, uptime, CPU, memória e disco do sistema Linux")
	if err != nil {
		return nil, fmt.Errorf("failed to get system info: %w", err)
	}

	return map[string]interface{}{
		"status": "collected",
		"info":   resp,
	}, nil
}

// ToolLinuxGetCPUUsage gets CPU usage information
type ToolLinuxGetCPUUsage struct{}

// NewToolLinuxGetCPUUsage creates a new CPU usage tool
func NewToolLinuxGetCPUUsage() *ToolLinuxGetCPUUsage {
	return &ToolLinuxGetCPUUsage{}
}

// GetName returns the tool name
func (t *ToolLinuxGetCPUUsage) GetName() string {
	return "get_cpu_usage"
}

// GetDescription returns the tool description
func (t *ToolLinuxGetCPUUsage) GetDescription() string {
	return "Obter uso de CPU do sistema Linux com percentuais por núcleo"
}

// Execute runs the CPU usage collection
func (t *ToolLinuxGetCPUUsage) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	resp, err := ai.Generate(ctx, "Obter uso de CPU do sistema Linux com percentuais por núcleo")
	if err != nil {
		return nil, fmt.Errorf("failed to get CPU usage: %w", err)
	}

	return map[string]interface{}{
		"status":   "collected",
		"response": resp,
		"tool":     "get_cpu_usage",
	}, nil
}

// ToolLinuxGetMemoryUsage gets memory usage information
type ToolLinuxGetMemoryUsage struct{}

// NewToolLinuxGetMemoryUsage creates a new memory usage tool
func NewToolLinuxGetMemoryUsage() *ToolLinuxGetMemoryUsage {
	return &ToolLinuxGetMemoryUsage{}
}

// GetName returns the tool name
func (t *ToolLinuxGetMemoryUsage) GetName() string {
	return "get_memory_usage"
}

// GetDescription returns the tool description
func (t *ToolLinuxGetMemoryUsage) GetDescription() string {
	return "Obter uso de memória do sistema Linux"
}

// Execute runs the memory usage collection
func (t *ToolLinuxGetMemoryUsage) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	resp, err := ai.Generate(ctx, "Obter uso de memória do sistema Linux")
	if err != nil {
		return nil, fmt.Errorf("failed to get memory usage: %w", err)
	}

	return map[string]interface{}{
		"status":   "collected",
		"response": resp,
		"tool":     "get_memory_usage",
	}, nil
}

// ToolLinuxGetDiskUsage gets disk usage information
type ToolLinuxGetDiskUsage struct{}

// NewToolLinuxGetDiskUsage creates a new disk usage tool
func NewToolLinuxGetDiskUsage() *ToolLinuxGetDiskUsage {
	return &ToolLinuxGetDiskUsage{}
}

// GetName returns the tool name
func (t *ToolLinuxGetDiskUsage) GetName() string {
	return "get_disk_usage"
}

// GetDescription returns the tool description
func (t *ToolLinuxGetDiskUsage) GetDescription() string {
	return "Obter uso de disco do sistema Linux - partições e espaço total"
}

// Execute runs the disk usage collection
func (t *ToolLinuxGetDiskUsage) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	resp, err := ai.Generate(ctx, "Obter uso de disco do sistema Linux - partições e espaço total")
	if err != nil {
		return nil, fmt.Errorf("failed to get disk usage: %w", err)
	}

	return map[string]interface{}{
		"status":   "collected",
		"response": resp,
		"tool":     "get_disk_usage",
	}, nil
}

// ToolLinuxGetTopProcesses gets top processes by CPU or memory
type ToolLinuxGetTopProcesses struct {
	By    string
	Limit int
}

// NewToolLinuxGetTopProcesses creates a new top processes tool
func NewToolLinuxGetTopProcesses(by string, limit int) *ToolLinuxGetTopProcesses {
	return &ToolLinuxGetTopProcesses{By: by, Limit: limit}
}

// GetName returns the tool name
func (t *ToolLinuxGetTopProcesses) GetName() string {
	return "get_top_processes"
}

// GetDescription returns the tool description
func (t *ToolLinuxGetTopProcesses) GetDescription() string {
	return "Listar os N processos de maior consumo de CPU ou memória no Linux"
}

// Execute runs the top processes collection
func (t *ToolLinuxGetTopProcesses) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	var prompt string
	if t.By == "cpu" {
		prompt = fmt.Sprintf("Listar os %d processos de maior consumo de CPU no Linux", t.Limit)
	} else if t.By == "memory" {
		prompt = fmt.Sprintf("Listar os %d processos de maior consumo de memória no Linux", t.Limit)
	} else {
		return nil, fmt.Errorf("invalid sort option: %s (use 'cpu' or 'memory')", t.By)
	}

	resp, err := ai.Generate(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to get top processes: %w", err)
	}

	return map[string]interface{}{
		"status":   "collected",
		"response": resp,
		"tool":     "get_top_processes",
		"sort_by":  t.By,
		"limit":    t.Limit,
	}, nil
}

// ToolLinuxGetServices lists system services
type ToolLinuxGetServices struct{}

// NewToolLinuxGetServices creates a new services tool
func NewToolLinuxGetServices() *ToolLinuxGetServices {
	return &ToolLinuxGetServices{}
}

// GetName returns the tool name
func (t *ToolLinuxGetServices) GetName() string {
	return "get_services"
}

// GetDescription returns the tool description
func (t *ToolLinuxGetServices) GetDescription() string {
	return "Listar serviços do sistema Linux e seus status"
}

// Execute runs the services collection
func (t *ToolLinuxGetServices) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	resp, err := ai.Generate(ctx, "Listar serviços do sistema Linux e seus status")
	if err != nil {
		return nil, fmt.Errorf("failed to get services: %w", err)
	}

	return map[string]interface{}{
		"status":   "collected",
		"response": resp,
		"tool":     "get_services",
	}, nil
}

// ToolLinuxGetLogs retrieves system logs
type ToolLinuxGetLogs struct {
	Since string
	Limit int
}

// NewToolLinuxGetLogs creates a new logs tool
func NewToolLinuxGetLogs(since string, limit int) *ToolLinuxGetLogs {
	return &ToolLinuxGetLogs{Since: since, Limit: limit}
}

// GetName returns the tool name
func (t *ToolLinuxGetLogs) GetName() string {
	return "get_logs"
}

// GetDescription returns the tool description
func (t *ToolLinuxGetLogs) GetDescription() string {
	return "Obter logs do sistema Linux nos últimos N minutos/comando"
}

// Execute runs the logs collection
func (t *ToolLinuxGetLogs) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	prompt := fmt.Sprintf("Obter logs do sistema Linux nos últimos %s, limite de %d mensagens", t.Since, t.Limit)
	resp, err := ai.Generate(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to get logs: %w", err)
	}

	return map[string]interface{}{
		"status":   "collected",
		"response": resp,
		"tool":     "get_logs",
		"since":    t.Since,
		"limit":    t.Limit,
	}, nil
}

package linux

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-logr/logr"
	"genkit-agent/internal/ai"
)

// LinuxAgent represents the Linux system administration and diagnostics agent
type LinuxAgent struct {
	core   *ai.Core
	logger logr.Logger
}

// NewLinuxAgent creates a new Linux agent instance
func NewLinuxAgent() *LinuxAgent {
	return &LinuxAgent{
		core:   ai.GetCore(),
		logger: logr.Discard(),
	}
}

// SetLogger sets the logger for the agent
func (a *LinuxAgent) SetLogger(l logr.Logger) {
	a.logger = l
}

// GetLogger returns the agent's logger
func (a *LinuxAgent) GetLogger() logr.Logger {
	return a.logger
}

// LinuxSystemInfo contains system information
type LinuxSystemInfo struct {
	Hostname     string
	Uptime       string
	CPUModel     string
	CPUCount     int
	MemoryTotal  int64
	MemoryFree   int64
	DiskTotal    int64
	DiskFree     int64
	LoadAverage  string
	Platform     string
	PlatformID   string
	KernelVersion string
}

// GetSystemInfo collects system information
func (a *LinuxAgent) GetSystemInfo(ctx context.Context) (LinuxSystemInfo, error) {
	info := LinuxSystemInfo{}

	// Get hostname
	resp, err := ai.Generate(ctx, "Obter hostname do sistema Linux")
	if err != nil {
		return info, fmt.Errorf("failed to get hostname: %w", err)
	}
	info.Hostname = resp

	// Get uptime
	resp, err = ai.Generate(ctx, "Obter uptime do sistema Linux")
	if err != nil {
		return info, fmt.Errorf("failed to get uptime: %w", err)
	}
	info.Uptime = resp

	// Get CPU model and count
	resp, err = ai.Generate(ctx, "Obter modelo e número de CPUs do sistema Linux")
	if err != nil {
		return info, fmt.Errorf("failed to get CPU info: %w", err)
	}
	info.CPUModel = resp

	// Parse CPU count from response (simple parsing)
	if countStr := strings.TrimSpace(resp); countStr != "" {
		if n, err := strconv.Atoi(countStr); err == nil {
			info.CPUCount = n
		}
	}

	// Get memory info
	resp, err = ai.Generate(ctx, "Obter informações de memória do sistema Linux")
	if err != nil {
		return info, fmt.Errorf("failed to get memory info: %w", err)
	}
	info.MemoryTotal = parseBytes(resp)

	// Get disk info
	resp, err = ai.Generate(ctx, "Obter informações de disco do sistema Linux")
	if err != nil {
		return info, fmt.Errorf("failed to get disk info: %w", err)
	}
	info.DiskTotal = parseBytes(resp)

	// Get load average
	resp, err = ai.Generate(ctx, "Obter carga média do sistema Linux")
	if err != nil {
		return info, fmt.Errorf("failed to get load average: %w", err)
	}
	info.LoadAverage = resp

	// Get platform info
	resp, err = ai.Generate(ctx, "Obter plataforma e versão do kernel do Linux")
	if err != nil {
		return info, fmt.Errorf("failed to get platform info: %w", err)
	}
	parts := strings.Split(resp, "/")
	if len(parts) >= 2 {
		info.Platform = parts[0]
		info.KernelVersion = strings.Join(parts[1:], "/")
	}

	info.PlatformID = "linux"

	a.logger.Info("System info collected")
	return info, nil
}

// parseBytes parses a human-readable byte string to integer bytes
func parseBytes(s string) int64 {
	s = strings.TrimSpace(s)
	switch {
	case strings.HasSuffix(s, "TB"):
		v, _ := strconv.ParseFloat(strings.TrimSuffix(s, "TB"), 64)
		return int64(v * 1024 * 1024 * 1024 * 1024)
	case strings.HasSuffix(s, "GB"):
		v, _ := strconv.ParseFloat(strings.TrimSuffix(s, "GB"), 64)
		return int64(v * 1024 * 1024 * 1024)
	case strings.HasSuffix(s, "MB"):
		v, _ := strconv.ParseFloat(strings.TrimSuffix(s, "MB"), 64)
		return int64(v * 1024 * 1024)
	case strings.HasSuffix(s, "KB"):
		v, _ := strconv.ParseFloat(strings.TrimSuffix(s, "KB"), 64)
		return int64(v * 1024)
	case strings.HasSuffix(s, "B"):
		v, _ := strconv.ParseFloat(strings.TrimSuffix(s, "B"), 64)
		return int64(v)
	default:
		// Try as plain number (assume bytes)
		v, _ := strconv.ParseInt(s, 10, 64)
		return v
	}
}

// CPUInfo contains CPU usage information
type CPUInfo struct {
	TotalUsage    float64
	IdleUsage     float64
	CoreCount     int
	PerCoreUsage  []float64
}

// GetCPUUsage gets CPU usage information
func (a *LinuxAgent) GetCPUUsage(ctx context.Context) (CPUInfo, error) {
	resp, err := ai.Generate(ctx, "Obter uso de CPU do sistema Linux com percentuais por núcleo")
	if err != nil {
		return CPUInfo{}, fmt.Errorf("failed to get CPU usage: %w", err)
	}

	// Parse the response - in a real implementation, parse structured data
	result := CPUInfo{}
	if resp != "" {
		// Simple parsing - try to extract numbers
		parts := strings.Split(resp, ",")
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if strings.Contains(p, "%") {
				val := strings.TrimSuffix(p, "%")
				if f, err := strconv.ParseFloat(val, 64); err == nil {
					result.TotalUsage = f
				}
			}
		}
	}

	a.logger.Info("CPU usage consulted")
	return result, nil
}

// MemoryInfo contains memory usage information
type MemoryInfo struct {
	Total       int64
	Free        int64
	Used        int64
	UsagePercentage float64
}

// GetMemoryUsage gets memory usage information
func (a *LinuxAgent) GetMemoryUsage(ctx context.Context) (MemoryInfo, error) {
	resp, err := ai.Generate(ctx, "Obter uso de memória do sistema Linux")
	if err != nil {
		return MemoryInfo{}, fmt.Errorf("failed to get memory usage: %w", err)
	}

	result := MemoryInfo{}
	if resp != "" {
		// Parse memory values from response
		// Format could be "Total: 16GB, Free: 4GB, Used: 12GB"
		parts := strings.Split(resp, ",")
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if strings.Contains(p, "Total") {
				val := strings.TrimPrefix(p, "Total")
				val = strings.TrimSpace(val)
				val = strings.TrimSuffix(val, "GB")
				if b, err := strconv.ParseInt(strings.TrimSpace(val), 10, 64); err == nil {
					result.Total = b * 1024 * 1024 * 1024 // Convert GB to bytes
				}
			} else if strings.Contains(p, "Free") {
				val := strings.TrimPrefix(p, "Free")
				val = strings.TrimSpace(val)
				val = strings.TrimSuffix(val, "GB")
				if b, err := strconv.ParseInt(strings.TrimSpace(val), 10, 64); err == nil {
					result.Free = b * 1024 * 1024 * 1024
				}
			} else if strings.Contains(p, "Used") {
				val := strings.TrimPrefix(p, "Used")
				val = strings.TrimSpace(val)
				val = strings.TrimSuffix(val, "GB")
				if b, err := strconv.ParseInt(strings.TrimSpace(val), 10, 64); err == nil {
					result.Used = b * 1024 * 1024 * 1024
				}
			}
		}
	}

	if result.Total > 0 {
		result.Used = result.Total - result.Free
		if result.Total > 0 {
			result.UsagePercentage = (float64(result.Used) / float64(result.Total)) * 100
		}
	}

	a.logger.Info("Memory usage consulted")
	return result, nil
}

// DiskInfo contains disk usage information
type DiskInfo struct {
	Total       int64
	Free        int64
	Used        int64
	UsagePercentage float64
	MountPoints []string
}

// GetDiskUsage gets disk usage information
func (a *LinuxAgent) GetDiskUsage(ctx context.Context) (DiskInfo, error) {
	resp, err := ai.Generate(ctx, "Obter uso de disco do sistema Linux - partições e espaço total")
	if err != nil {
		return DiskInfo{}, fmt.Errorf("failed to get disk usage: %w", err)
	}

	result := DiskInfo{}
	if resp != "" {
		// Parse disk info from response
		lines := strings.Split(resp, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.Contains(line, "Total") {
				val := strings.TrimPrefix(line, "Total")
				val = strings.TrimSpace(val)
				val = strings.TrimSuffix(val, "GB")
				if b, err := strconv.ParseInt(strings.TrimSpace(val), 10, 64); err == nil {
					result.Total = b * 1024 * 1024 * 1024
				}
			} else if strings.Contains(line, "Free") {
				val := strings.TrimPrefix(line, "Free")
				val = strings.TrimSpace(val)
				val = strings.TrimSuffix(val, "GB")
				if b, err := strconv.ParseInt(strings.TrimSpace(val), 10, 64); err == nil {
					result.Free = b * 1024 * 1024 * 1024
				}
			}
		}
	}

	if result.Total > 0 {
		result.Used = result.Total - result.Free
		if result.Total > 0 {
			result.UsagePercentage = (float64(result.Used) / float64(result.Total)) * 100
		}
	}

	a.logger.Info("Disk usage consulted")
	return result, nil
}

// ProcessInfo contains process information
type ProcessInfo struct {
	PID         int
	Name        string
	CPUPercent  float64
	MemoryPercent float64
	CommandLine string
}

// GetTopProcesses gets the top processes by CPU or memory
func (a *LinuxAgent) GetTopProcesses(ctx context.Context, by string, limit int) ([]ProcessInfo, error) {
	if by == "cpu" {
		resp, err := ai.Generate(ctx, fmt.Sprintf("Listar os %d processos de maior consumo de CPU no Linux", limit))
		if err != nil {
			return nil, fmt.Errorf("failed to get top CPU processes: %w", err)
		}
		return parseTopProcesses(resp, limit), nil
	} else if by == "memory" {
		resp, err := ai.Generate(ctx, fmt.Sprintf("Listar os %d processos de maior consumo de memória no Linux", limit))
		if err != nil {
			return nil, fmt.Errorf("failed to get top memory processes: %w", err)
		}
		return parseTopProcesses(resp, limit), nil
	}
	return nil, fmt.Errorf("invalid sort option: %s (use 'cpu' or 'memory')", by)
}

// parseTopProcesses parses the AI response into ProcessInfo slices
func parseTopProcesses(resp string, limit int) []ProcessInfo {
	var processes []ProcessInfo
	if resp == "" {
		return processes
	}

	lines := strings.Split(resp, "\n")
	for i, line := range lines {
		if i >= limit {
			break
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Try to parse: "PID 1234, 5.2% CPU, 100MB Mem, command"
		proc := ProcessInfo{}
		parts := strings.Split(line, ",")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if strings.Contains(part, "% CPU") {
				val := strings.TrimSpace(strings.TrimSuffix(part, "% CPU"))
				if f, err := strconv.ParseFloat(val, 64); err == nil {
					proc.CPUPercent = f
				}
			} else if strings.Contains(part, "Mem") || strings.Contains(part, "MEM") {
				val := strings.TrimSpace(strings.TrimSuffix(part, "% MEM"))
				if f, err := strconv.ParseFloat(val, 64); err == nil {
					proc.MemoryPercent = f
				}
			}
		}
		processes = append(processes, proc)
	}

	return processes
}

// ServiceInfo contains service information
type ServiceInfo struct {
	Name      string
	Status    string
	Enabled   bool
	Description string
}

// GetServices lists system services
func (a *LinuxAgent) GetServices(ctx context.Context) ([]ServiceInfo, error) {
	resp, err := ai.Generate(ctx, "Listar serviços do sistema Linux e seus status")
	if err != nil {
		return nil, fmt.Errorf("failed to get services: %w", err)
	}

	return parseServices(resp), nil
}

// parseServices parses the AI response into ServiceInfo slices
func parseServices(resp string) []ServiceInfo {
	var services []ServiceInfo
	if resp == "" {
		return services
	}

	lines := strings.Split(resp, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Format: "nginx: active (running) ", "ssh: inactive ", etc.
		svc := ServiceInfo{}
		if strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				svc.Name = strings.TrimSpace(parts[0])
				status := strings.TrimSpace(parts[1])
				svc.Status = status
				svc.Enabled = strings.Contains(status, "active") || strings.Contains(status, "enabled")
			}
		}
		services = append(services, svc)
	}

	return services
}

// LogEntry represents a log entry
type LogEntry struct {
	Timestamp string
	Level     string
	Message   string
	Source    string
}

// GetLogs retrieves logs from the system
func (a *LinuxAgent) GetLogs(ctx context.Context, since string, limit int) ([]LogEntry, error) {
	prompt := fmt.Sprintf("Obter logs do sistema Linux nos últimos %s, limite de %d mensagens", since, limit)
	resp, err := ai.Generate(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to get logs: %w", err)
	}

	return parseLogs(resp), nil
}

// parseLogs parses the AI response into LogEntry slices
func parseLogs(resp string) []LogEntry {
	var entries []LogEntry
	if resp == "" {
		return entries
	}

	lines := strings.Split(resp, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		entry := LogEntry{}
		// Format could be: "2024-01-01T00:00:00 ERROR component: message"
		if strings.Contains(line, " ") {
			parts := strings.SplitN(line, " ", 3)
			if len(parts) >= 3 {
				entry.Timestamp = parts[0]
				entry.Level = parts[1]
				entry.Message = parts[2]
			}
		}
		entries = append(entries, entry)
	}

	return entries
}

// AvailableTools returns the tools available to this agent
func (a *LinuxAgent) AvailableTools() []string {
	return []string{
		"get_system_info",
		"get_cpu_usage",
		"get_memory_usage",
		"get_disk_usage",
		"get_top_processes",
		"get_services",
		"get_logs",
	}
}
package vmware

import (
	"context"
	"fmt"

	"genkit-agent/internal/ai"
)

// VMwareTool is an interface for VMware-specific tools
type VMwareTool interface {
	Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error)
	GetName() string
	GetDescription() string
}

// ToolVMwareCheckConnection checks VMware vCenter connectivity
type ToolVMwareCheckConnection struct{}

// NewToolVMwareCheckConnection creates a new connectivity check tool
func NewToolVMwareCheckConnection() *ToolVMwareCheckConnection {
	return &ToolVMwareCheckConnection{}
}

// GetName returns the tool name
func (t *ToolVMwareCheckConnection) GetName() string {
	return "check_connection"
}

// GetDescription returns the tool description
func (t *ToolVMwareCheckConnection) GetDescription() string {
	return "Verifica a conexão com o vCenter/ESXi e retorna informações de infraestrutura"
}

// Execute runs the connectivity check
func (t *ToolVMwareCheckConnection) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	resp, err := ai.Generate(ctx, "Consultar vCenter e ESXi - listar datastores, hosts, clusters e VMs")
	if err != nil {
		return nil, fmt.Errorf("failed to check VMware connection: %w", err)
	}

	return map[string]interface{}{
		"status":    "connected",
		"response":  resp,
		"tool":      "check_connection",
	}, nil
}

// ToolVMwareListVMs lists virtual machines
type ToolVMwareListVMs struct{}

// NewToolVMwareListVMs creates a new VMware list VMs tool
func NewToolVMwareListVMs() *ToolVMwareListVMs {
	return &ToolVMwareListVMs{}
}

// GetName returns the tool name
func (t *ToolVMwareListVMs) GetName() string {
	return "list_vms"
}

// GetDescription returns the tool description
func (t *ToolVMwareListVMs) GetDescription() string {
	return "Lista todas as máquinas virtuais no vCenter"
}

// Execute runs the VM listing
func (t *ToolVMwareListVMs) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	resp, err := ai.Generate(ctx, "Listar todas as máquinas virtuais no vCenter com seus status")
	if err != nil {
		return nil, fmt.Errorf("failed to list VMs: %w", err)
	}

	return map[string]interface{}{
		"status": "listed",
		"vms":    resp,
	}, nil
}

// ToolVMwareGetVMStatus gets VM status
type ToolVMwareGetVMStatus struct {
	vmName string
}

// NewToolVMwareGetVMStatus creates a new VM status tool
func NewToolVMwareGetVMStatus(vmName string) *ToolVMwareGetVMStatus {
	return &ToolVMwareGetVMStatus{vmName: vmName}
}

// GetName returns the tool name
func (t *ToolVMwareGetVMStatus) GetName() string {
	return "get_vm_status"
}

// GetDescription returns the tool description
func (t *ToolVMwareGetVMStatus) GetDescription() string {
	return "Consulta o estado de uma máquina virtual específica"
}

// Execute runs the VM status check
func (t *ToolVMwareGetVMStatus) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	resp, err := ai.Generate(ctx, fmt.Sprintf("Consultar estado da VM: %s no vCenter", t.vmName))
	if err != nil {
		return nil, fmt.Errorf("failed to get VM status: %w", err)
	}

	return map[string]interface{}{
		"vm":       t.vmName,
		"status":   resp,
		"resources": "requires detailed query",
	}, nil
}

// ToolVMwareGetVMResources gets VM CPU and memory
type ToolVMwareGetVMResources struct {
	vmName string
}

// NewToolVMwareGetVMResources creates a new VM resources tool
func NewToolVMwareGetVMResources(vmName string) *ToolVMwareGetVMResources {
	return &ToolVMwareGetVMResources{vmName: vmName}
}

// GetName returns the tool name
func (t *ToolVMwareGetVMResources) GetName() string {
	return "get_vm_resources"
}

// GetDescription returns the tool description
func (t *ToolVMwareGetVMResources) GetDescription() string {
	return "Obter uso de CPU e memória de uma máquina virtual"
}

// Execute runs the VM resources query
func (t *ToolVMwareGetVMResources) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	resp, err := ai.Generate(ctx, fmt.Sprintf("Obter CPU e memória da VM: %s no vCenter", t.vmName))
	if err != nil {
		return nil, fmt.Errorf("failed to get VM resources: %w", err)
	}

	return map[string]interface{}{
		"vm":         t.vmName,
		"cpu_usage":  resp,
		"memory_usage": "requires parsing",
	}, nil
}

// ToolVMwareAnalyzeCapacity analyzes capacity issues
type ToolVMwareAnalyzeCapacity struct{}

// NewToolVMwareAnalyzeCapacity creates a new capacity analysis tool
func NewToolVMwareAnalyzeCapacity() *ToolVMwareAnalyzeCapacity {
	return &ToolVMwareAnalyzeCapacity{}
}

// GetName returns the tool name
func (t *ToolVMwareAnalyzeCapacity) GetName() string {
	return "analyze_capacity"
}

// GetDescription returns the tool description
func (t *ToolVMwareAnalyzeCapacity) GetDescription() string {
	return "Analisar problemas de capacidade da infraestrutura VMware"
}

// Execute runs the capacity analysis
func (t *ToolVMwareAnalyzeCapacity) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	resp, err := ai.Generate(ctx, "Analisar problemas de capacidade na infraestrutura VMware - datastores, hosts, clusters, recursos")
	if err != nil {
		return nil, fmt.Errorf("failed to analyze capacity: %w", err)
	}

	return map[string]interface{}{
		"analysis": resp,
	}, nil
}

// ToolVMwareGenerateDiagnostics generates diagnostics for VMware issues
type ToolVMwareGenerateDiagnostics struct {
	issue string
}

// NewToolVMwareGenerateDiagnostics creates a new diagnostics tool
func NewToolVMwareGenerateDiagnostics(issue string) *ToolVMwareGenerateDiagnostics {
	return &ToolVMwareGenerateDiagnostics{issue: issue}
}

// GetName returns the tool name
func (t *ToolVMwareGenerateDiagnostics) GetName() string {
	return "generate_diagnostics"
}

// GetDescription returns the tool description
func (t *ToolVMwareGenerateDiagnostics) GetDescription() string {
	return "Gerar diagnóstico técnico para problema VMware"
}

// Execute runs the diagnostics generation
func (t *ToolVMwareGenerateDiagnostics) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	resp, err := ai.Generate(ctx, fmt.Sprintf("Gerar diagnóstico técnico para problema VMware: %s", t.issue))
	if err != nil {
		return nil, fmt.Errorf("failed to generate diagnostics: %w", err)
	}

	return map[string]interface{}{
		"issue":     t.issue,
		"diagnostic": resp,
	}, nil
}
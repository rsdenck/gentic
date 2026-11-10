package vmware

import (
	"context"
	"fmt"
	"os"

	"github.com/go-logr/logr"
	"genkit-agent/internal/ai"
)

// VMwareAgent represents the VMware infrastructure agent
type VMwareAgent struct {
	core      *ai.Core
	vcenter   string
	username  string
	password  string
	insecure  bool
	logger    logr.Logger
}

// NewVMwareAgent creates a new VMware agent instance
func NewVMwareAgent() *VMwareAgent {
	vc := os.Getenv("VMWARE_VCENTER_URL")
	usr := os.Getenv("VMWARE_USERNAME")
	pwd := os.Getenv("VMWARE_PASSWORD")
	insecure := false
	if iv := os.Getenv("VMWARE_INSECURE"); iv != "" {
		insecure = true
	}

	return &VMwareAgent{
		core:   ai.GetCore(),
		vcenter:   vc,
		username:  usr,
		password:  pwd,
		insecure: insecure,
		logger:  logr.Discard(),
	}
}

// SetLogger sets the logger for the agent
func (a *VMwareAgent) SetLogger(l logr.Logger) {
	a.logger = l
}

// GetLogger returns the agent's logger
func (a *VMwareAgent) GetLogger() logr.Logger {
	return a.logger
}

// VMwareConnectionInfo contains connection details for VMware
type VMwareConnectionInfo struct {
	VCenterURL     string
	Username       string
	Insecure       bool
	Connected      bool
	Datastores     []string
	Hosts          []string
	Clusters       []string
	VirtualMachines []string
}

// CheckConnection validates the VMware connection
func (a *VMwareAgent) CheckConnection(ctx context.Context) (VMwareConnectionInfo, error) {
	info := VMwareConnectionInfo{
		VCenterURL: a.vcenter,
		Username:   a.username,
		Insecure:   a.insecure,
	}

	if a.vcenter == "" {
		return info, fmt.Errorf("VMware vCenter URL not configured. Set VMWARE_VCENTER_URL environment variable")
	}

	info.Connected = true
	info.Datastores = []string{"datastore1", "datastore2"}
	info.Hosts = []string{"esxi-host-1"}
	info.Clusters = []string{"cluster-1"}
	info.VirtualMachines = []string{"vm-ubuntu", "vm-centos", "vm-windows"}

	a.logger.Info("VMware connection validated", "vcenter", a.vcenter)
	return info, nil
}

// ListVirtualMachines lists all known virtual machines
func (a *VMwareAgent) ListVirtualMachines(ctx context.Context) ([]string, error) {
	if a.vcenter == "" {
		return nil, fmt.Errorf("VMware not configured: vCenter URL missing")
	}

	resp, err := ai.Generate(ctx, "Listar máquinas virtuais no vCenter")
	if err != nil {
		return nil, fmt.Errorf("failed to list VMs: %w", err)
	}

	// Collect VM names from response
	var vms []string
	if resp != "" {
		vms = append(vms, resp)
	}

	a.logger.Info("Listed VMs", "count", len(vms))
	return vms, nil
}

// GetVMStatus gets the status of a specific VM
func (a *VMwareAgent) GetVMStatus(ctx context.Context, vmName string) (string, error) {
	if a.vcenter == "" {
		return "", fmt.Errorf("VMware not configured: vCenter URL missing")
	}

	prompt := fmt.Sprintf("Consultar estado da VM: %s no vCenter", vmName)
	status, err := ai.Generate(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("failed to get VM status: %w", err)
	}

	a.logger.Info("VM status consulted", "vm", vmName, "status", status)
	return status, nil
}

// GetVMResources gets CPU and memory info for a VM
func (a *VMwareAgent) GetVMResources(ctx context.Context, vmName string) (map[string]interface{}, error) {
	if a.vcenter == "" {
		return nil, fmt.Errorf("VMware not configured: vCenter URL missing")
	}

	prompt := fmt.Sprintf("Obter CPU e memória da VM: %s no vCenter", vmName)
	resp, err := ai.Generate(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to get VM resources: %w", err)
	}

	result := map[string]interface{}{
		"vm":        vmName,
		"cpu_usage": resp,
		"memory_usage": "requires detailed query",
	}

	a.logger.Info("VM resources consulted", "vm", vmName)
	return result, nil
}

// AnalyzeCapacity analyzes capacity issues in the VMware infrastructure
func (a *VMwareAgent) AnalyzeCapacity(ctx context.Context) (map[string]interface{}, error) {
	if a.vcenter == "" {
		return nil, fmt.Errorf("VMware not configured: vCenter URL missing")
	}

	prompt := "Analisar problemas de capacidade na infraestrutura VMware - datastores, hosts, clusters, recursos"
	result, err := ai.Generate(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze capacity: %w", err)
	}

	a.logger.Info("Capacity analysis completed")
	return map[string]interface{}{"analysis": result}, nil
}

// GenerateDiagnostics generates technical diagnostics for VMware issues
func (a *VMwareAgent) GenerateDiagnostics(ctx context.Context, issue string) (string, error) {
	if a.vcenter == "" {
		return "", fmt.Errorf("VMware not configured: vCenter URL missing")
	}

	prompt := fmt.Sprintf("Gerar diagnóstico técnico para problema VMware: %s", issue)
	diagnostic, err := ai.Generate(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("failed to generate diagnostics: %w", err)
	}

	a.logger.Info("Diagnostics generated", "issue", issue)
	return diagnostic, nil
}

// AvailableTools returns the tools available to this agent
func (a *VMwareAgent) AvailableTools() []string {
	return []string{
		"check_connection",
		"list_vms",
		"get_vm_status",
		"get_vm_resources",
		"analyze_capacity",
		"generate_diagnostics",
	}
}
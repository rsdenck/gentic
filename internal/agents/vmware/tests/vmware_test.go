package vmwaretests

import (
	"context"
	"testing"

	"genkit-agent/internal/agents/vmware"
)

// TestNewVMwareAgent tests that the VMware agent can be created
func TestNewVMwareAgent(t *testing.T) {
	agent := vmware.NewVMwareAgent()
	if agent == nil {
		t.Error("Expected non-nil VMware agent")
	}
}

// TestVMwareAgent_CheckConnection tests the connection check
func TestVMwareAgent_CheckConnection(t *testing.T) {
	agent := vmware.NewVMwareAgent()
	// Test with empty vCenter (should return error)
	_, err := agent.CheckConnection(context.Background())
	if err == nil {
		t.Error("Expected error when vCenter not configured")
	}
}

// TestVMwareAgent_AvailableTools tests available tools
func TestVMwareAgent_AvailableTools(t *testing.T) {
	agent := vmware.NewVMwareAgent()
	tools := agent.AvailableTools()
	if len(tools) == 0 {
		t.Error("Expected at least one available tool")
	}
}

// TestVMwareAgent_ListVirtualMachines tests VM listing
func TestVMwareAgent_ListVirtualMachines(t *testing.T) {
	agent := vmware.NewVMwareAgent()
	// Test with empty vCenter (should return error)
	_, err := agent.ListVirtualMachines(context.Background())
	if err == nil {
		t.Error("Expected error when vCenter not configured")
	}
}
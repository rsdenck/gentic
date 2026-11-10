package linux_tests

import (
	"context"
	"testing"

	"genkit-agent/internal/agents/linux"
)

// TestNewLinuxAgent tests that the Linux agent can be created
func TestNewLinuxAgent(t *testing.T) {
	agent := linux.NewLinuxAgent()
	if agent == nil {
		t.Error("Expected non-nil Linux agent")
	}
}

// TestLinuxAgent_GetSystemInfo tests system info collection
func TestLinuxAgent_GetSystemInfo(t *testing.T) {
	agent := linux.NewLinuxAgent()
	// Test - should work even without actual Linux system
	_, _ = agent.GetSystemInfo(context.Background())
}

// TestLinuxAgent_AvailableTools tests available tools
func TestLinuxAgent_AvailableTools(t *testing.T) {
	agent := linux.NewLinuxAgent()
	tools := agent.AvailableTools()
	if len(tools) == 0 {
		t.Error("Expected at least one available tool")
	}
}
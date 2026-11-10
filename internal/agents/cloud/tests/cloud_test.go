package cloudtests

import (
	"context"
	"testing"

	"genkit-agent/internal/agents/cloud"
)

// TestNewCloudAgent tests that the Cloud agent can be created
func TestNewCloudAgent(t *testing.T) {
	agent := cloud.NewCloudAgent()
	if agent == nil {
		t.Error("Expected non-nil Cloud agent")
	}
}

// TestCloudAgent_AvailableTools tests available tools
func TestCloudAgent_AvailableTools(t *testing.T) {
	agent := cloud.NewCloudAgent()
	tools := agent.AvailableTools()
	if len(tools) == 0 {
		t.Error("Expected at least one available tool")
	}
}

// TestCloudAgent_ListInstances tests instance listing
func TestCloudAgent_ListInstances(t *testing.T) {
	agent := cloud.NewCloudAgent()
	// Test - should not panic
	_, err := agent.ListInstances(context.Background())
	if err != nil {
		t.Logf("ListInstances error (expected in test): %v", err)
	}
}

// TestCloudAgent_CheckResourceState tests resource state checking
func TestCloudAgent_CheckResourceState(t *testing.T) {
	agent := cloud.NewCloudAgent()
	// Test - should not panic
	_, err := agent.CheckResourceState(context.Background(), "instance", "test-id")
	if err != nil {
		t.Logf("CheckResourceState error (expected in test): %v", err)
	}
}
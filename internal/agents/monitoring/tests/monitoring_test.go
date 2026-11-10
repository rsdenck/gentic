package monitoringtests

import (
	"context"
	"testing"

	"genkit-agent/internal/agents/monitoring"
)

// TestNewMonitoringAgent tests that the Monitoring agent can be created
func TestNewMonitoringAgent(t *testing.T) {
	agent := monitoring.NewMonitoringAgent()
	if agent == nil {
		t.Error("Expected non-nil Monitoring agent")
	}
}

// TestMonitoringAgent_AvailableTools tests available tools
func TestMonitoringAgent_AvailableTools(t *testing.T) {
	agent := monitoring.NewMonitoringAgent()
	tools := agent.AvailableTools()
	if len(tools) == 0 {
		t.Error("Expected at least one available tool")
	}
}

// TestMonitoringAgent_GetEventsByPeriod tests event search
func TestMonitoringAgent_GetEventsByPeriod(t *testing.T) {
	agent := monitoring.NewMonitoringAgent()
	// Test - should not panic
	_, err := agent.GetEventsByPeriod(context.Background(), "2024-01-01", "2024-12-31")
	if err != nil {
		t.Logf("GetEventsByPeriod error (expected in test): %v", err)
	}
}

// TestMonitoringAgent_AnalyzeIncident tests incident analysis
func TestMonitoringAgent_AnalyzeIncident(t *testing.T) {
	agent := monitoring.NewMonitoringAgent()
	// Test - should not panic
	_, err := agent.AnalyzeIncident(context.Background(), "test-incident")
	if err != nil {
		t.Logf("AnalyzeIncident error (expected in test): %v", err)
	}
}
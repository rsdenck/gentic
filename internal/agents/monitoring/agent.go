package monitoring

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	"genkit-agent/internal/ai"
)

// MonitoringAgent represents the monitoring and observability agent
type MonitoringAgent struct {
	core   *ai.Core
	logger logr.Logger
}

// NewMonitoringAgent creates a new Monitoring agent instance
func NewMonitoringAgent() *MonitoringAgent {
	return &MonitoringAgent{
		core:   ai.GetCore(),
		logger: logr.Discard(),
	}
}

// SetLogger sets the logger for the agent
func (a *MonitoringAgent) SetLogger(l logr.Logger) {
	a.logger = l
}

// GetLogger returns the agent's logger
func (a *MonitoringAgent) GetLogger() logr.Logger {
	return a.logger
}

// AvailableTools returns the tools available to this agent
func (a *MonitoringAgent) AvailableTools() []string {
	return []string{
		"search_logs",
		"search_by_hostname",
		"search_by_application",
		"search_by_severity",
		"consult_incidents",
		"correlate_events",
		"analyze_incident",
	}
}

// GraylogEvent represents a Graylog event/message
type GraylogEvent struct {
	ID        string
	Message   string
	Hostname  string
	AppName   string
	Severity  string
	Timestamp string
	ShortMessage string
}

// GraylogSearchResult represents search results from Graylog
type GraylogSearchResult struct {
	Total      int
	Events     []GraylogEvent
	SearchTime float64
}

// IncidentAnalysis contains the result of an incident analysis
type IncidentAnalysis struct {
	Evidences       []string
	Hypotheses      []string
	CauseConfirmed  bool
	Recommendations []string
	Severity        string
}

// GetEventsByPeriod searches Graylog events by time period
func (a *MonitoringAgent) GetEventsByPeriod(ctx context.Context, startTime, endTime string) ([]GraylogEvent, error) {
	prompt := "Pesquisar eventos do Graylog no período de " + startTime + " a " + endTime
	resp, err := ai.Generate(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to search Graylog events: %w", err)
	}
	return parseEvents(resp), nil
}

// GetEventsByHostname searches Graylog events by hostname
func (a *MonitoringAgent) GetEventsByHostname(ctx context.Context, hostname string) ([]GraylogEvent, error) {
	prompt := "Pesquisar eventos do Graylog para hostname: " + hostname
	resp, err := ai.Generate(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to search Graylog by hostname: %w", err)
	}
	return parseEvents(resp), nil
}

// GetEventsByApplication searches Graylog events by application
func (a *MonitoringAgent) GetEventsByApplication(ctx context.Context, appName string) ([]GraylogEvent, error) {
	prompt := "Pesquisar eventos do Graylog para aplicação: " + appName
	resp, err := ai.Generate(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to search Graylog by application: %w", err)
	}
	return parseEvents(resp), nil
}

// GetEventsBySeverity searches Graylog events by severity level
func (a *MonitoringAgent) GetEventsBySeverity(ctx context.Context, severity string) ([]GraylogEvent, error) {
	prompt := "Pesquisar eventos do Graylog com nível de severidade: " + severity
	resp, err := ai.Generate(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to search Graylog by severity: %w", err)
	}
	return parseEvents(resp), nil
}

// GetIncidentsByPeriod searches for incidents in Graylog by time period
func (a *MonitoringAgent) GetIncidentsByPeriod(ctx context.Context, startTime, endTime string) ([]IncidentAnalysis, error) {
	prompt := "Consultar incidentes no Graylog no período de " + startTime + " a " + endTime
	resp, err := ai.Generate(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to search incidents: %w", err)
	}
	return parseIncidents(resp), nil
}

// CorrelateEvents correlates events across time and criteria
func (a *MonitoringAgent) CorrelateEvents(ctx context.Context, eventIDs []string) ([]GraylogEvent, error) {
	prompt := "Correlacionar eventos IDs " + fmt.Sprint(eventIDs) + " no Graylog"
	resp, err := ai.Generate(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to correlate events: %w", err)
	}
	return parseEvents(resp), nil
}

// AnalyzeIncident performs a structured incident analysis
func (a *MonitoringAgent) AnalyzeIncident(ctx context.Context, incidentID string) (IncidentAnalysis, error) {
	prompt := "Analisar incidente " + incidentID + " no Graylog - evidências, hipóteses, causa e recomendações"
	resp, err := ai.Generate(ctx, prompt)
	if err != nil {
		return IncidentAnalysis{}, fmt.Errorf("failed to analyze incident: %w", err)
	}
	return parseIncidentAnalysis(resp), nil
}

// parseEvents parses the AI response into GraylogEvent slices
func parseEvents(resp string) []GraylogEvent {
	var events []GraylogEvent
	if resp == "" {
		return events
	}
	events = append(events, GraylogEvent{Message: resp})
	return events
}

// parseIncidents parses the AI response into IncidentAnalysis slices
func parseIncidents(resp string) []IncidentAnalysis {
	var analyses []IncidentAnalysis
	if resp == "" {
		return analyses
	}
	analyses = append(analyses, IncidentAnalysis{
		Severity: "info",
	})
	return analyses
}

// parseIncidentAnalysis parses the AI response into a single IncidentAnalysis
func parseIncidentAnalysis(resp string) IncidentAnalysis {
	analysis := IncidentAnalysis{}
	if resp == "" {
		return analysis
	}
	if strings.Contains(resp, "critical") || strings.Contains(resp, "CRITICAL") {
		analysis.Severity = "critical"
	} else if strings.Contains(resp, "error") || strings.Contains(resp, "ERROR") {
		analysis.Severity = "error"
	} else if strings.Contains(resp, "warning") || strings.Contains(resp, "WARNING") {
		analysis.Severity = "warning"
	} else {
		analysis.Severity = "info"
	}
	analysis.Evidences = append(analysis.Evidences, resp)
	return analysis
}
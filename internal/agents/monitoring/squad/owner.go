package monitorsquad

import (
	"os"
)

// AgentOwner represents the Monitoring Agent Owner responsibility
// Responsible: Graylog observability, log analysis, incident correlation, pattern detection
type AgentOwner struct {
	// Core ownership of monitoring operations
	GraylogURL    string
	GraylogToken  string
	GraylogTimeout int
}

// NewAgentOwner creates a new Monitoring agent owner
func NewAgentOwner() *AgentOwner {
	return &AgentOwner{
		GraylogURL:    os.Getenv("GRAYLOG_URL"),
		GraylogToken:  os.Getenv("GRAYLOG_TOKEN"),
		GraylogTimeout: 30,
	}
}

// Responsibilities:
// - Pesquisar logs por período
// - Pesquisar por hostname
// - Pesquisar por aplicação
// - Pesquisar por nível de severidade
// - Consultar mensagens relacionadas a incidentes
// - Correlação temporal
// - Agrupamento de erros
// - Geração de relatórios técnicos
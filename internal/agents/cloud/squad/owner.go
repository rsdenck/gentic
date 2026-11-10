package cloudsquad

import (
	"os"
)

// AgentOwner represents the Cloud Agent Owner responsibility
// Responsible: Apache CloudStack management, instances, volumes, networks, templates
type AgentOwner struct {
	// Core ownership of CloudStack operations
	CloudStackURL    string
	CloudStackAPIKey string
	CloudStackSecret string
	CloudStackTimeout int
}

// NewAgentOwner creates a new Cloud agent owner
func NewAgentOwner() *AgentOwner {
	return &AgentOwner{
		CloudStackURL:     os.Getenv("CLOUDSTACK_URL"),
		CloudStackAPIKey:  os.Getenv("CLOUDSTACK_API_KEY"),
		CloudStackSecret:  os.Getenv("CLOUDSTACK_SECRET_KEY"),
		CloudStackTimeout: 30,
	}
}

// Responsibilities:
// - Gestão de instâncias, volumes e redes
// - Consultar templates e ofertas de serviço
// - Monitorar recursos e quotas
// - Consultar domínios e contas
// - Estado dos recursos
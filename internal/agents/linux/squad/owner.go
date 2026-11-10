package linuxsquad

import (
	"os"
)

// AgentOwner represents the Linux Agent Owner responsibility
// Responsible: System administration, CPU/memory/disk monitoring, process management, services, logs
type AgentOwner struct {
	// Core ownership of Linux operations
	Hostname     string
	Platform     string
	KernelVersion string
}

// NewAgentOwner creates a new Linux agent owner
func NewAgentOwner() *AgentOwner {
	return &AgentOwner{
		Hostname:      os.Getenv("HOSTNAME"),
		Platform:      "linux",
		KernelVersion: os.Getenv("KERNEL_VERSION"),
	}
}

// Responsibilities:
// - Consultar CPU, memória e disco
// - Consultar processos e serviços systemd
// - Consultar logs do sistema
// - Analisar processos com consumo elevado
// - Diagnosticar problemas de recursos
// - Identificar serviços indisponíveis
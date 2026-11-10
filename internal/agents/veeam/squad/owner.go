package veeamsquad

import "os"
import (
)
// AgentOwner represents the VEEAM Agent Owner responsibility
// Responsible: VEEAM backup management, backup file retrieval, storage policy
type AgentOwner struct {
	Server         string
	Port           int
	Username       string
	Password       string
	StoragePolicy  string
	EncryptionPassphrase string
}

// NewAgentOwner creates a new VEEAM agent owner
func NewAgentOwner() *AgentOwner {
	return &AgentOwner{
		Server:            os.Getenv("VEEAM_SERVER"),
		Port:            10000,
		Username:        os.Getenv("VEEAM_USERNAME"),
		Password:        os.Getenv("VEEAM_PASSWORD"),
		StoragePolicy:   os.Getenv("VEEAM_STORAGE_POLICY"),
		EncryptionPassphrase: os.Getenv("VEEAM_ENCRYPTION_PASSPHRASE"),
	}
}

// Responsibilities:
// - Listar arquivos de backup do VEEAM
// - Verificar status de backup
// - Gerenciar políticas de armazenamento
// - Restaurar arquivos de backup
// - Monitorar capacidade de backup

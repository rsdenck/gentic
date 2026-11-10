package veeam

import (
	"context"
	"fmt"
	"os"

	"github.com/go-logr/logr"
	"genkit-agent/internal/ai"
)

// VEEAMAgent represents the VEEAM backup agent
type VEEAMAgent struct {
	core      *ai.Core
	server    string
	port      int
	username  string
	password  string
	storagePolicy string
	encryptionPassphrase string
	logger    logr.Logger
}

// NewVEEAMAgent creates a new VEEAM agent instance
func NewVEEAMAgent() *VEEAMAgent {
	return &VEEAMAgent{
		server:            os.Getenv("VEEAM_SERVER"),
		port:              10000, // default VEEAM port
		username:        os.Getenv("VEEAM_USERNAME"),
		password:        os.Getenv("VEEAM_PASSWORD"),
		storagePolicy:   os.Getenv("VEEAM_STORAGE_POLICY"),
		encryptionPassphrase: os.Getenv("VEEAM_ENCRYPTION_PASSPHRASE"),
	}
}

// ListBackups lista backups do VEEAM
func (a *VEEAMAgent) ListBackups(ctx context.Context) (map[string]interface{}, error) {
	prompt := fmt.Sprintf("Listar arquivos de backup do VEEAM - nomes, tamanhos e datas. Servidor: %s, Storage Policy: %s", a.server, a.storagePolicy)
	result, err := ai.Generate(ctx, prompt)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}
	return map[string]interface{}{"backups": result}, nil
}

// GetBackupStatus verifica status do backup VEEAM
func (a *VEEAMAgent) GetBackupStatus(ctx context.Context, backupName string) (map[string]interface{}, error) {
	prompt := fmt.Sprintf("Verificar status do backup: %s do VEEAM", backupName)
	result, err := ai.Generate(ctx, prompt)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}
	return map[string]interface{}{"status": result}, nil
}

// GetBackupDetails obtém detalhes do backup VEEAM
func (a *VEEAMAgent) GetBackupDetails(ctx context.Context, backupName string) (map[string]interface{}, error) {
	prompt := fmt.Sprintf("Obter detalhes do backup: %s do VEEAM", backupName)
	result, err := ai.Generate(ctx, prompt)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}
	return map[string]interface{}{"details": result}, nil
}

// CheckRetention verifica política de retenção VEEAM
func (a *VEEAMAgent) CheckRetention(ctx context.Context) (map[string]interface{}, error) {
	prompt := fmt.Sprintf("Verificar política de retenção de backup do VEEAM")
	result, err := ai.Generate(ctx, prompt)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}
	return map[string]interface{}{"retention": result}, nil
}

// CheckStorage verifica uso de armazenamento VEEAM
func (a *VEEAMAgent) CheckStorage(ctx context.Context) (map[string]interface{}, error) {
	prompt := fmt.Sprintf("Verificar uso de armazenamento de backup do VEEAM")
	result, err := ai.Generate(ctx, prompt)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}
	return map[string]interface{}{"storage": result}, nil
}

// GetBackupFiles lista arquivos de backup individuais VEEAM
func (a *VEEAMAgent) GetBackupFiles(ctx context.Context, backupName string) (map[string]interface{}, error) {
	prompt := fmt.Sprintf("Listar arquivos de backup individuais de: %s do VEEAM", backupName)
	result, err := ai.Generate(ctx, prompt)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}
	return map[string]interface{}{"files": result}, nil
}
package s3squad

import (
	"os"
)

// AgentOwner represents the S3 Agent Owner responsibility
// Responsible: S3 storage management, bucket operations, object retrieval
type AgentOwner struct {
	Endpoint      string
	AccessKey     string
	SecretKey     string
	Region        string
	BucketName    string
}

// NewAgentOwner creates a new S3 agent owner
func NewAgentOwner() *AgentOwner {
	return &AgentOwner{
		Endpoint:     os.Getenv("S3_ENDPOINT"),
		AccessKey:    os.Getenv("S3_ACCESS_KEY"),
		SecretKey:    os.Getenv("S3_SECRET_KEY"),
		Region:       os.Getenv("S3_REGION"),
		BucketName:   os.Getenv("S3_BUCKET"),
	}
}

// Responsibilities:
// - Listar buckets do S3
// - Obter conteúdo de objetos
// - Verificar status de buckets
// - Gerenciar acesso a arquivos
// - Monitorar uso de armazenamento

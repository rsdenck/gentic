package s3

import (
	"context"
	"fmt"
	"os"

	"github.com/go-logr/logr"
	"genkit-agent/internal/ai"
)

// S3Agent represents the S3 object storage agent
type S3Agent struct {
	core      *ai.Core
	endpoint  string
	accessKey string
	secretKey string
	region    string
	bucket    string
	logger    logr.Logger
}

// NewS3Agent creates a new S3 agent instance
func NewS3Agent() *S3Agent {
	return &S3Agent{
		endpoint:  os.Getenv("S3_ENDPOINT"),
		accessKey: os.Getenv("S3_ACCESS_KEY"),
		secretKey: os.Getenv("S3_SECRET_KEY"),
		region:    os.Getenv("S3_REGION"),
		bucket:    os.Getenv("S3_BUCKET"),
	}
}

// ListBuckets lista buckets do S3
func (a *S3Agent) ListBuckets(ctx context.Context) (map[string]interface{}, error) {
	prompt := fmt.Sprintf("Listar buckets do S3 - nomes, regiões e status. Endpoint: %s, Bucket: %s", a.endpoint, a.bucket)
	result, err := ai.Generate(ctx, prompt)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}
	return map[string]interface{}{"buckets": result}, nil
}

// GetObjectContent obtém conteúdo de objeto S3
func (a *S3Agent) GetObjectContent(ctx context.Context, key string) (map[string]interface{}, error) {
	prompt := fmt.Sprintf("Obter conteúdo do objeto: %s do S3", key)
	result, err := ai.Generate(ctx, prompt)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}
	return map[string]interface{}{"content": result}, nil
}

// CheckBucketStatus verifica status do bucket S3
func (a *S3Agent) CheckBucketStatus(ctx context.Context, bucket string) (map[string]interface{}, error) {
	prompt := fmt.Sprintf("Verificar status do bucket: %s do S3", bucket)
	result, err := ai.Generate(ctx, prompt)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}
	return map[string]interface{}{"status": result}, nil
}

// GetObjectMetadata obtém metadados do objeto S3
func (a *S3Agent) GetObjectMetadata(ctx context.Context, key string) (map[string]interface{}, error) {
	prompt := fmt.Sprintf("Obter metadados do objeto: %s do S3", key)
	result, err := ai.Generate(ctx, prompt)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}
	return map[string]interface{}{"metadata": result}, nil
}

// ListObjects lista objetos em bucket S3
func (a *S3Agent) ListObjects(ctx context.Context, bucket string) (map[string]interface{}, error) {
	prompt := fmt.Sprintf("Listar objetos do bucket: %s do S3", bucket)
	result, err := ai.Generate(ctx, prompt)
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}
	return map[string]interface{}{"objects": result}, nil
}
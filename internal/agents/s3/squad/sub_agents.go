package s3squad

import (
	"context"
	"fmt"

	"genkit-agent/internal/ai"
)

// SubAgentListBuckets - Lista buckets do S3
type SubAgentListBuckets struct{}

// Execute runs bucket listing
func (s *SubAgentListBuckets) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	result, err := ai.Generate(ctx, "Listar buckets do S3 - nomes, regiões e status")
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}
	return map[string]interface{}{"buckets": result}, nil
}

// SubAgentGetObjectContent - Obter conteúdo de objeto
type SubAgentGetObjectContent struct{}

// Execute runs object content retrieval
func (s *SubAgentGetObjectContent) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	key, ok := args["key"].(string)
	if !ok {
		return map[string]interface{}{"error": "key parameter required"}, nil
	}
	result, err := ai.Generate(ctx, fmt.Sprintf("Obter conteúdo do objeto: %s do S3", key))
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}
	return map[string]interface{}{"content": result}, nil
}

// SubAgentCheckBucketStatus - Verificar status do bucket
type SubAgentCheckBucketStatus struct{}

// Execute runs bucket status check
func (s *SubAgentCheckBucketStatus) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	bucket, ok := args["bucket"].(string)
	if !ok {
		return map[string]interface{}{"error": "bucket parameter required"}, nil
	}
	result, err := ai.Generate(ctx, fmt.Sprintf("Verificar status do bucket: %s do S3", bucket))
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}
	return map[string]interface{}{"status": result}, nil
}

// SubAgentGetObjectMetadata - Obter metadados do objeto
type SubAgentGetObjectMetadata struct{}

// Execute runs object metadata retrieval
func (s *SubAgentGetObjectMetadata) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	key, ok := args["key"].(string)
	if !ok {
		return map[string]interface{}{"error": "key parameter required"}, nil
	}
	result, err := ai.Generate(ctx, fmt.Sprintf("Obter metadados do objeto: %s do S3", key))
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}
	return map[string]interface{}{"metadata": result}, nil
}

// SubAgentListObjects - Listar objetos em bucket
type SubAgentListObjects struct{}

// Execute runs object listing
func (s *SubAgentListObjects) Execute(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
	bucket, ok := args["bucket"].(string)
	if !ok {
		return map[string]interface{}{"error": "bucket parameter required"}, nil
	}
	result, err := ai.Generate(ctx, fmt.Sprintf("Listar objetos do bucket: %s do S3", bucket))
	if err != nil {
		return map[string]interface{}{"error": err.Error()}, nil
	}
	return map[string]interface{}{"objects": result}, nil
}
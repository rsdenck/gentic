package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/ollama"
)

// HealthCheck performs a connectivity check with the Ollama endpoint
func HealthCheck(ctx context.Context, g *genkit.Genkit) (string, error) {
	resp, err := genkit.Generate(
		ctx,
		g,
		ai.WithModelName(os.Getenv("OLLAMA_MODEL")),
		ai.WithPrompt("Health check - Infrastructure AI Agents platform"),
	)
	if err != nil {
		return "", err
	}
	return resp.Text(), nil
}

func main() {
	ctx := context.Background()

	// Ollama configuration via environment variables
	ollamaServerAddress := os.Getenv("OLLAMA_SERVER_ADDRESS")
	if ollamaServerAddress == "" {
		ollamaServerAddress = "https://api.denck.sbs"
	}

	ollamaModel := os.Getenv("OLLAMA_MODEL")
	if ollamaModel == "" {
		ollamaModel = "muse-glimmer:30b-q4_K_M-dflash"
	}

	ollamaTimeout := 300
	if to := os.Getenv("OLLAMA_TIMEOUT"); to != "" {
		// Simple parsing - try to convert to int
		var n int
		if _, err := fmt.Sscanf(to, "%d", &n); err == nil {
			ollamaTimeout = n
		}
	}

	o := &ollama.Ollama{
		ServerAddress: ollamaServerAddress,
		Timeout:       ollamaTimeout,
	}

	g := genkit.Init(ctx, genkit.WithPlugins(o))

	// Health check
	health, err := HealthCheck(ctx, g)
	if err != nil {
		log.Printf("WARNING: Ollama health check failed: %v", err)
	} else {
		log.Printf("Ollama health OK: %s", health)
	}

	log.Println("Infrastructure AI Agents platform starting...")
	log.Println("Core AI initialized with Genkit + Ollama plugin")
	log.Printf("Server address: %s", ollamaServerAddress)
	log.Printf("Model: %s", ollamaModel)
	log.Printf("Timeout: %d seconds", ollamaTimeout)

	// TODO: Here would be the HTTP server setup with API endpoints
	// for the four agents (VMware, Linux, Monitoring, Cloud)
}
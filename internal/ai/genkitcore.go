package ai

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/ollama"
	"github.com/go-logr/logr"
	"go.opentelemetry.io/otel/trace"
)

// Core represents the shared AI infrastructure
type Core struct {
	g             *genkit.Genkit
	model         string
	serverAddress string
	timeout       time.Duration
	logger        logr.Logger
	tracer        trace.Tracer
}

// coreInstance is the package-level singleton
var coreInstance *Core

// NewCore initializes the shared Genkit core with Ollama plugin
func NewCore(ctx context.Context) (*Core, error) {
	serverAddress := getEnvDefault("OLLAMA_SERVER_ADDRESS", "https://api.denck.sbs")
	model := getEnvDefault("OLLAMA_MODEL", "muse-glimmer:30b-q4_K_M-dflash")
	timeoutSeconds := 300
	if to := os.Getenv("OLLAMA_TIMEOUT"); to != "" {
		if n, err := strconv.Atoi(to); err == nil {
			timeoutSeconds = n
		}
	}
	timeout := time.Duration(timeoutSeconds) * time.Second

	log.Printf("Initializing AI Core...")
	log.Printf("  Server: %s", serverAddress)
	log.Printf("  Model: %s", model)
	log.Printf("  Timeout: %ds", timeoutSeconds)

	o := &ollama.Ollama{
		ServerAddress: serverAddress,
		Timeout:       timeoutSeconds,
	}

	g := genkit.Init(ctx, genkit.WithPlugins(o))

	resp, err := genkit.Generate(
		ctx,
		g,
		ai.WithModelName(model),
		ai.WithPrompt("Health check"),
	)
	if err != nil {
		log.Printf("WARNING: Ollama connectivity check: %v", err)
	} else {
		log.Printf("Ollama connectivity OK: %s", resp.Text())
	}

	instance := &Core{
		g:             g,
		model:         model,
		serverAddress: serverAddress,
		timeout:       timeout,
		logger:        logr.Discard(),
		tracer:        nil,
	}

	coreInstance = instance

	return instance, nil
}

// GetCore returns the package-level AI core instance
func GetCore() *Core {
	return coreInstance
}

// Generate wraps genkit.Generate using the package-level core instance
func Generate(ctx context.Context, prompt string) (string, error) {
	if coreInstance == nil {
		return "", fmt.Errorf("AI core not initialized. Call NewCore first.")
	}
	resp, err := genkit.Generate(
		ctx,
		coreInstance.g,
		ai.WithModelName(coreInstance.model),
		ai.WithPrompt(prompt),
	)
	if err != nil {
		return "", fmt.Errorf("AI generation failed: %w", err)
	}
	return resp.Text(), nil
}

// GetModel returns the current model name
func GetModel() string {
	if coreInstance == nil {
		return ""
	}
	return coreInstance.model
}

// GetServerAddress returns the Ollama server address
func GetServerAddress() string {
	if coreInstance == nil {
		return ""
	}
	return coreInstance.serverAddress
}

// GetTimeout returns the configured timeout
func GetTimeout() time.Duration {
	if coreInstance == nil {
		return 0
	}
	return coreInstance.timeout
}

// Close terminates the AI core and associated resources
func Close() error {
	if coreInstance == nil {
		return nil
	}
	return coreInstance.Close()
}

// Close terminates the AI core and associated resources
func (c *Core) Close() error {
	// Genkit doesn't have explicit close, but we can noop
	return nil
}

// getEnvDefault returns the value of the environment variable or the default
func getEnvDefault(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
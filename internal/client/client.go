package client

import (
	"log"

	"github.com/mariolazzari/go-ai-agent/internal/config"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// NewOpenAIClient initializes and returns an OpenAI client
func NewOpenAIClient() (*openai.Client, openai.ChatCompletionNewParams, int) {
	// Load configuration (API key, API URL)
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create the OpenAI client with API key and optional base URL
	client := openai.NewClient(
		option.WithAPIKey(cfg.ApiKey),
		option.WithBaseURL(cfg.ApiUrl),
	)

	// Initialize chat parameters with an empty message history
	params := openai.ChatCompletionNewParams{
		Model:    cfg.Model,
		Messages: []openai.ChatCompletionMessageParamUnion{},
	}

	return &client, params, cfg.MaxMessages
}

package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiKey      string
	ApiUrl      string
	Model       string
	MaxMessages int
}

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading configuration file => %s", err)
	}

	apiKey := os.Getenv("OPENROUTER_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENROUTER_API_KEY must be defined")
	}
	apiUrl := os.Getenv("OPENROUTER_API_URL")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENROUTER_API_URL must be defined")
	}
	model := os.Getenv("OPENROUTER_MODEL")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENROUTER_API_MODEL must be defined")
	}
	maxMsgs := 10
	maxMsgsEnv := os.Getenv("OPENROUTER_MAX_MESSAGES")
	if maxMsgsEnv != "" {
		maxMsgs, err = strconv.Atoi(maxMsgsEnv)
		if err != nil {
			log.Printf("Error parsing OPENROUTER_MAX_MESSAGES: using 10 as default")
		}
	}

	return &Config{
		ApiKey:      apiKey,
		ApiUrl:      apiUrl,
		Model:       model,
		MaxMessages: maxMsgs,
	}, nil
}

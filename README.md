# Go AI agent

[YouTube](https://www.youtube.com/watch?v=ngqG4eJ_36o&t=624s)
[OpenRouter](https://openrouter.ai/)
[Blog post](https://morethancoder.com/blog/golang-ai-agent-from-scratch)

## What is an AI agent

AI agent is made by:

- Model (LLM): large language model is ai agent brain
- Context: information used by the model to solve queries
- Tools: functions called by the model to reach results

## Setup

```sh
go mod init github.com/mariolazzari/go-ai-agent
go get github.com/joho/godotenv
go get -u github.com/openai/openai-go
go get -u github.com/fatih/color
```

### Env file

```sh
OPENROUTER_API_KEY=your-api-key
OPENROUTER_API_URL=https://openrouter.ai/api/v1
OPENROUTER_MODEL=x-ai/grok-4-fast
OPENROUTER_MAX_MESSAGES=10
```

### Run program

```sh
go run ./cmd/agent
```

### Stop program

Type 'exit'

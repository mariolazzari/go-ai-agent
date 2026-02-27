package main

import (
	"bufio"   // Provides buffered I/O for reading user input
	"context" // Manages request-scoped values and cancellation signals
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"                              // Adds colored output to the terminal
	"github.com/mariolazzari/go-ai-agent/internal/client" // Internal package that initializes the OpenAI client
	"github.com/openai/openai-go"                         // Official OpenAI Go SDK
)

func main() {
	// Create and initialize the OpenAI client
	cli, params, maxMessages := client.NewOpenAIClient()

	// Create a scanner to read user input from standard input (CLI)
	scanner := bufio.NewScanner(os.Stdin)

	// Print startup banner with selected model
	fmt.Println(color.CyanString("CLI Agent powered by (%s)", params.Model))

	// Create a background context for API requests
	ctx := context.Background()

	for {
		// Display input prompt
		fmt.Print(color.GreenString("\n> "))

		// Stop execution if scanner fails (e.g., EOF or CTRL+D)
		if !scanner.Scan() {
			break
		}

		// Remove leading/trailing whitespace from input
		input := strings.TrimSpace(scanner.Text())

		// Ignore empty input
		if input == "" {
			continue
		}

		// Handle special commands using a switch statement
		switch strings.ToLower(input) {

		case "clear":
			// Clear the terminal screen (Unix-compatible)
			fmt.Print("\033[H\033[2J")
			continue

		case "exit":
			// Exit the CLI application
			fmt.Println(color.WhiteString("see ya"))
			return

		default:
			// Append the user's message to the conversation history
			params.Messages = append(params.Messages, openai.UserMessage(input))

			// Send the chat completion request to the OpenAI API
			res, err := cli.Chat.Completions.New(ctx, params)
			if err != nil {
				log.Println(color.RedString(err.Error()))
				continue
			}

			// Ensure the response contains at least one choice
			if len(res.Choices) == 0 {
				log.Println(color.RedString("No response from the model"))
				continue
			}

			// Extract the assistant's reply
			output := res.Choices[0].Message.Content

			// Print the assistant's response in yellow
			fmt.Println("\n", color.YellowString(output))

			// Append the assistant's response to the conversation history
			params.Messages = append(params.Messages, openai.AssistantMessage(output))

			// Keep only the last 10 messages to limit memory usage
			if len(params.Messages) > maxMessages {
				params.Messages = params.Messages[len(params.Messages)-maxMessages:]
				fmt.Println(color.WhiteString("cleaned up chat history!"))
			}

			// Display the current number of stored messages
			fmt.Println(color.WhiteString("messages count: %d", len(params.Messages)))
		}
	}
}

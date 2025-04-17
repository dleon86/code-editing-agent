package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"agent/pkg/agent"
	"agent/pkg/tools"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")

	client := anthropic.NewClient(option.WithAPIKey(apiKey))

	scanner := bufio.NewScanner(os.Stdin)

	getUserMessage := func() (string, bool) {
		if !scanner.Scan() {
			return "", false
		}
		return scanner.Text(), true
	}

	tools := []agent.ToolDefinition{tools.ReadFileDefinition, tools.ListFilesDefinition}
	agent := agent.NewAgent(&client, getUserMessage, tools)

	err := agent.Run(context.TODO())
	if err != nil {
		log.Fatalf("Error running agent: %v", err)
	}
}

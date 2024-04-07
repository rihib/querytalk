package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
	"golang.org/x/exp/slog"
)

// FIXME: Use go routines to call LLM asynchronously
// FIXME: Clean Architecture

func gpt4(sysPrompt string, userPrompt string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		slog.Error("error loading .env file", "msg", err.Error())
		return "", fmt.Errorf("error loading .env file: %v", err)
	}

	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")
	c := openai.NewClient(OPENAI_API_KEY)

	resp, err := c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: sysPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userPrompt,
				},
			},
		},
	)
	if err != nil {
		slog.Warn("openai chat completion error", "msg", err.Error())
		return "", fmt.Errorf("openai chat completion error: %v", err)
	}

	output := resp.Choices[0].Message.Content
	slog.Info("openai chat completion response", "output", output)
	return output, nil
}

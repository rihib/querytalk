package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func gpt4(prompt string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")

	client := openai.NewClient(OPENAI_API_KEY)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("ChatCompletion error: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
}

func prompt2sql(prompt string) string {
	sql, ok := gpt4(prompt)
	if ok != nil {
		return "Error: Could not generate SQL"
	}
	return sql
}

func getData(sql string) string {
	data := sql
	return data
}

func createVisualizableData(data string) string {
	visualizableData := data
	return visualizableData
}

func getVisualizableData(prompt string) string {
	sql := prompt2sql(prompt)
	d := getData(sql)
	vd := createVisualizableData(d)
	return vd
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func readSchema() string {
	filePath := "schema.yaml"
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	schema := string(content)
	return schema
}

func createSysPrompt() string {
	schema := readSchema()
	sysPrompt := fmt.Sprintf("This is the database schema:\n\n```\n%s\n```\n\n", schema)
	return sysPrompt
}

func createUserPrompt(prompt string) string {
	userPrompt := fmt.Sprintf(
		`Create a mysql 5.7 query for '%s'.
		Use only tables, columns and relationships in the database schema.
		Answer in short.`, prompt)
	return userPrompt
}

func gpt4(sysPrompt string, userPrompt string) (string, error) {
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
		return "", fmt.Errorf("ChatCompletion error: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
}

func extarctSQL(text string) string {
	pattern := regexp.MustCompile("(?s)```sql\n(.+?)\n```")
	matches := pattern.FindStringSubmatch(text)

	if len(matches) <= 1 {
		fmt.Println(text)
		return "Error: No SQL snippet found"
	}
	return matches[1]
}

func prompt2sql(sysPrompt string, userPrompt string) string {
	text, ok := gpt4(sysPrompt, userPrompt)
	sql := extarctSQL(text)
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
	sp := createSysPrompt()
	up := createUserPrompt(prompt)
	sql := prompt2sql(sp, up)
	d := getData(sql)
	vd := createVisualizableData(d)
	return vd
}

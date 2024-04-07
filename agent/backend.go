package main

import (
	"fmt"
	"log/slog"
	"regexp"
)

// FIXME: Clean Architecture

const (
	SYS_PROMPT  = "This is the database schema:\n\n```\n%s\n```\n\n"
	USER_PROMPT = `
		Given the database schema, create a %s query for '%s'.
		- Use only the tables, columns, and relationships described in the schema.
		- Provide the query in a short, concise format.
	` + "- Enclose the SQL query within triple backticks (```) and prefix it with 'sql', like so: ```sql"
)

func createSysPrompt(schema string) string {
	return fmt.Sprintf(SYS_PROMPT, schema)
}

func createUserPrompt(dbType string, prompt string) string {
	return fmt.Sprintf(USER_PROMPT, dbType, prompt)
}

func extarctQuery(output string) (string, error) {
	pattern := regexp.MustCompile("(?s)```sql\n(.+?)\n```")
	matches := pattern.FindStringSubmatch(output)

	if len(matches) <= 1 {
		slog.Info("query not found", "output", output)
		return "", fmt.Errorf("query not found")
	}

	return matches[1], nil
}

func createQuery(dbType string, schema string, prompt string) (string, error) {
	sp := createSysPrompt(schema)
	up := createUserPrompt(dbType, prompt)

	output, err := gpt4(sp, up)
	if err != nil {
		return "", err
	}

	query, err := extarctQuery(output)
	if err != nil {
		return "", err
	}

	return query, nil
}

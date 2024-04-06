package main

import (
	"fmt"
	"regexp"
)

const (
	SYS_PROMPT  = "This is the database schema:\n\n```\n%s\n```\n\n"
	USER_PROMPT = `
		Create a mysql 5.7 query for '%s'.
		Use only tables, columns and relationships in the database schema.
		Answer in short.
	`
)

func createSysPrompt(schema string) string {
	return fmt.Sprintf(SYS_PROMPT, schema)
}

func createUserPrompt(prompt string) string {
	return fmt.Sprintf(USER_PROMPT, prompt)
}

func extarctQuery(text string) string {
	pattern := regexp.MustCompile("(?s)```sql\n(.+?)\n```")
	matches := pattern.FindStringSubmatch(text)

	if len(matches) <= 1 {
		return "Error: No SQL snippet found"
	}
	return matches[1]
}

func createQuery(schema string, prompt string) string {
	sp := createSysPrompt(schema)
	up := createUserPrompt(prompt)
	text, ok := gpt4(sp, up)
	sql := extarctQuery(text)
	if ok != nil {
		return "Error: Could not generate SQL"
	}
	return sql
}

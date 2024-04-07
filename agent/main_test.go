package main_test

import (
	"context"
	"fmt"
	"testing"

	ogen "querychat/ogen"
)

// TODO: Create comprehensive tests for chat API

const (
	BASEURL  = "http://localhost:8080"
	DB_TYPE  = "SQLite3"
	PROMPT   = "How many customers?"
	EXPECTED = `{"rows":[{"COUNT(*)":59}]}`
)

func run(ctx context.Context, dbType string, prompt string) (*ogen.VisualizableData, error) {
	c, err := ogen.NewClient(BASEURL)
	if err != nil {
		return nil, fmt.Errorf("create client: %w", err)
	}

	request := ogen.OptMSG{
		Value: ogen.MSG{
			DbType: ogen.OptString{Value: dbType, Set: true},
			Prompt: ogen.OptString{Value: prompt, Set: true},
		},
		Set: true,
	}
	r, err := c.SendMSG(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("send prompt: %w", err)
	}

	return r, nil
}

func TestChatAPI(t *testing.T) {
	r, err := run(context.Background(), DB_TYPE, PROMPT)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	if r.VisualizableData != EXPECTED {
		t.Errorf("test failed: got `%v` expected `%v`", r.VisualizableData, EXPECTED)
	}
}

package main_test

import (
	"context"
	"fmt"
	"testing"

	ogen "querychat/ogen"
)

func run(ctx context.Context) (*ogen.VisualizableData, error) {
	BaseURL := "http://localhost:8080"

	var res *ogen.VisualizableData

	client, err := ogen.NewClient(BaseURL)
	if err != nil {
		return res, fmt.Errorf("create client: %w", err)
	}

	request := ogen.OptPrompt{Value: ogen.Prompt{Prompt: "How many users?"}, Set: true}
	res, err = client.SendPrompt(ctx, request)
	if err != nil {
		return res, fmt.Errorf("send prompt: %w", err)
	}

	return res, nil
}

func TestChatAPI(t *testing.T) {
	res, err := run(context.Background())
	if err != nil {
		t.Fatalf("Unexpected error: %+v", err)
	}

	expected := "SELECT COUNT(*) FROM users;"
	if res.VisualizableData != expected {
		t.Errorf("Expected: %v, got: %v", expected, res.VisualizableData)
	}
}

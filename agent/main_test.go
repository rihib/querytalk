package main_test

import (
	"context"
	"fmt"
	"testing"

	ogen "querychat/ogen"
)

const (
	BASEURL = "http://localhost:8080"
)

func run(ctx context.Context, prompt string) (*ogen.VisualizableData, error) {
	var r *ogen.VisualizableData

	c, err := ogen.NewClient(BASEURL)
	if err != nil {
		return r, fmt.Errorf("create client: %w", err)
	}

	request := ogen.OptPrompt{Value: ogen.Prompt{Prompt: prompt}, Set: true}
	r, err = c.SendPrompt(ctx, request)
	if err != nil {
		return r, fmt.Errorf("send prompt: %w", err)
	}

	return r, nil
}

func TestChatAPI(t *testing.T) {
	p := "How many users?"
	e := "SELECT COUNT(*) FROM users;"

	r, err := run(context.Background(), p)
	if err != nil {
		t.Fatalf("Unexpected error: %+v", err)
	}

	if r.VisualizableData != e {
		t.Errorf("Test Failed: got `%v` expected `%v`", r.VisualizableData, e)
	}
}

package main_test

import (
	"context"
	"fmt"
	"testing"

	ogen "querychat/ogen"
)

// TODO: Create comprehensive tests for chat API

const (
	BASEURL = "http://localhost:8080"
)

func run(ctx context.Context, dbType string, prompt string) (*ogen.VisualizableData, error) {
	var r *ogen.VisualizableData

	c, err := ogen.NewClient(BASEURL)
	if err != nil {
		return r, fmt.Errorf("create client: %w", err)
	}

	request := ogen.OptMSG{
		Value: ogen.MSG{
			DbType: ogen.OptString{Value: dbType, Set: true},
			Prompt: ogen.OptString{Value: prompt, Set: true},
		},
		Set: true,
	}
	r, err = c.SendMSG(ctx, request)
	if err != nil {
		return r, fmt.Errorf("send prompt: %w", err)
	}

	return r, nil
}

func TestChatAPI(t *testing.T) {
	d := "SQLite3"
	p := "How many customers?"
	e := "SELECT COUNT(*) FROM customers;"

	r, err := run(context.Background(), d, p)
	if err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	if r.VisualizableData != e {
		t.Errorf("test failed: got `%v` expected `%v`", r.VisualizableData, e)
	}
}

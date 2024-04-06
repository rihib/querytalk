package main

import (
	"context"
	"fmt"
	"log/slog"

	ogen "querychat/ogen"
)

type chatService struct{}

func (s *chatService) SendPrompt(ctx context.Context, req ogen.OptPrompt) (*ogen.VisualizableData, error) {
	var res ogen.VisualizableData

	if !req.Set {
		slog.Info("prompt not set")
		return &res, fmt.Errorf("prompt not set")
	}

	vd, err := getVisualizableData(req.Value.Prompt)
	if err != nil {
		return &res, err
	}

	res.VisualizableData = vd
	return &res, nil
}

func (s *chatService) NewError(ctx context.Context, err error) *ogen.ErrorStatusCode {
	slog.Error("internal server error", "message", err.Error())
	return &ogen.ErrorStatusCode{StatusCode: 500, Response: ogen.Error{Code: 0, Message: "internal server error"}}
}

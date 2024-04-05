package main

import (
	"context"
	"fmt"

	ogen "querychat/ogen"
)

type chatService struct{}

func (s *chatService) SendPrompt(ctx context.Context, req ogen.OptPrompt) (*ogen.VisualizableData, error) {
	var res ogen.VisualizableData

	if !req.Set {
		return &res, fmt.Errorf("bad request")
	}

	res.VisualizableData = getVisualizableData(req.Value.Prompt)
	return &res, nil
}

func (s *chatService) NewError(ctx context.Context, err error) *ogen.ErrorStatusCode {
	return &ogen.ErrorStatusCode{StatusCode: 500, Response: ogen.Error{Message: err.Error()}}
}

package main

import (
	"context"
	"fmt"

	ogen "querytalk/ogen"
)

type talkService struct{}

func (s *talkService) SendPrompt(ctx context.Context, req ogen.OptPrompt) (*ogen.VisualizableData, error) {
	var res ogen.VisualizableData

	if !req.Set {
		return &res, fmt.Errorf("bad request")
	}

	res.VisualizableData = getVisualizableData(req.Value.Prompt)
	return &res, nil
}

func (s *talkService) NewError(ctx context.Context, err error) *ogen.ErrorStatusCode {
	return &ogen.ErrorStatusCode{StatusCode: 500, Response: ogen.Error{Message: err.Error()}}
}

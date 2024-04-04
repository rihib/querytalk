package main

import (
	"context"
	"log/slog"
	"net/http"

	ogen "querytalk/ogen"
)

type talkService struct{}
type BadRequestError struct{}

func (e BadRequestError) Error() string {
	return "Bad Request"
}

func (s *talkService) SendPrompt(ctx context.Context, req ogen.OptPrompt) (*ogen.VisualizableData, error) {
	var res ogen.VisualizableData

	if !req.Set {
		return &res, BadRequestError{}
	}

	res.VisualizableData = getVisualizableData(req.Value.Prompt)
	return &res, nil
}

func (s *talkService) NewError(ctx context.Context, err error) *ogen.ErrorStatusCode {
	return &ogen.ErrorStatusCode{StatusCode: 500, Response: ogen.Error{Message: err.Error()}}
}

func main() {
	service := &talkService{}
	srv, err := ogen.NewServer(service)
	if err != nil {
		slog.Error(err.Error())
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		slog.Error(err.Error())
	}
}

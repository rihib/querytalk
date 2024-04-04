package main

import (
	"context"
	"log/slog"
	"net/http"

	ogen "querytalk/ogen"
)

type talkService struct{}

func (s *talkService) SendPrompt(ctx context.Context, req ogen.OptPrompt) (*ogen.VisualizableData, error) {
	// if !req.Set {
	// 	return nil, errors.New("no prompt provided")
	// }

	var res ogen.VisualizableData
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

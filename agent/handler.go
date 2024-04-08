package main

import (
	"context"
	"fmt"
	"log/slog"

	ogen "querychat/ogen"
)

type chatService struct{}

func (s *chatService) SendMSG(ctx context.Context, req ogen.OptMSG) (*ogen.VisualizableDataHeaders, error) {
	if !req.Set {
		slog.Info("prompt not set")
		return nil, fmt.Errorf("prompt not set")
	}

	vd, err := getVisualizableData(req.Value.DbType.Value, req.Value.Prompt.Value)
	if err != nil {
		return nil, err
	}

	var res ogen.VisualizableDataHeaders

	res.AccessControlAllowHeaders.Value = "Content-Type"
	res.AccessControlAllowHeaders.Set = true

	res.AccessControlAllowMethods.Value = "POST, GET, OPTIONS"
	res.AccessControlAllowMethods.Set = true

	res.AccessControlAllowOrigin.Value = "http://localhost:3000"
	res.AccessControlAllowOrigin.Set = true

	res.AccessControlMaxAge.Value = 3600
	res.AccessControlMaxAge.Set = true

	res.Response.VisualizableData = vd
	return &res, nil
}

func (s *chatService) NewError(ctx context.Context, err error) *ogen.ErrorStatusCode {
	slog.Error("failed to visualize data", "msg", err.Error())
	return &ogen.ErrorStatusCode{StatusCode: 500, Response: ogen.Error{Code: 0, Message: "Failed to visualize data"}}
}

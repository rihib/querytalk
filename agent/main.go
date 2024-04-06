package main

import (
	"log/slog"
	"net/http"

	ogen "querychat/ogen"
)

func main() {
	service := &chatService{}
	srv, err := ogen.NewServer(service)
	if err != nil {
		slog.Error(err.Error())
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		slog.Error(err.Error())
	}
}
package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	ogen "querychat/ogen"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func main() {
	service := &chatService{}

	s, err := ogen.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(":8080", s)
	if err != nil {
		log.Fatal(err)
	}
}

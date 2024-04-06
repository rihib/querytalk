package main

import (
	"log"
	"net/http"

	ogen "querychat/ogen"
)

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

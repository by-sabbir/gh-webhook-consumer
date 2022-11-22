package main

import (
	"log"

	httpTransport "github.com/by-sabbir/gh-webhook-consumer/internal/http"
)

func Run() error {
	httpHandler := httpTransport.NewHandler()

	httpHandler.Serve()
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatalf("could not run the service: %+v\n", err)
	}
}

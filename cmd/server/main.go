package main

import (
	"github.com/by-sabbir/gh-webhook-consumer/internal/db"
	"github.com/by-sabbir/gh-webhook-consumer/internal/events"
	httpTransport "github.com/by-sabbir/gh-webhook-consumer/internal/http"
	"github.com/sirupsen/logrus"
)

func Run() error {
	db, err := db.NewDatabase()
	if err != nil {
		panic(err)
	}

	svc := events.NewService(db)
	httpHandler := httpTransport.NewHandler(svc)

	httpHandler.Serve()
	return nil
}

func main() {
	if err := Run(); err != nil {
		logrus.Fatalf("could not run the service: %+v\n", err)
	}
}

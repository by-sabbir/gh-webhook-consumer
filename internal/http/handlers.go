package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Service EventServices
	Server  *http.Server
}

func NewHandler(svc EventServices) *Handler {
	h := &Handler{
		Service: svc,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()

	h.Server = &http.Server{
		Addr:         "0.0.0.0:3000",
		Handler:      h.Router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	h.Router.Use(LogMiddlewire)
	h.Router.Use(JSONMiddlewire)

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/webhook", h.GetIssues).Methods("POST")
	h.Router.HandleFunc("/healthz", healthCheck).Methods("GET")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}
func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Fatalf("error serving http%+v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("Shut down gracefully...")
	return nil
}

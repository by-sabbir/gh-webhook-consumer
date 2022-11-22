package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	hook "github.com/by-sabbir/gh-webhook-consumer/internal/hooks"
)

type EventServices interface {
	GetIssues(context.Context, hook.GitIssues) (string, error)
}

func (h *Handler) GetIssues(w http.ResponseWriter, r *http.Request) {
	var hook hook.GitIssues
	if err := json.NewDecoder(r.Body).Decode(&hook); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	insertedId, err := h.Service.GetIssues(context.Background(), hook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	log.Println("GetHooks")
	json.NewEncoder(w).Encode(map[string]string{
		"docId": insertedId,
	})
}

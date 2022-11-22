package http

import (
	"encoding/json"
	"log"
	"net/http"

	hook "github.com/by-sabbir/gh-webhook-consumer/internal/hooks"
)

func (h *Handler) GetHooks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	log.Println("GetHooks")
	var hook hook.GitIssues
	json.NewDecoder(r.Body).Decode(&hook)
	json.NewEncoder(w).Encode(hook)
}

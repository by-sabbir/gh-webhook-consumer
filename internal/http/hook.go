package http

import (
	"encoding/json"
	"log"
	"net/http"

	hook "github.com/by-sabbir/gh-webhook-consumer/internal/hooks"
)

func (h *Handler) GetHooks(w http.ResponseWriter, r *http.Request) {
	log.Println("GetHooks")
	var hook hook.AllEvents
	json.NewDecoder(r.Body).Decode(&hook)
	json.NewEncoder(w).Encode(hook.Commits)
}

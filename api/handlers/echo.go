package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"reminder_bot/api/types"
)

func (h *HTTPHandler) Echo(w http.ResponseWriter, r *http.Request) {

	var update types.Update

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Printf("error: %q", err)
		return
	}

	if update.Message == nil {
		return
	}

	if update.Message.From.Username != "" && update.Message.Text != "" {
		log.Printf("[%s] %s", update.Message.From.Username, update.Message.Text)
	}

	if update.Message.Chat != nil && update.Message.Text != "" {
		h.Service.SendMessage(int(update.Message.Chat.ID), update.Message.Text)
	}

}

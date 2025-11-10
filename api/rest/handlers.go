package api

import (
	"encoding/json"
	"log"
	"net/http"
	"reminder_bot/api/types"
)

func (s *HTTPServer) MessageHandler(w http.ResponseWriter, r *http.Request) {

	var update types.Update

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Printf("error: %q", err)
		return
	}

	if update.Message != nil {
		s.MessageProcessor.Echo(update.Message)
	}

}

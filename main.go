package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type UpdateResponse struct {
	Message *Message `json:"message,omitempty"`
}

type Message struct {
	ID   int    `json:"message_id"`
	From *User  `json:"from,omitempty"`
	Chat *Chat  `json:"chat,omitempty"`
	Text string `json:"text,omitempty"`
}

type User struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name,omitempty"`
	Username     string `json:"username,omitempty"`
	LanguageCode string `json:"language_code,omitempty"`
}

type Chat struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title,omitempty"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

func HandleWebhook(w http.ResponseWriter, r *http.Request) {

	var update UpdateResponse

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if update.Message == nil {
		return
	}

	if update.Message.From.Username != "" && update.Message.Text != "" {
		log.Printf("[%s] %s", update.Message.From.Username, update.Message.Text)
	}

}

func main() {

	r := chi.NewRouter()

	r.Get("/", HandleWebhook)

	fmt.Println("Запускаем сервер")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
	fmt.Println("Завершаем работу")
}

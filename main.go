package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Update struct {
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

	var update Update

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
		sendMessage(int(update.Message.Chat.ID), update.Message.Text)
	}

}

type SendMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func sendMessage(chatID int, message string) {
	url := "https://api.telegram.org/bot8217269610:AAHInG-LgHmjjU-ET9qMwNO2EfwBT7ac2O0/sendMessage"

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(SendMessage{
		ChatID: int64(chatID),
		Text:   message,
	})

	resp, err := http.Post(url, "application/json", buf)
	if err != nil {
		log.Printf("error :%q", err)
	}
	defer resp.Body.Close()

	log.Println("Статус: ", resp.Status)

}
func main() {

	r := chi.NewRouter()

	r.Post("/", HandleWebhook)

	fmt.Println("Запускаем сервер")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
	fmt.Println("Завершаем работу")
}

package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"reminder_bot/api/types"
)

type MessageProcessor struct{}

func NewMessageProcessor() *MessageProcessor {
	return &MessageProcessor{}
}

var URL = "https://api.telegram.org/bot8217269610:AAHInG-LgHmjjU-ET9qMwNO2EfwBT7ac2O0"

func (mp *MessageProcessor) Echo(message *types.Message) {

	if message.From.Username != "" && message.Text != "" {
		log.Printf("[%s] %s", message.From.Username, message.Text)

	}

	if message.Chat != nil && message.Text != "" {
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(types.SendMessage{
			ChatID: int64(message.Chat.ID),
			Text:   message.Text + " 10.07.2025",
		})

		resp, err := http.Post(URL+"/sendMessage", "application/json", buf)
		if err != nil {
			log.Printf("error :%q", err)
		}
		defer resp.Body.Close()

		log.Println("Статус: ", resp.Status)
	}

}

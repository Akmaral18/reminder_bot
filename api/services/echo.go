package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"reminder_bot/api/types"
)

var URL = "https://api.telegram.org/bot8217269610:AAHInG-LgHmjjU-ET9qMwNO2EfwBT7ac2O0"

func (s *MyService) SendMessage(chatID int, message string) {

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(types.SendMessage{
		ChatID: int64(chatID),
		Text:   message,
	})

	resp, err := http.Post(URL+"/sendMessage", "application/json", buf)
	if err != nil {
		log.Printf("error :%q", err)
	}
	defer resp.Body.Close()

	log.Println("Статус: ", resp.Status)

}

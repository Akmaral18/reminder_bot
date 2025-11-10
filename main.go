package main

import (
	api "reminder_bot/api/rest"
	"reminder_bot/internal/services"
)

func main() {
	mp := services.NewMessageProcessor()

	server := &api.HTTPServer{
		Addr:             ":8080",
		MessageProcessor: mp,
	}

	server.Run()
}

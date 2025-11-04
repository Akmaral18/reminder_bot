package main

import (
	"reminder_bot/api"
	"reminder_bot/api/handlers"
	"reminder_bot/api/services"
)

func main() {
	service := &services.MyService{}
	handler := &handlers.HTTPHandler{Service: service}

	server := &api.HTTPServer{
		Addr:  ":8080",
		Hndlr: handler,
	}

	server.Run()
}

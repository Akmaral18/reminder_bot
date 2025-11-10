package api

import (
	"fmt"
	"net/http"

	"reminder_bot/internal/services"

	"github.com/go-chi/chi"
)

type Server interface {
	Run()
}

type HTTPServer struct {
	Addr             string
	MessageProcessor *services.MessageProcessor
}

func (s *HTTPServer) Run() {
	r := chi.NewRouter()

	r.Post("/", s.MessageHandler)

	fmt.Println("Запускаем сервер")
	err := http.ListenAndServe(s.Addr, r)
	if err != nil {
		panic(err)
	}
	fmt.Println("Завершаем работу")
}

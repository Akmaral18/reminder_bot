package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type Server interface {
	Run()
}

type HTTPServer struct {
	Addr  string
	Hndlr *handlers.HTTPHandler
}

func (s *HTTPServer) Run() {
	r := chi.NewRouter()

	r.Post("/", Echo)

	fmt.Println("Запускаем сервер")
	err := http.ListenAndServe(s.Addr, r)
	if err != nil {
		panic(err)
	}
	fmt.Println("Завершаем работу")
}

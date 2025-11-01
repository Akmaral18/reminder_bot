package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	fmt.Println("Запускаем сервер")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
	fmt.Println("Завершаем работу")
}

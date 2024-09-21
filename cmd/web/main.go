package main

import (
	"fmt"
	"log"
	"net/http"
)

// Адрес и порт для запуска сервера
var addr string = "localhost"
var port string = "8080"

func main() {
	// Роуты
	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Printf("Запуск веб-сервера на http://%s:%s\n", addr, port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", addr, port), mux); err != nil {
		log.Fatal(err)
	}
}

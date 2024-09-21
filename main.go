package main

import (
	"fmt"
	"log"
	"net/http"
)

var addr string = "localhost"
var port string = "8080"

func home(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	res.Write([]byte("Домашняя страница"))
}

func showSnippet(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Отображение заметки"))
}

func createSnippet(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.Header().Set("Allow", http.MethodPost)
		http.Error(res, "Метод запрещен!", http.StatusMethodNotAllowed)
		return
	}
	res.Write([]byte("Форма создания новой заметки"))
}

func main() {
	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Printf("Запуск веб-сервера на http://%s:%s\n", addr, port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", addr, port), mux); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
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

	// Роут для обработки статических файлов (css, js, images)
	var fileServer http.Handler = http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Запуск веб-сервера на http://%s:%s\n", addr, port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", addr, port), mux); err != nil {
		log.Fatal(err)
	}
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}
	return f, nil
}

package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Флаги для командой строки
	addr := flag.String("addr", "localhost:8080", "Сетевой адрес HTTP")
	flag.Parse()

	// Логгеры
	var infoLog *log.Logger = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	var errorLog *log.Logger = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Инициализация структуры с зависимостями
	var app *application = &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Роуты
	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// Роут для обработки статических файлов (css, js, images)
	var fileServer http.Handler = http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Инициализация HTTP сервера
	var httpServer *http.Server = &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// Запуск HTTP сервера
	infoLog.Printf("Запуск веб-сервера на http://%s\n", *addr)
	if err := httpServer.ListenAndServe(); err != nil {
		errorLog.Fatal(err)
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

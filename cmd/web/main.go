package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Флаги для командой строки
	var addr *string = flag.String("addr", "localhost:8080", "Сетевой адрес HTTP")
	flag.Parse()

	// Логгеры
	var infoLog *log.Logger = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	var errorLog *log.Logger = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Инициализация структуры с зависимостями
	var app *application = &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Получение роутов
	var mux *http.ServeMux = app.routes()

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

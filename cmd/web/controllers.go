package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// Контроллер домашней страницы
func (app *application) home(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	var files []string = []string{
		"./ui/html/home.html",
		"./ui/html/base.html",
		"./ui/html/footer.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err = ts.Execute(res, nil); err != nil {
		app.errorLog.Println(err.Error())
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Контроллер показа заметки по id в URL
func (app *application) showSnippet(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(res, req)
		return
	}
	res.Write([]byte(fmt.Sprintf("Заметка с ID: %d\n", id)))
	// Также можно использовать fmt.Fprintf(),
	// так как http.ResponseWriter удовлетворяет интерфейсу io.Writer
	// fmt.Fprintf(res, "Заметка с ID: %d\n", id)
}

// Контроллер создания заметки (только POST)
func (app *application) createSnippet(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.Header().Set("Allow", http.MethodPost)
		http.Error(res, "Метод запрещен!", http.StatusMethodNotAllowed)
		return
	}
	res.Write([]byte("Форма создания новой заметки"))
}

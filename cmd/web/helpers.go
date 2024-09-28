package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Обработка серверных ошибок (Internal Server Error)
func (app *application) serverError(res http.ResponseWriter, err error) {
	var trace string = fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Обработка клиентских ошибок
func (app *application) clientError(res http.ResponseWriter, status int) {
	http.Error(res, http.StatusText(status), status)
}

// Обработка ошибки Not found (404)
func (app *application) notFound(res http.ResponseWriter) {
	app.clientError(res, http.StatusNotFound)
}

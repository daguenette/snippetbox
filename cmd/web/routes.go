package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.homeHandler))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetViewHandler))
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.snippetFormHandler))
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.snippetCreateHandler))
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}

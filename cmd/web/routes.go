package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.addIPToContext)
	mux.Use(app.Session.LoadAndSave)

	mux.Get("/", app.Home)
	mux.Post("/login", app.Login)

	mux.Route("/user", func(r chi.Router) {
		r.Use(app.auth)
		r.Get("/profile", app.Profile)
	})
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

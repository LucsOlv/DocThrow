package main

import (
	"github.com/LucsOliv/DocThrow/web/templates"
	"github.com/a-h/templ"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/hello", templ.Handler(templates.HomePage()).ServeHTTP)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("welcome"))
		if err != nil {
			return
		}
	})
	_ = http.ListenAndServe(":3000", r)
}

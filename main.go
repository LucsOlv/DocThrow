package main

import (
	"github.com/LucsOliv/DocThrow/web/layout"
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	//add public assets
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	// access to public css
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	r.Get("/hello", templ.Handler(layout.BaseLayout()).ServeHTTP)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("welcome"))
		if err != nil {
			return
		}
	})
	_ = http.ListenAndServe(":3000", r)
}

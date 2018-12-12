package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type JSON
		middleware.RequestID,                          // adds request ID
		middleware.Logger,                             // logs request details
		middleware.DefaultCompress,                    // compress results
		middleware.RedirectSlashes,                    // redirect slashes to no slash version
		middleware.Recoverer,                          // recover from panics without crashing
	)

	router.Route("/v1/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome"))
		})
	})

	return router
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := Routes()

	log.Fatal(http.ListenAndServe(":"+port, router))
}

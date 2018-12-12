package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":"+port, r)
}

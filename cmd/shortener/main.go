package main

import (
	"net/http"

	"github.com/Far04ka/LinkShortener/internal/app/endpoints"
	"github.com/Far04ka/LinkShortener/internal/storage"
	"github.com/go-chi/chi/v5"
)

func main() {
	storage.Storage.Lnks = make(map[string]string)

	router := chi.NewRouter()

	router.Get("/{id}", endpoints.GetURL(&storage.Storage))
	router.Post("/", endpoints.PostURL(&storage.Storage))

	http.ListenAndServe(":8080", router)
}

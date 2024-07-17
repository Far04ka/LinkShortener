package main

import (
	"net/http"

	"github.com/Far04ka/LinkShortener/internal/app/endpoints"
	"github.com/Far04ka/LinkShortener/internal/storage"
)

func main() {
	storage.Storage.Lnks = make(map[string]string)

	mux := http.NewServeMux()

	mux.HandleFunc("/", endpoints.Router)

	http.ListenAndServe(":8080", mux)
}

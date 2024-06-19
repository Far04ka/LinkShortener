package main

import (
	"net/http"

	"github.com/Far04ka/LinkShortener/internal/app/endpoints"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", endpoints.Router)

	http.ListenAndServe(":8080", mux)
}

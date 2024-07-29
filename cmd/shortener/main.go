package main

import (
	"net/http"

	conf "github.com/Far04ka/LinkShortener/internal"
	"github.com/Far04ka/LinkShortener/internal/app/endpoints"
	"github.com/Far04ka/LinkShortener/internal/storage"
	"github.com/go-chi/chi/v5"
)

func init() {
	conf.Conf = &conf.Config{Addr: &conf.AddrField{Val: "localhost:8080"}, Finaladdr: &conf.FinalAddrField{}}
	conf.CreateConfig()
	storage.Storage.Lnks = make(map[string]string)
}

func main() {
	router := chi.NewRouter()

	router.Get("/"+conf.Conf.Finaladdr.Val+"{id}", endpoints.GetURL(&storage.Storage))
	router.Post("/", endpoints.PostURL(&storage.Storage))

	http.ListenAndServe(conf.Conf.Addr.Val, router)
}

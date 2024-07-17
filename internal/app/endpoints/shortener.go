package endpoints

import (
	"io"
	"net/http"

	"github.com/Far04ka/LinkShortener/internal/storage"
	"github.com/teris-io/shortid"
)

func Router(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetURL(&storage.Storage).ServeHTTP(w, r)
	} else if r.Method == http.MethodPost {
		PostURL(&storage.Storage).ServeHTTP(w, r)
	} else {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
}

func GetURL(stor storage.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqURL := r.URL.Path[1:]
		if len(reqURL) == 0 {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		redir, err := stor.GetLink(reqURL)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, redir, http.StatusTemporaryRedirect)
	}
}
func PostURL(stor storage.Repo) http.HandlerFunc {
	storage.Storage.Lnks = make(map[string]string)
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		req, err := io.ReadAll(r.Body)
		url := string(req)

		if len(req) == 0 || err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		id := stor.GetID(url)
		if id != "" {
			w.WriteHeader(http.StatusCreated)
			io.WriteString(w, storage.URL+id)
			return
		}

		id, _ = shortid.Generate()
		flag := stor.SetLink(id, url)
		for !flag {
			id, _ = shortid.Generate()
			flag = stor.SetLink(id, url)
		}

		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, storage.URL+id)

	}
}

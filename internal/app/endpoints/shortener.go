package endpoints

import (
	"io"
	"net/http"

	conf "github.com/Far04ka/LinkShortener/internal"
	"github.com/Far04ka/LinkShortener/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/teris-io/shortid"
)

func GetURL(stor storage.Repo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqURL := chi.URLParam(r, "id")
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
	return func(w http.ResponseWriter, r *http.Request) {

		req, err := io.ReadAll(r.Body)
		url := string(req)

		if len(req) == 0 || err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		id := stor.GetID(url)
		if id != "" {
			w.WriteHeader(http.StatusCreated)
			io.WriteString(w, conf.Conf.Addr.Val+"/"+conf.Conf.Finaladdr.Val+id)
			return
		}

		id, _ = shortid.Generate()
		flag := stor.SetLink(id, url)
		for !flag {
			id, _ = shortid.Generate()
			flag = stor.SetLink(id, url)
		}

		w.WriteHeader(http.StatusCreated)
		io.WriteString(w, conf.Conf.Addr.Val+"/"+conf.Conf.Finaladdr.Val+id)

	}
}

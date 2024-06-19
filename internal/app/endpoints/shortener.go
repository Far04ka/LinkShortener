package endpoints

import (
	"io"
	"net/http"
	"strings"

	"github.com/Far04ka/LinkShortener/internal/storage"
	"github.com/teris-io/shortid"
)

func Router(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.HandlerFunc(GetURL).ServeHTTP(w, r)
	} else if r.Method == http.MethodPost {
		http.HandlerFunc(PostURL).ServeHTTP(w, r)
	} else {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
}

func GetURL(w http.ResponseWriter, r *http.Request) {
}

func PostURL(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	req, err := io.ReadAll(r.Body)
	url := strings.Split(string(req), "//")[1]
	if len(req) == 0 || err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	for key, val := range storage.Lnks {
		if key == url {
			io.WriteString(w, storage.URL+val)
			return
		}
	}

	id := ""

	for len(id) == 0 {
		id, _ = shortid.Generate()
		for _, val := range storage.Lnks {
			if val == id {
				id = ""
				break
			}
		}
	}
	storage.Lnks[url] = id

	shortURL := storage.URL + id
	io.WriteString(w, shortURL)

}

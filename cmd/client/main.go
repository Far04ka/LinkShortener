package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Far04ka/LinkShortener/internal/storage"
)

func PostClient(client *http.Client, path string) map[string]string {
	answ := make(map[string]string)
	req, err := http.NewRequest(http.MethodPost, storage.URL, strings.NewReader(path))
	req.Header.Add("Content-Type", "text/html")
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	answ["Status"] = resp.Status

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	answ["Body"] = string(respBody)
	return answ
}

func GetClient(client *http.Client, path string) {
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		panic(err)
	}
	client.Do(req)
}

func main() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Printf("Redirect to: %v\n", req.URL)
			return nil
		},
	}

	mapp := PostClient(client, "https://pkg.go.dev/io#Reader.Read")
	fmt.Println(mapp["Body"])
	GetClient(client, mapp["Body"])
}

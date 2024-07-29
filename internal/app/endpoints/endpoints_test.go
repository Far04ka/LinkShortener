package endpoints

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	conf "github.com/Far04ka/LinkShortener/internal"
	"github.com/Far04ka/LinkShortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetURL(t *testing.T) {
	conf.Conf = &conf.Config{Addr: &conf.AddrField{Val: "localhost:8080"}, Finaladdr: &conf.FinalAddrField{Val: "http://localhost:8080/", ShortAddr: "/"}}
	type want struct {
		value      string
		statusCode int
	}

	tests := []struct {
		name   string
		method string
		val    string
		want   want
	}{

		{
			name:   "Simple GET err test",
			method: http.MethodGet,
			val:    "0000",
			want: want{
				value:      "bad request\n",
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name:   "Empty GET test",
			method: http.MethodGet,
			val:    "",
			want: want{
				value:      "bad request\n",
				statusCode: http.StatusBadRequest,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := httptest.NewRequest(test.method, conf.Conf.Finaladdr.Val+test.val, nil)
			w := httptest.NewRecorder()
			GetURL(&storage.Storage)(w, r)

			res := w.Result()
			require.Equal(t, test.want.statusCode, res.StatusCode)

			resBody, _ := io.ReadAll(res.Body)
			defer res.Body.Close()

			assert.Equal(t, test.want.value, string(resBody))
		})
	}
}

func TestPostURL(t *testing.T) {
	conf.Conf = &conf.Config{Addr: &conf.AddrField{Val: "localhost:8080"}, Finaladdr: &conf.FinalAddrField{Val: "http://localhost:8080/", ShortAddr: "/"}}
	storage.Storage.Lnks = make(map[string]string)
	type want struct {
		value      string
		statusCode int
	}

	tests := []struct {
		name   string
		method string
		val    string
		want   want
	}{
		{
			name:   "Simple POST test",
			method: http.MethodGet,
			val:    "ya.ru",
			want: want{
				value:      "",
				statusCode: http.StatusCreated,
			},
		},
		{
			name:   "POST err test",
			method: http.MethodPost,
			val:    "",
			want: want{
				value:      "",
				statusCode: http.StatusBadRequest,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reader := strings.NewReader(test.val)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(test.method, conf.Conf.Finaladdr.Val, reader)

			PostURL(&storage.Storage)(w, r)
			res := w.Result()
			defer res.Body.Close()

			require.Equal(t, test.want.statusCode, res.StatusCode)
		})
	}
}

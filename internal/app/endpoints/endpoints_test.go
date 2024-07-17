package endpoints

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Far04ka/LinkShortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetURL(t *testing.T) {
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
			r := httptest.NewRequest(test.method, storage.URL+test.val, nil)
			w := httptest.NewRecorder()
			GetURL(&storage.Storage)(w, r)

			res := w.Result()
			require.Equal(t, test.want.statusCode, res.StatusCode)

			resBody, _ := io.ReadAll(res.Body)

			assert.Equal(t, test.want.value, string(resBody))
		})
	}
}

func TestPostURL(t *testing.T) {
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
			r := httptest.NewRequest(test.method, storage.URL, reader)

			PostURL(&storage.Storage)(w, r)
			res := w.Result()

			require.Equal(t, test.want.statusCode, res.StatusCode)
		})
	}
}

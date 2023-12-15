package link

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/inancgumus/gobyexample/bite/sqlx/sqlxtest"
)

func TestServer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		path     string
		method   string
		wantCode int
	}{
		{path: "/health", method: http.MethodGet, wantCode: http.StatusOK},
		{path: "/notfound", method: http.MethodGet, wantCode: http.StatusNotFound},

		/* exercise: add more test cases here for other routes */
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(tt.method, tt.path, nil)

			NewServer(nil).ServeHTTP(w, r)

			if w.Code != tt.wantCode {
				t.Errorf("got status code = %d, want %d", w.Code, tt.wantCode)
			}
		})
	}
}

func TestHealth(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	Health(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("got status code = %d, want %d", w.Code, http.StatusOK)
	}
	if got := w.Body.String(); !strings.Contains(got, "OK") {
		t.Errorf("got body = %s\twant contains %s", got, "OK")
	}
}

func TestShortenInternalError(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(
		`{"key": "go", "url": "https://go.dev"}`,
	))

	store := NewStore(
		sqlxtest.Dial(context.Background(), t),
	)
	// disconnect from the database to force an error
	_ = store.db.Close()
	Shorten(store).ServeHTTP(w, r)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("got status code = %d, want %d", w.Code, http.StatusInternalServerError)
	}
}

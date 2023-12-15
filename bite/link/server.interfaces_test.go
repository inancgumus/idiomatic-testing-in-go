//go:build exclude

package link

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShorten(t *testing.T) {
	t.Parallel()

	creator := fakeLinkCreator(func(context.Context, Link) error {
		return nil
	})
	handler := shorten(creator)

	body := `{"key": "go", "url": "https://go.dev"}`
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	handler.ServeHTTP(w, r)

	if want := http.StatusCreated; w.Code != want {
		t.Errorf("got status code = %d, want %d", w.Code, want)
	}
}

type fakeLinkCreator func(context.Context, Link) error

func (f fakeLinkCreator) Create(ctx context.Context, link Link) error {
	return f(ctx, link)
}

package link

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServerHealth(t *testing.T) {
	t.Parallel()

	var srv Server
	rec := httptest.NewRecorder()
	srv.Health(rec, httptest.NewRequest(http.MethodGet, "/", nil))

	if rec.Code != http.StatusOK {
		t.Errorf("got status code = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Body.String(); !strings.Contains(got, "OK") {
		t.Errorf("got body = %s\twant contains %s", got, "OK")
	}
}

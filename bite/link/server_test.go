package link

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServerHealth(t *testing.T) {
	t.Parallel()

	var srv Server
	rec := httptest.NewRecorder()
	srv.Health(rec, newRequest(t, http.MethodGet, "/", nil))

	if rec.Code != http.StatusOK {
		t.Errorf("got status code = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Body.String(); !strings.Contains(got, "OK") {
		t.Errorf("got body = %s\twant contains %s", got, "OK")
	}
}

// newRequest creates a new [http.Request] for testing.
// It fails the calling test if it cannot create the request.
func newRequest(
	tb testing.TB, method string, target string, body io.Reader,
) *http.Request {
	tb.Helper()
	r, err := http.NewRequest(method, target, body)
	if err != nil {
		tb.Fatalf("http.NewRequest() err = %v, want nil", err)
	}
	return r
}

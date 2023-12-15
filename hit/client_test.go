package hit

import (
	"context"
	"net/http"
	"testing"
)

func TestClientDoError(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest(http.MethodGet, "/", http.NoBody)
	if err != nil {
		t.Fatalf("new http request: %v", err)
	}

	fail := func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	client := &Client{
		Transport: fakeTripper(fail),
	}
	sum := client.Do(context.Background(), req, 5)
	if got := sum.Errors; got != 5 {
		t.Errorf("Errors=%d; want 5", got)
	}
}

type fakeTripper func(*http.Request) (*http.Response, error)

func (f fakeTripper) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

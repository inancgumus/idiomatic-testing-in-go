package hit

import (
	"context"
	"net/http"
	"sync/atomic"
	"testing"
)

// roundTripperFunc is an adapter to allow the use of ordinary functions as an
// [http.RoundTripper]. If f is a function with the appropriate signature,
// roundTripperFunc(f) is an [http.RoundTripper] that calls f.
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestSendNRequestResultCount(t *testing.T) {
	t.Parallel()

	const wantRequests = 10

	var totalRequests atomic.Int64
	countingRoundTripper := roundTripperFunc(
		func(*http.Request) (*http.Response, error) {
			totalRequests.Add(1)
			return &http.Response{StatusCode: http.StatusOK}, nil
		},
	)

	req, err := http.NewRequest(http.MethodGet, "/any", http.NoBody)
	if err != nil {
		t.Fatalf("http.NewRequest() err = %q, want <nil>", err)
	}
	results, err := SendN(context.Background(), wantRequests, req, &Options{
		Client:      &http.Client{Transport: countingRoundTripper},
		Concurrency: 2,
	})
	if err != nil {
		t.Fatalf("SendN() err = %q, want <nil>", err)
	}

	var totalResults int
	for range results {
		totalResults++
	}

	if got := totalRequests.Load(); got != wantRequests {
		t.Fatalf("totalRequests = %d, want %d", got, wantRequests)
	}
	if totalResults != wantRequests {
		t.Fatalf("totalResults = %d, want %d", totalResults, wantRequests)
	}
}

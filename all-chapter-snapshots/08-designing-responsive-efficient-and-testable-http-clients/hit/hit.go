package hit

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// SendN sends n HTTP requests using the [Options] configuration and the
// provided [http.Request]. It returns a [Results] iterator for each [Result].
// The iterator stops when the consumer stops or when all requests are done.
func SendN(
	ctx context.Context, n int, req *http.Request, opts *Options,
) (Results, error) {
	opts = opts.clone().setDefaults()
	if n <= 0 {
		return nil, fmt.Errorf("invalid request count: %d", n)
	}

	ctx, cancel := context.WithCancel(ctx)
	results := runPipeline(ctx, n, req, opts)

	return func(yield func(Result) bool) {
		defer cancel()

		for result := range results {
			if !yield(result) {
				return
			}
		}
	}, nil
}

// Send sends an HTTP request and returns a performance [Result].
func Send(client *http.Client, req *http.Request) Result {
	var (
		t     = time.Now()
		code  int
		bytes int64
	)
	response, err := client.Do(req)
	if err == nil { // no error
		code = response.StatusCode
		bytes, err = io.Copy(io.Discard, response.Body)
		_ = response.Body.Close()
	}

	return Result{
		Duration: time.Since(t),
		Bytes:    bytes,
		Status:   code,
		Error:    err,
	}
}

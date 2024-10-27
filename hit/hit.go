package hit

import (
	"fmt"
	"net/http"
	"time"
)

// SendN sends n HTTP requests using the [Options] configuration and the
// provided [http.Request]. It returns a [Results] iterator for each [Result].
// The iterator stops when the consumer stops or when all requests are done.
func SendN(n int, req *http.Request, opts *Options) (Results, error) {
	opts = opts.clone().setDefaults()
	if n <= 0 {
		return nil, fmt.Errorf("invalid request count: %d", n)
	}
	// other condition checks are omitted for brevity

	return func(yield func(Result) bool) {
		for range n {
			if !yield(opts.Send(req)) {
				return
			}
		}
	}, nil
}

// Send sends an HTTP request and returns a performance [Result].
func Send(_ *http.Request) Result {
	const roundTripTime = 100 * time.Millisecond

	time.Sleep(roundTripTime)

	return Result{
		Status:   http.StatusOK,
		Bytes:    10,
		Duration: roundTripTime,
	}
}

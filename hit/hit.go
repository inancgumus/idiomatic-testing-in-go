package hit

import (
	"net/http"
	"time"
)

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

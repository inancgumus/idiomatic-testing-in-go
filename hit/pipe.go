package hit

import (
	"net/http"
	"time"
)

func runPipeline(n int, req *http.Request, opts *Options) <-chan Result {
	requests := produce(n, generateRequest(req))
	_ = requests

	return nil
}

func generateRequest(req *http.Request) requestFunc {
	return func() *http.Request {
		// NOTE: clone the request if you want to modify it.
		// req = req.Clone(req.Context())
		return req
	}
}

// requestFunc is a type of function that creates a request.
type requestFunc func() *http.Request

// produce produces n requests and sends them to out.
func produce(n int, genReq requestFunc) <-chan *http.Request {
	out := make(chan *http.Request)
	go func() {
		defer close(out)
		for range n {
			out <- genReq()
		}
	}()
	return out
}

// throttle throttles the incoming requests with the given delay.
func throttle(
	in <-chan *http.Request, delay time.Duration,
) <-chan *http.Request {
	out := make(chan *http.Request)
	go func() {
		defer close(out)

		t := time.NewTicker(delay)
		for r := range in {
			<-t.C
			out <- r
		}
	}()
	return out
}

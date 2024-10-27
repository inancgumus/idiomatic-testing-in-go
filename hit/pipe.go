package hit

import (
	"net/http"
	"sync"
	"time"
)

func runPipeline(n int, req *http.Request, opts *Options) <-chan Result {
	requests := produce(n, generateRequest(req))
	if opts.RPS > 0 {
		requests = throttle(requests, time.Second/time.Duration(opts.RPS))
	}

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

// dispatch sends requests from in using send and sends
// the results it receives to the returned channel.
// The concurrency parameter specifies the number of
// concurrent dispatcher workers to use.
func dispatch(
	in <-chan *http.Request,
	concurrency int,
	send SendFunc,
) <-chan Result {
	out := make(chan Result)

	var wg sync.WaitGroup
	wg.Add(concurrency)

	for range concurrency {
		go func() {
			defer wg.Done()
			dispatchRequest(in, out, send)
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// dispatchRequest receives requests from in and sends
// results to out.
func dispatchRequest(
	in <-chan *http.Request,
	out chan<- Result,
	send SendFunc,
) {
	for req := range in {
		out <- send(req)
	}
}

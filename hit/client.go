package hit

import (
	"context"
	"net/http"
	"runtime"
	"time"
)

// Client makes HTTP requests and returns an aggregated performance
// result. The fields should not be changed after initializing.
type Client struct {
	C         int               // C is the concurrency level (default: # of CPUs)
	RPS       int               // RPS throttles the requests per second
	Timeout   time.Duration     // Timeout per request
	Transport http.RoundTripper // Transport to use for the requests
}

// Do sends n HTTP requests and returns an aggregated [Result].
func (c *Client) Do(ctx context.Context, r *http.Request, n int) Result {
	t := time.Now()

	var sum Result
	for result := range c.do(ctx, r, n) {
		sum = sum.Merge(result)
	}

	return sum.Finalize(time.Since(t))
}

func (c *Client) do(ctx context.Context, r *http.Request, n int) <-chan Result {
	pipe := produce(ctx, n, func() *http.Request {
		return r.Clone(ctx)
	})

	if c.RPS > 0 {
		t := time.Second / time.Duration(c.RPS*c.concurrency())
		pipe = throttle(pipe, t)
	}

	client := c.client()
	defer client.CloseIdleConnections()

	return split(pipe, c.concurrency(), func(r *http.Request) Result {
		// skipping error handling because it's already
		// handled by the performance result summary.
		result, _ := Send(client, r)
		return result
	})
}

func (c *Client) client() *http.Client {
	transport := c.Transport
	if transport == nil {
		transport = &http.Transport{
			MaxIdleConnsPerHost: c.concurrency(),
		}
	}

	return &http.Client{
		Timeout: c.Timeout,
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: transport,
	}
}

func (c *Client) concurrency() int {
	if c.C > 0 {
		return c.C
	}
	return runtime.NumCPU()
}

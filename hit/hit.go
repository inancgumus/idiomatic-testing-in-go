// Package hit provides facilities for making concurrent requests
// to a server, aggregating and returning statistics.
//
// # Client Example
//
// [Client.Do] method can send many requests to a server.
//
// Example:
//
//	c := &hit.Client{
//		C:       10,        		// makes ten parallell requests (default is the number of CPUs).
//		RPS:     100,       		// throttles requests to one hundred per second.
//		Timeout: 10 * time.Second,	// each request will time after ten seconds.
//	}
//	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080", http.NoBody)
//	if err != nil {
//		log.Fatal(err)
//	}
//	sum := c.Do(context.Background(), request, 1_000_000)
//	sum.Fprint(os.Stdout)
//
// # Do Example
//
// [SendN] function is a convenience around [Client.Do].
//
// The following example sends one million requests to localhost:
//
//	sum, err := hit.SendN(context.Background(), "http://localhost:8080", 1_000_000)
//	if err != nil {
//		log.Fatal(err)
//	}
//	sum.Fprint(os.Stdout)
//
// You can customize [SendN] as follows:
//
//	sum, err := hit.SendN(
//		context.Background(), "http://localhost:8080", 1_000_000,
//		hit.Timeout(time.Second),
//		hit.Concurrency(10),
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
package hit

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Option allows changing [Client]'s behavior.
type Option func(*Client)

// Concurrency changes [Client]'s concurrency level.
func Concurrency(n int) Option {
	return func(c *Client) { c.C = n }
}

// RequestsPerSecond changes [Client]'s RPS (requests per second).
func RequestsPerSecond(d int) Option {
	return func(c *Client) { c.RPS = d }
}

// Timeout changes [Client]'s timeout per request.
func Timeout(d time.Duration) Option {
	return func(c *Client) { c.Timeout = d }
}

// SendN sends n HTTP requests to the url and returns an aggregated [Result].
func SendN(ctx context.Context, url string, n int, opts ...Option) (Result, error) {
	r, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return Result{}, fmt.Errorf("new http request: %w", err)
	}

	var c Client
	for _, o := range opts {
		o(&c)
	}

	return c.Do(ctx, r, n), nil
}

// Send sends an HTTP request and returns a performance [Result].
func Send(c *http.Client, r *http.Request) (Result, error) {
	t := time.Now()

	var (
		code  int
		bytes int64
	)

	response, err := c.Do(r)
	if err == nil { // no error
		code = response.StatusCode
		bytes, err = io.Copy(io.Discard, response.Body)
		_ = response.Body.Close()
	}

	result := Result{
		Duration: time.Since(t),
		Bytes:    bytes,
		Status:   code,
		Error:    err,
	}

	return result, err
}

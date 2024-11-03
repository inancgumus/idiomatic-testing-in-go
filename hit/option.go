package hit

import "net/http"

// SendFunc is a type of function that sends a request
// and returns a result.
type SendFunc func(*http.Request) Result

// Options defines the options for sending requests.
type Options struct {
	// Concurrency is the number of concurrent requests to send.
	// Default: 1
	Concurrency int

	// RPS is the requests to send per second.
	// Default: 0 (no rate limiting)
	RPS int

	// Send allows customizing the request processing.
	// Default: Uses [Send].
	Send SendFunc

	// Client is the HTTP client to use for sending requests.
	// Default: http.DefaultClient
	Client *http.Client
}

// DefaultOptions returns a new [Options] with the defaults.
func DefaultOptions() *Options {
	return new(Options).setDefaults()
}

// setDefaults sets the defaults.
// It returns a new [Options] if the current one is nil.
func (o *Options) setDefaults() *Options {
	if o == nil {
		o = new(Options)
	}
	if o.Concurrency <= 0 {
		o.Concurrency = 1
	}
	if o.Client == nil {
		o.Client = http.DefaultClient
	}
	if o.Send == nil {
		o.Send = func(r *http.Request) Result {
			return Send(o.Client, r)
		}
	}
	return o
}

// clone returns a shallow copy of the [Options].
// It returns nil if the current [Options] is nil.
func (o *Options) clone() *Options {
	if o == nil {
		return nil
	}
	c := *o
	return &c
}

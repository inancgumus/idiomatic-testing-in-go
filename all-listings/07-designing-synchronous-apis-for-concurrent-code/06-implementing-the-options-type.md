# Listing 7.6: Implementing the Options type

## [hit](https://github.com/inancgumus/gobyexample/blob/3fa192a4be25f366b93ef69761827cdcd54bda00/hit) / [option.go](https://github.com/inancgumus/gobyexample/blob/3fa192a4be25f366b93ef69761827cdcd54bda00/hit/option.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
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
	if o.Send == nil {
		o.Send = Send
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
```


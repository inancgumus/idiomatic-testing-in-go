package hit

import (
	"context"
	"net/http"
	"sync"
	"time"
)

// Produce calls fn n times and sends results to out.
func Produce(ctx context.Context, out chan<- *http.Request, n int, fn func() *http.Request) {
	for range n {
		select {
		case <-ctx.Done():
			return
		case out <- fn():
		}
	}
}

// produce runs [Produce] in a goroutine.
func produce(ctx context.Context, n int, fn func() *http.Request) <-chan *http.Request {
	out := make(chan *http.Request)
	go func() {
		defer close(out)
		Produce(ctx, out, n, fn)
	}()
	return out
}

// Throttle slows down receiving from in by delay and
// sends what it receives from in to out.
func Throttle(in <-chan *http.Request, out chan<- *http.Request, delay time.Duration) {
	t := time.NewTicker(delay)
	defer t.Stop()

	for r := range in {
		<-t.C
		out <- r
	}
}

// throttle runs [Throttle] in a goroutine.
func throttle(in <-chan *http.Request, delay time.Duration) <-chan *http.Request {
	out := make(chan *http.Request)
	go func() {
		defer close(out)
		Throttle(in, out, delay)
	}()
	return out
}

// SendFunc is the type of the function that sends an HTTP
// request and returns a performance result.
type SendFunc func(*http.Request) Result

// Split splits the pipeline into c goroutines, each running fn with
// what split receives from in, and sends results to out.
func Split(in <-chan *http.Request, out chan<- Result, c int, fn SendFunc) {
	send := func() {
		for r := range in {
			out <- fn(r)
		}
	}
	var wg sync.WaitGroup
	wg.Add(c)
	for range c {
		go func() {
			defer wg.Done()
			send()
		}()
	}
	wg.Wait()
}

// split runs [Split] in a goroutine.
func split(in <-chan *http.Request, c int, fn SendFunc) <-chan Result {
	out := make(chan Result)
	go func() {
		defer close(out)
		Split(in, out, c, fn)
	}()
	return out
}

// Alternative (not used in the book)

// SplitLimit is like [Split], but it uses a semaphore.
func SplitLimit(in <-chan *http.Request, out chan<- Result, c int, fn func(*http.Request) Result) {
	type token struct{}

	consume := func(lim chan token, r *http.Request) {
		// acquire the token (similar to wg.Add)
		lim <- token{}
		go func() {
			// release the token (similar to wg.Done)
			defer func() { <-lim }()
			out <- fn(r)
		}()
	}

	// limit the channel capacity to c. the channel blocks the
	// goroutine that wants to send to the channel if the channel's
	// buffer is full. allowing us to run C goroutines at a time.
	lim := make(chan token, c)

	// the go scheduler will ~evenly distribute the incoming values to
	// the goroutines.
	for r := range in {
		consume(lim, r)
	}
	// wait until all slots are released by the goroutines
	// (similar to wg.Wait)
	for range c {
		lim <- token{} // acquire the token
	}
}

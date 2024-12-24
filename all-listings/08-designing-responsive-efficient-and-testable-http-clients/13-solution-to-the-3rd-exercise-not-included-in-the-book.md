# Listing 8.13: Solution to the 3rd exercise (not included in the book)

## [hit](https://github.com/inancgumus/gobyexample/blob/67bba939e66ca618094390a11f77a6c1aa2c1d9f/hit) / [pipe.go](https://github.com/inancgumus/gobyexample/blob/67bba939e66ca618094390a11f77a6c1aa2c1d9f/hit/pipe.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package hit

import (
	"context"
	"net/http"
	"sync"
	"time"
)

func runPipeline(
	ctx context.Context, n int, req *http.Request, opts *Options,
) <-chan Result {
	requests := produce(ctx, n, generateRequest(req))
	if opts.RPS > 0 {
		requests = throttle(ctx, requests, time.Second/time.Duration(opts.RPS))
	}
	return dispatch(ctx, requests, opts.Concurrency, opts.Send)
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
func produce(
	ctx context.Context, n int, genReq requestFunc,
) <-chan *http.Request {
	out := make(chan *http.Request)

	go func() {
		defer close(out)
		for range n {
			select {
			case <-ctx.Done():
				return
			case out <- genReq():
			}
		}
	}()

	return out
}

// throttle throttles the incoming requests with the given delay.
func throttle(
	ctx context.Context, in <-chan *http.Request, delay time.Duration,
) <-chan *http.Request {
	out := make(chan *http.Request)
	go func() {
		defer close(out)

		t := time.NewTicker(delay)
		for r := range in {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
			}
			select {
			case <-ctx.Done():
				return
			case out <- r:
			}
		}
	}()
	return out
}

// dispatch sends requests from in using send and sends
// the results it receives to the returned channel.
// The concurrency parameter specifies the number of
// concurrent dispatcher workers to use.
func dispatch(
	ctx context.Context,
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
			dispatchRequest(ctx, in, out, send)
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
	ctx context.Context,
	in <-chan *http.Request,
	out chan<- Result,
	send SendFunc,
) {
	for req := range in {
		select {
		case out <- send(req):
		case <-ctx.Done():
			return
		}
	}
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -10,11 +10,11 @@ import (
 func runPipeline(
 	ctx context.Context, n int, req *http.Request, opts *Options,
 ) <-chan Result {
 	requests := produce(ctx, n, generateRequest(req))
 	if opts.RPS > 0 {
-		requests = throttle(requests, time.Second/time.Duration(opts.RPS))
+		requests = throttle(ctx, requests, time.Second/time.Duration(opts.RPS))
 	}
-	return dispatch(requests, opts.Concurrency, opts.Send)
+	return dispatch(ctx, requests, opts.Concurrency, opts.Send)
 }
 
 func generateRequest(req *http.Request) requestFunc {
@@ -50,58 +50,72 @@ func produce(
 
 // throttle throttles the incoming requests with the given delay.
 func throttle(
-	in <-chan *http.Request, delay time.Duration,
+	ctx context.Context, in <-chan *http.Request, delay time.Duration,
 ) <-chan *http.Request {
 	out := make(chan *http.Request)
 	go func() {
 		defer close(out)
 
 		t := time.NewTicker(delay)
 		for r := range in {
-			<-t.C
-			out <- r
+			select {
+			case <-ctx.Done():
+				return
+			case <-t.C:
+			}
+			select {
+			case <-ctx.Done():
+				return
+			case out <- r:
+			}
 		}
 	}()
 	return out
 }
 
 // dispatch sends requests from in using send and sends
 // the results it receives to the returned channel.
 // The concurrency parameter specifies the number of
 // concurrent dispatcher workers to use.
 func dispatch(
+	ctx context.Context,
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
-			dispatchRequest(in, out, send)
+			dispatchRequest(ctx, in, out, send)
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
+	ctx context.Context,
 	in <-chan *http.Request,
 	out chan<- Result,
 	send SendFunc,
 ) {
 	for req := range in {
-		out <- send(req)
+		select {
+		case out <- send(req):
+		case <-ctx.Done():
+			return
+		}
 	}
 }
```


# Listing 7.11: Implementing the throttler stage

## [hit](https://github.com/inancgumus/gobyexample/blob/e4086bc6739daaa5d56c8cff0b3a2bcfeda6ff1d/hit) / [pipe.go](https://github.com/inancgumus/gobyexample/blob/e4086bc6739daaa5d56c8cff0b3a2bcfeda6ff1d/hit/pipe.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
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
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -1,6 +1,9 @@
 package hit
 
-import "net/http"
+import (
+	"net/http"
+	"time"
+)
 
 func runPipeline(n int, req *http.Request, opts *Options) <-chan Result {
 	requests := produce(n, generateRequest(req))
@@ -31,3 +34,20 @@ func produce(n int, genReq requestFunc) <-chan *http.Request {
 	}()
 	return out
 }
+
+// throttle throttles the incoming requests with the given delay.
+func throttle(
+	in <-chan *http.Request, delay time.Duration,
+) <-chan *http.Request {
+	out := make(chan *http.Request)
+	go func() {
+		defer close(out)
+
+		t := time.NewTicker(delay)
+		for r := range in {
+			<-t.C
+			out <- r
+		}
+	}()
+	return out
+}
```


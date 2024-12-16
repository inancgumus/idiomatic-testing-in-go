# Listing 7.10: Integrating the producer stage

## [hit](https://github.com/inancgumus/gobyexample/blob/f80a95ea001ed03903b564a532e75b41738939ee/hit) / [pipe.go](https://github.com/inancgumus/gobyexample/blob/f80a95ea001ed03903b564a532e75b41738939ee/hit/pipe.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package hit

import "net/http"

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
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -2,6 +2,21 @@ package hit
 
 import "net/http"
 
+func runPipeline(n int, req *http.Request, opts *Options) <-chan Result {
+	requests := produce(n, generateRequest(req))
+	_ = requests
+
+	return nil
+}
+
+func generateRequest(req *http.Request) requestFunc {
+	return func() *http.Request {
+		// NOTE: clone the request if you want to modify it.
+		// req = req.Clone(req.Context())
+		return req
+	}
+}
+
 // requestFunc is a type of function that creates a request.
 type requestFunc func() *http.Request
 
```


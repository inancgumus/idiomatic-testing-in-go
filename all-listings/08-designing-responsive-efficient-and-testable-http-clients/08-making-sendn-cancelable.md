# Listing 8.8: Making SendN cancelable

## [hit](https://github.com/inancgumus/gobyexample/blob/40c435a0ef7aee095f22e796230f44c79252deb0/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/40c435a0ef7aee095f22e796230f44c79252deb0/hit/hit.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package hit

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// SendN sends n HTTP requests using the [Options] configuration and the
// provided [http.Request]. It returns a [Results] iterator for each [Result].
// The iterator stops when the consumer stops or when all requests are done.
func SendN(
	ctx context.Context, n int, req *http.Request, opts *Options,
) (Results, error) {
	opts = opts.clone().setDefaults()
	if n <= 0 {
		return nil, fmt.Errorf("invalid request count: %d", n)
	}

	ctx, cancel := context.WithCancel(ctx)
	results := runPipeline(ctx, n, req, opts)

	return func(yield func(Result) bool) {
		defer cancel()

		for result := range results {
			if !yield(result) {
				return
			}
		}
	}, nil
}

// Send sends an HTTP request and returns a performance [Result].
func Send(client *http.Client, req *http.Request) Result {
	var (
		t     = time.Now()
		code  int
		bytes int64
	)
	response, err := client.Do(req)
	if err == nil { // no error
		code = response.StatusCode
		bytes, err = io.Copy(io.Discard, response.Body)
		_ = response.Body.Close()
	}

	return Result{
		Duration: time.Since(t),
		Bytes:    bytes,
		Status:   code,
		Error:    err,
	}
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -1,30 +1,36 @@
 package hit
 
 import (
+	"context"
 	"fmt"
 	"io"
 	"net/http"
 	"time"
 )
 
 // SendN sends n HTTP requests using the [Options] configuration and the
 // provided [http.Request]. It returns a [Results] iterator for each [Result].
 // The iterator stops when the consumer stops or when all requests are done.
-func SendN(n int, req *http.Request, opts *Options) (Results, error) {
+func SendN(
+	ctx context.Context, n int, req *http.Request, opts *Options,
+) (Results, error) {
 	opts = opts.clone().setDefaults()
 	if n <= 0 {
 		return nil, fmt.Errorf("invalid request count: %d", n)
 	}
 
-	results := runPipeline(n, req, opts)
+	ctx, cancel := context.WithCancel(ctx)
+	results := runPipeline(ctx, n, req, opts)
 
 	return func(yield func(Result) bool) {
+		defer cancel()
+
 		for result := range results {
 			if !yield(result) {
 				return
 			}
 		}
 	}, nil
 }
 
 // Send sends an HTTP request and returns a performance [Result].
```


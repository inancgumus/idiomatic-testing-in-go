# Listing 8.3: Saving the status code and error in a Result

## [hit](https://github.com/inancgumus/gobyexample/blob/5b9dd1c442e095d72413a0ee9e66a29a23ebc644/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/5b9dd1c442e095d72413a0ee9e66a29a23ebc644/hit/hit.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package hit

import (
	"fmt"
	"net/http"
	"time"
)

// SendN sends n HTTP requests using the [Options] configuration and the
// provided [http.Request]. It returns a [Results] iterator for each [Result].
// The iterator stops when the consumer stops or when all requests are done.
func SendN(n int, req *http.Request, opts *Options) (Results, error) {
	opts = opts.clone().setDefaults()
	if n <= 0 {
		return nil, fmt.Errorf("invalid request count: %d", n)
	}

	results := runPipeline(n, req, opts)

	return func(yield func(Result) bool) {
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
		t    = time.Now()
		code int
	)
	response, err := client.Do(req)
	if err == nil { // no error
		code = response.StatusCode
	}

	return Result{
		Duration: time.Since(t),
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
@@ -28,10 +28,18 @@ func SendN(n int, req *http.Request, opts *Options) (Results, error) {
 
 // Send sends an HTTP request and returns a performance [Result].
 func Send(client *http.Client, req *http.Request) Result {
-	t := time.Now()
-
+	var (
+		t    = time.Now()
+		code int
+	)
 	response, err := client.Do(req)
-	_, _ = response, err
+	if err == nil { // no error
+		code = response.StatusCode
+	}
 
-	return Result{Duration: time.Since(t)}
+	return Result{
+		Duration: time.Since(t),
+		Status:   code,
+		Error:    err,
+	}
 }
```


# Listing 8.2: Sending a request using http.Client

## [hit](https://github.com/inancgumus/gobyexample/blob/2c2a678d9fa9965eb9e644b46c633f43d4ef7534/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/2c2a678d9fa9965eb9e644b46c633f43d4ef7534/hit/hit.go)

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
	t := time.Now()

	response, err := client.Do(req)
	_, _ = response, err

	return Result{Duration: time.Since(t)}
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -27,14 +27,11 @@ func SendN(n int, req *http.Request, opts *Options) (Results, error) {
 }
 
 // Send sends an HTTP request and returns a performance [Result].
-func Send(_ *http.Request) Result {
-	const roundTripTime = 100 * time.Millisecond
+func Send(client *http.Client, req *http.Request) Result {
+	t := time.Now()
 
-	time.Sleep(roundTripTime)
+	response, err := client.Do(req)
+	_, _ = response, err
 
-	return Result{
-		Status:   http.StatusOK,
-		Bytes:    10,
-		Duration: roundTripTime,
-	}
+	return Result{Duration: time.Since(t)}
 }
```


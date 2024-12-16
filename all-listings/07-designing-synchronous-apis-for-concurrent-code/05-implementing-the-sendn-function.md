# Listing 7.5: Implementing the SendN function

## [hit](https://github.com/inancgumus/gobyexample/blob/4c778e47538bdbb4a29fa6cb4f859554d55404f4/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/4c778e47538bdbb4a29fa6cb4f859554d55404f4/hit/hit.go)

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
	// other condition checks are omitted for brevity

	return func(yield func(Result) bool) {
		for range n {
			if !yield(opts.Send(req)) {
				return
			}
		}
	}, nil
}

// Send sends an HTTP request and returns a performance [Result].
func Send(_ *http.Request) Result {
	const roundTripTime = 100 * time.Millisecond

	time.Sleep(roundTripTime)

	return Result{
		Status:   http.StatusOK,
		Bytes:    10,
		Duration: roundTripTime,
	}
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -1,10 +1,30 @@
 package hit
 
 import (
+	"fmt"
 	"net/http"
 	"time"
 )
 
+// SendN sends n HTTP requests using the [Options] configuration and the
+// provided [http.Request]. It returns a [Results] iterator for each [Result].
+// The iterator stops when the consumer stops or when all requests are done.
+func SendN(n int, req *http.Request, opts *Options) (Results, error) {
+	opts = opts.clone().setDefaults()
+	if n <= 0 {
+		return nil, fmt.Errorf("invalid request count: %d", n)
+	}
+	// other condition checks are omitted for brevity
+
+	return func(yield func(Result) bool) {
+		for range n {
+			if !yield(opts.Send(req)) {
+				return
+			}
+		}
+	}, nil
+}
+
 // Send sends an HTTP request and returns a performance [Result].
 func Send(_ *http.Request) Result {
 	const roundTripTime = 100 * time.Millisecond
```


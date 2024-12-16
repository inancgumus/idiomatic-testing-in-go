# Listing 7.4: Implementing the Send function

## [hit](https://github.com/inancgumus/gobyexample/blob/870185038cb6209e74b53b84e465ea21bee0ba3a/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/870185038cb6209e74b53b84e465ea21bee0ba3a/hit/hit.go)

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


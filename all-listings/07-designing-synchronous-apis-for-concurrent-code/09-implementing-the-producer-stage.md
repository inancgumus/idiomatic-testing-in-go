# Listing 7.9: Implementing the producer stage

## [hit](https://github.com/inancgumus/gobyexample/blob/72a56eaa1d7fba2c34d00e99bce568097075228f/hit) / [pipe.go](https://github.com/inancgumus/gobyexample/blob/72a56eaa1d7fba2c34d00e99bce568097075228f/hit/pipe.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package hit

import "net/http"

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


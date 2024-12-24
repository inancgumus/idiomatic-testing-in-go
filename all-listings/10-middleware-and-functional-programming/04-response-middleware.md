# Listing 10.4: Response middleware

## [bite](https://github.com/inancgumus/gobyexample/blob/25c1648cd622caacb1f1a3082298f9f0bab981ad/bite) / [httplog](https://github.com/inancgumus/gobyexample/blob/25c1648cd622caacb1f1a3082298f9f0bab981ad/bite/httplog) / [response.go](https://github.com/inancgumus/gobyexample/blob/25c1648cd622caacb1f1a3082298f9f0bab981ad/bite/httplog/response.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package httplog

import (
	"log/slog"
	"net/http"
	"time"
)

// Response provides HTTP-response logging.
// Its zero value is useful and ready to use.
type Response struct {
	requestStart time.Time
}

// Wrap is a middleware that records response related data.
func (res *Response) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res.onRequestStart(r)
		next.ServeHTTP(w, r)
	})
}

func (res *Response) onRequestStart(_ *http.Request) {
	res.requestStart = time.Now()
}

// Time returns an [slog.Attr] that contains the response time.
func (res *Response) Time(_ *http.Request) slog.Attr {
	if res.requestStart.IsZero() {
		return slog.Attr{}
	}
	return slog.Duration("response_time", time.Since(res.requestStart))
}
```


# Listing 10.3: Built-in attribute functions

## [bite](https://github.com/inancgumus/gobyexample/blob/e20a38a6aff785a6972d2466038706e26f7aeb91/bite) / [httplog](https://github.com/inancgumus/gobyexample/blob/e20a38a6aff785a6972d2466038706e26f7aeb91/bite/httplog) / [request.go](https://github.com/inancgumus/gobyexample/blob/e20a38a6aff785a6972d2466038706e26f7aeb91/bite/httplog/request.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package httplog

import (
	"log/slog"
	"net/http"
)

// URL returns an [slog.Attr] that contains the URL.
func URL(r *http.Request) slog.Attr {
	return slog.Any("url", r.URL)
}

// Method returns an [slog.Attr] that contains the method.
func Method(r *http.Request) slog.Attr {
	return slog.String("method", r.Method)
}

// RemoteAddr returns an [slog.Attr] that contains the remote address.
func RemoteAddr(r *http.Request) slog.Attr {
	return slog.String("remote_addr", r.RemoteAddr)
}

// Header returns an [AttrFunc] that returns an [slog.Attr] with the key.
func Header(key string) AttrFunc {
	return func(r *http.Request) slog.Attr {
		if v := r.Header.Get(key); v != "" {
			// Skipping "key" formatting for the sake of simplicity of the example
			// (e.g., user_agent instead of User-Agent).
			return slog.String(key, r.Header.Get(key))
		}
		return slog.Attr{}
	}
}
```


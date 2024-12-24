# Listing 10.1: Logger middleware

## [bite](https://github.com/inancgumus/gobyexample/blob/06df96ac9c42a4e429c6fa69841f8e0f7e5852fe/bite) / [httplog](https://github.com/inancgumus/gobyexample/blob/06df96ac9c42a4e429c6fa69841f8e0f7e5852fe/bite/httplog) / [httplog.go](https://github.com/inancgumus/gobyexample/blob/06df96ac9c42a4e429c6fa69841f8e0f7e5852fe/bite/httplog/httplog.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
// Package httplog provides HTTP request and response logging.
package httplog

import (
	"log/slog"
	"net/http"
)

// Logger logs incoming requests and responses.
type Logger struct {
	lg *slog.Logger
}

// New returns a new [Logger].
func New(lg *slog.Logger) *Logger {
	return &Logger{lg: lg}
}

// Wrap is a middleware that logs requests and responses.
func (lg *Logger) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		lg.lg.LogAttrs(r.Context(), slog.LevelInfo, "httplog",
			slog.String("url", r.URL.String()),
		)
	})
}
```


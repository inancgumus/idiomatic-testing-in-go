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

// Package httplog provides HTTP request and response logging.
package httplog

import (
	"log/slog"
	"net/http"
)

// AttrFunc returns an [slog.Attr] to log.
type AttrFunc func(*http.Request) slog.Attr

// Logger logs incoming requests and responses.
type Logger struct {
	lg        *slog.Logger
	attrFuncs []AttrFunc
}

// New returns a new [Logger].
func New(lg *slog.Logger) *Logger {
	return &Logger{lg: lg}
}

// With augments logs with additional attributes.
func (lg *Logger) With(funcs ...AttrFunc) *Logger {
	lg.attrFuncs = append(lg.attrFuncs, funcs...)
	return lg
}

// Wrap is a middleware that logs requests and responses.
func (lg *Logger) Wrap(next http.Handler) http.Handler {
	attrs := make([]slog.Attr, len(lg.attrFuncs))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		for i, f := range lg.attrFuncs {
			attrs[i] = f(r)
		}
		lg.lg.LogAttrs(r.Context(), slog.LevelInfo, "httplog", attrs...)
	})
}

package httpio

import (
	"context"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"time"
)

// LoggingMiddleware logs requests.
func LoggingMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rr := &responseRecorder{
			ResponseWriter: w,
		}
		next.ServeHTTP(rr, r)

		slog.Log(r.Context(), slog.LevelInfo,
			"request",
			"url", r.URL, "method", r.Method,
			"status", rr.status,
			"took", time.Since(start))
	}
}

type responseRecorder struct {
	http.ResponseWriter
	status int
}

func (r *responseRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}

func (r *responseRecorder) Unwrap() http.ResponseWriter {
	return r.ResponseWriter
}

type traceIDKey struct{}

// TraceMiddleware adds a trace ID to the context.
func TraceMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := SetTraceID(r.Context(), rand.Int64())
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// SetTraceID sets the trace ID to the context.
func SetTraceID(ctx context.Context, id int64) context.Context {
	return context.WithValue(ctx, traceIDKey{}, id)
}

// TraceID returns the trace ID from the context.
func TraceID(ctx context.Context) (int64, bool) {
	n, ok := ctx.Value(traceIDKey{}).(int64)
	return n, ok
}

// LogHandler adds the trace ID to the log records if
// the context has a trace ID. Use [TraceMiddleware] to set
// the trace ID.
type LogHandler struct{ slog.Handler }

func (h *LogHandler) Handle(ctx context.Context, r slog.Record) error {
	if id, ok := TraceID(ctx); ok {
		r.Add("trace_id", id)
	}
	return h.Handler.Handle(ctx, r)
}

func (h *LogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &LogHandler{Handler: h.Handler.WithAttrs(attrs)}
}

func (h *LogHandler) WithGroup(name string) slog.Handler {
	return &LogHandler{Handler: h.Handler.WithGroup(name)}
}

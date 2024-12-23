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

// responseInterceptor intercepts [http.ResponseWriter] calls.
type responseInterceptor struct {
	http.ResponseWriter
	writeHeader func(int)
}

func (r *responseInterceptor) WriteHeader(code int) {
	if r.writeHeader != nil {
		r.writeHeader(code)
	}
	r.ResponseWriter.WriteHeader(code)
}

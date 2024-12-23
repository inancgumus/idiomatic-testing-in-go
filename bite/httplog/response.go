package httplog

import (
	"log/slog"
	"net/http"
	"time"
)

// Response provides HTTP-response logging.
// Its zero value is useful and ready to use.
type Response struct {
	statusCode   int
	requestStart time.Time
}

// Wrap is a middleware that records response related data.
func (res *Response) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res.onRequestStart(r)
		next.ServeHTTP(&responseInterceptor{
			ResponseWriter: w,
			writeHeader:    res.onWriteHeader,
		}, r)
	})
}

// onRequestStart is called when a new request starts.
func (res *Response) onRequestStart(_ *http.Request) {
	res.statusCode = http.StatusOK // default status code
	res.requestStart = time.Now()
}

// onWriteHeader is called when the response status code is written.
func (res *Response) onWriteHeader(code int) {
	res.statusCode = code
}

// Time returns an [slog.Attr] that contains the response time.
func (res *Response) Time(_ *http.Request) slog.Attr {
	if res.requestStart.IsZero() {
		return slog.Attr{}
	}
	return slog.Duration("response_time", time.Since(res.requestStart))
}

// StatusCode returns an [slog.Attr] that contains the HTTP status code.
func (res *Response) StatusCode(_ *http.Request) slog.Attr {
	return slog.Int("status", res.statusCode)
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

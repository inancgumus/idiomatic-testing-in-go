# Listing 10.7: Injecting a ResponseWriter

## [bite](https://github.com/inancgumus/gobyexample/blob/333f9f530a5311bf73e9e77d4ed69f648a4cb768/bite) / [httplog](https://github.com/inancgumus/gobyexample/blob/333f9f530a5311bf73e9e77d4ed69f648a4cb768/bite/httplog) / [response.go](https://github.com/inancgumus/gobyexample/blob/333f9f530a5311bf73e9e77d4ed69f648a4cb768/bite/httplog/response.go)

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
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -9,29 +9,45 @@ import (
 // Response provides HTTP-response logging.
 // Its zero value is useful and ready to use.
 type Response struct {
+	statusCode   int
 	requestStart time.Time
 }
 
 // Wrap is a middleware that records response related data.
 func (res *Response) Wrap(next http.Handler) http.Handler {
 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
 		res.onRequestStart(r)
-		next.ServeHTTP(w, r)
+		next.ServeHTTP(&responseInterceptor{
+			ResponseWriter: w,
+			writeHeader:    res.onWriteHeader,
+		}, r)
 	})
 }
 
+// onRequestStart is called when a new request starts.
 func (res *Response) onRequestStart(_ *http.Request) {
+	res.statusCode = http.StatusOK // default status code
 	res.requestStart = time.Now()
 }
 
+// onWriteHeader is called when the response status code is written.
+func (res *Response) onWriteHeader(code int) {
+	res.statusCode = code
+}
+
 // Time returns an [slog.Attr] that contains the response time.
 func (res *Response) Time(_ *http.Request) slog.Attr {
 	if res.requestStart.IsZero() {
 		return slog.Attr{}
 	}
 	return slog.Duration("response_time", time.Since(res.requestStart))
 }
 
+// StatusCode returns an [slog.Attr] that contains the HTTP status code.
+func (res *Response) StatusCode(_ *http.Request) slog.Attr {
+	return slog.Int("status", res.statusCode)
+}
+
 // responseInterceptor intercepts [http.ResponseWriter] calls.
 type responseInterceptor struct {
 	http.ResponseWriter
```


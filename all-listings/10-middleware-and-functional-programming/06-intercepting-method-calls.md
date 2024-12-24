# Listing 10.6: Intercepting method calls

## [bite](https://github.com/inancgumus/gobyexample/blob/e4e71e2934f4cc1ed4958ee126cfe732fb1c6138/bite) / [httplog](https://github.com/inancgumus/gobyexample/blob/e4e71e2934f4cc1ed4958ee126cfe732fb1c6138/bite/httplog) / [response.go](https://github.com/inancgumus/gobyexample/blob/e4e71e2934f4cc1ed4958ee126cfe732fb1c6138/bite/httplog/response.go)

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
@@ -31,3 +31,16 @@ func (res *Response) Time(_ *http.Request) slog.Attr {
 	}
 	return slog.Duration("response_time", time.Since(res.requestStart))
 }
+
+// responseInterceptor intercepts [http.ResponseWriter] calls.
+type responseInterceptor struct {
+	http.ResponseWriter
+	writeHeader func(int)
+}
+
+func (r *responseInterceptor) WriteHeader(code int) {
+	if r.writeHeader != nil {
+		r.writeHeader(code)
+	}
+	r.ResponseWriter.WriteHeader(code)
+}
```


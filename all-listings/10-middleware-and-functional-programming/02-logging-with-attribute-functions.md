# Listing 10.2: Logging with attribute functions

## [bite](https://github.com/inancgumus/gobyexample/blob/709c4b0ee7f6ee2299d46cc2964bc39c012f4088/bite) / [httplog](https://github.com/inancgumus/gobyexample/blob/709c4b0ee7f6ee2299d46cc2964bc39c012f4088/bite/httplog) / [httplog.go](https://github.com/inancgumus/gobyexample/blob/709c4b0ee7f6ee2299d46cc2964bc39c012f4088/bite/httplog/httplog.go)

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
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -4,24 +4,36 @@ package httplog
 import (
 	"log/slog"
 	"net/http"
 )
 
+// AttrFunc returns an [slog.Attr] to log.
+type AttrFunc func(*http.Request) slog.Attr
+
 // Logger logs incoming requests and responses.
 type Logger struct {
-	lg *slog.Logger
+	lg        *slog.Logger
+	attrFuncs []AttrFunc
 }
 
 // New returns a new [Logger].
 func New(lg *slog.Logger) *Logger {
 	return &Logger{lg: lg}
 }
 
+// With augments logs with additional attributes.
+func (lg *Logger) With(funcs ...AttrFunc) *Logger {
+	lg.attrFuncs = append(lg.attrFuncs, funcs...)
+	return lg
+}
+
 // Wrap is a middleware that logs requests and responses.
 func (lg *Logger) Wrap(next http.Handler) http.Handler {
+	attrs := make([]slog.Attr, len(lg.attrFuncs))
 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
 		next.ServeHTTP(w, r)
-		lg.lg.LogAttrs(r.Context(), slog.LevelInfo, "httplog",
-			slog.String("url", r.URL.String()),
-		)
+		for i, f := range lg.attrFuncs {
+			attrs[i] = f(r)
+		}
+		lg.lg.LogAttrs(r.Context(), slog.LevelInfo, "httplog", attrs...)
 	})
 }
```


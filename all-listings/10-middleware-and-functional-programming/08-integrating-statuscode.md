# Listing 10.8: Integrating StatusCode

## [bite](https://github.com/inancgumus/gobyexample/blob/a47d2b44fbbb2609c638be598227fb6093bad84b/bite) / [cmd](https://github.com/inancgumus/gobyexample/blob/a47d2b44fbbb2609c638be598227fb6093bad84b/bite/cmd) / [linkd](https://github.com/inancgumus/gobyexample/blob/a47d2b44fbbb2609c638be598227fb6093bad84b/bite/cmd/linkd) / [linkd.go](https://github.com/inancgumus/gobyexample/blob/a47d2b44fbbb2609c638be598227fb6093bad84b/bite/cmd/linkd/linkd.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

Integration `Response.StatusCode`.

```go
func main() {
    fmt.Println("bok")
```

Hello.

 # Hello

 ## Hello2

```go
package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/inancgumus/gobyexample/bite/httplog"
	"github.com/inancgumus/gobyexample/bite/link"
)

func main() {
	const addr = "localhost:8080"

	log := slog.New(slog.NewTextHandler(os.Stderr, nil)).With("app", "linkd")
	log.Info("starting", "addr", addr)

	linkServer := link.NewServer(log, new(link.Store))

	response := new(httplog.Response)
	logger := httplog.New(log).With(
		httplog.URL,
		httplog.Method,
		httplog.RemoteAddr,
		response.Time,
		response.StatusCode,
	)

	srv := &http.Server{
		Addr:        addr,
		Handler:     logger.Wrap(response.Wrap(linkServer)),
		ReadTimeout: 20 * time.Second,
		IdleTimeout: 40 * time.Second,
	}
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Error("server closed unexpectedly", "message", err)
	}
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -14,26 +14,27 @@ import (
 func main() {
 	const addr = "localhost:8080"
 
 	log := slog.New(slog.NewTextHandler(os.Stderr, nil)).With("app", "linkd")
 	log.Info("starting", "addr", addr)
 
 	linkServer := link.NewServer(log, new(link.Store))
 
 	response := new(httplog.Response)
 	logger := httplog.New(log).With(
 		httplog.URL,
 		httplog.Method,
 		httplog.RemoteAddr,
 		response.Time,
+		response.StatusCode,
 	)
 
 	srv := &http.Server{
 		Addr:        addr,
 		Handler:     logger.Wrap(response.Wrap(linkServer)),
 		ReadTimeout: 20 * time.Second,
 		IdleTimeout: 40 * time.Second,
 	}
 	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
 		log.Error("server closed unexpectedly", "message", err)
 	}
 }
```


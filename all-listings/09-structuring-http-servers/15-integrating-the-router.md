# Listing 9.15: Integrating the router

## [bite](https://github.com/inancgumus/gobyexample/blob/cc8c61dd2440833b8ce3e6ebb3b3bf558d408b9a/bite) / [cmd](https://github.com/inancgumus/gobyexample/blob/cc8c61dd2440833b8ce3e6ebb3b3bf558d408b9a/bite/cmd) / [linkd](https://github.com/inancgumus/gobyexample/blob/cc8c61dd2440833b8ce3e6ebb3b3bf558d408b9a/bite/cmd/linkd) / [linkd.go](https://github.com/inancgumus/gobyexample/blob/cc8c61dd2440833b8ce3e6ebb3b3bf558d408b9a/bite/cmd/linkd/linkd.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/inancgumus/gobyexample/bite/link"
)

func main() {
	const addr = "localhost:8080"

	log := slog.New(slog.NewTextHandler(os.Stderr, nil)).With("app", "linkd")
	log.Info("starting", "addr", addr)

	linkServer := link.NewServer(log, new(link.Store))

	srv := &http.Server{
		Addr:        addr,
		Handler:     linkServer,
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
@@ -13,18 +13,18 @@ import (
 func main() {
 	const addr = "localhost:8080"
 
 	log := slog.New(slog.NewTextHandler(os.Stderr, nil)).With("app", "linkd")
 	log.Info("starting", "addr", addr)
 
-	linkServer := link.NewServer(log)
+	linkServer := link.NewServer(log, new(link.Store))
 
 	srv := &http.Server{
 		Addr:        addr,
-		Handler:     http.HandlerFunc(linkServer.Health),
+		Handler:     linkServer,
 		ReadTimeout: 20 * time.Second,
 		IdleTimeout: 40 * time.Second,
 	}
 	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
 		log.Error("server closed unexpectedly", "message", err)
 	}
 }
```


# Listing 9.9: Running an HTTP server

## [bite](https://github.com/inancgumus/gobyexample/blob/8c1853b33ca0e77f3098d6927d537b6856652c47/bite) / [cmd](https://github.com/inancgumus/gobyexample/blob/8c1853b33ca0e77f3098d6927d537b6856652c47/bite/cmd) / [linkd](https://github.com/inancgumus/gobyexample/blob/8c1853b33ca0e77f3098d6927d537b6856652c47/bite/cmd/linkd) / [linkd.go](https://github.com/inancgumus/gobyexample/blob/8c1853b33ca0e77f3098d6927d537b6856652c47/bite/cmd/linkd/linkd.go)

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

	"github.com/inancgumus/gobyexample/bite/link"
)

func main() {
	const addr = "localhost:8080"

	log := slog.New(slog.NewTextHandler(os.Stderr, nil)).With("app", "linkd")
	log.Info("starting", "addr", addr)

	linkServer := link.NewServer(log)

	err := http.ListenAndServe(addr, http.HandlerFunc(linkServer.Health))
	if !errors.Is(err, http.ErrServerClosed) {
		log.Error("server closed unexpectedly", "message", err)
	}
}
```

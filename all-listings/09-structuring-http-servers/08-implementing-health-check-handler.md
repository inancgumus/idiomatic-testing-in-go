# Listing 9.8: Implementing health check handler

## [bite](https://github.com/inancgumus/gobyexample/blob/bd98dccb80e5571ed235a60b1f6cd7c5ec85893c/bite) / [link](https://github.com/inancgumus/gobyexample/blob/bd98dccb80e5571ed235a60b1f6cd7c5ec85893c/bite/link) / [server.go](https://github.com/inancgumus/gobyexample/blob/bd98dccb80e5571ed235a60b1f6cd7c5ec85893c/bite/link/server.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package link

import (
	"fmt"
	"log/slog"
	"net/http"
)

// Server is a URL shortener HTTP server.
type Server struct {
	lg *slog.Logger
}

// NewServer returns a new [Server].
func NewServer(lg *slog.Logger) *Server {
	return &Server{lg: lg}
}

// Health serves the health check requests.
func (srv *Server) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}
```


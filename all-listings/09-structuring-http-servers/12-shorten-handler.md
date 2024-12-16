# Listing 9.12: Shorten handler

## [bite](https://github.com/inancgumus/gobyexample/blob/27bff3e8f6c4d1a0cfdc4a82f1398cc352e9a1b9/bite) / [link](https://github.com/inancgumus/gobyexample/blob/27bff3e8f6c4d1a0cfdc4a82f1398cc352e9a1b9/bite/link) / [server.go](https://github.com/inancgumus/gobyexample/blob/27bff3e8f6c4d1a0cfdc4a82f1398cc352e9a1b9/bite/link/server.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package link

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/inancgumus/gobyexample/bite"
)

// Server is a URL shortener HTTP server.
type Server struct {
	lg *slog.Logger
}

// NewServer returns a new [Server].
func NewServer(lg *slog.Logger) *Server {
	return &Server{lg: lg}
}

// Shorten handles the URL shortening requests.
func (srv *Server) Shorten(links *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		link := Link{
			Key: r.PostFormValue("key"),
			URL: r.PostFormValue("url"),
		}
		if err := links.Create(r.Context(), link); err != nil {
			httpError(w, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, link.Key)
	}
}

// Health serves the health check requests.
func (srv *Server) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func httpError(w http.ResponseWriter, err error) {
	code := http.StatusInternalServerError
	switch {
	case errors.Is(err, bite.ErrInvalidRequest):
		code = http.StatusBadRequest
	case errors.Is(err, bite.ErrExists):
		code = http.StatusConflict
	case errors.Is(err, bite.ErrNotExist):
		code = http.StatusNotFound
	}
	http.Error(w, err.Error(), code)
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -19,6 +19,22 @@ func NewServer(lg *slog.Logger) *Server {
 	return &Server{lg: lg}
 }
 
+// Shorten handles the URL shortening requests.
+func (srv *Server) Shorten(links *Store) http.HandlerFunc {
+	return func(w http.ResponseWriter, r *http.Request) {
+		link := Link{
+			Key: r.PostFormValue("key"),
+			URL: r.PostFormValue("url"),
+		}
+		if err := links.Create(r.Context(), link); err != nil {
+			httpError(w, err)
+			return
+		}
+		w.WriteHeader(http.StatusCreated)
+		fmt.Fprint(w, link.Key)
+	}
+}
+
 // Health serves the health check requests.
 func (srv *Server) Health(w http.ResponseWriter, r *http.Request) {
 	fmt.Fprintln(w, "OK")
```


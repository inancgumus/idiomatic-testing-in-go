# Listing 9.14: Routing with ServeMux

## [bite](https://github.com/inancgumus/gobyexample/blob/4299571cded8d7028491ddc6e4d7881a57638912/bite) / [link](https://github.com/inancgumus/gobyexample/blob/4299571cded8d7028491ddc6e4d7881a57638912/bite/link) / [server.go](https://github.com/inancgumus/gobyexample/blob/4299571cded8d7028491ddc6e4d7881a57638912/bite/link/server.go)

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
	lg  *slog.Logger
	mux *http.ServeMux
}

// NewServer returns a new [Server].
func NewServer(lg *slog.Logger, store *Store) *Server {
	mux := http.NewServeMux()
	srv := &Server{
		lg:  lg,
		mux: mux,
	}
	mux.HandleFunc("POST /shorten", srv.Shorten(store))
	mux.HandleFunc("GET /r/{key}", srv.Resolve(store))
	mux.HandleFunc("GET /health", srv.Health)
	return srv
}

func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.mux.ServeHTTP(w, r)
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

// Resolve handles the URL resolving requests for the shortened links.
func (srv *Server) Resolve(links *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		link, err := links.Retrieve(r.Context(), r.PathValue("key"))
		if err != nil {
			httpError(w, err)
			return
		}
		http.Redirect(w, r, link.URL, http.StatusFound)
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
@@ -11,12 +11,25 @@ import (
 
 // Server is a URL shortener HTTP server.
 type Server struct {
-	lg *slog.Logger
+	lg  *slog.Logger
+	mux *http.ServeMux
 }
 
 // NewServer returns a new [Server].
-func NewServer(lg *slog.Logger) *Server {
-	return &Server{lg: lg}
+func NewServer(lg *slog.Logger, store *Store) *Server {
+	mux := http.NewServeMux()
+	srv := &Server{
+		lg:  lg,
+		mux: mux,
+	}
+	mux.HandleFunc("POST /shorten", srv.Shorten(store))
+	mux.HandleFunc("GET /r/{key}", srv.Resolve(store))
+	mux.HandleFunc("GET /health", srv.Health)
+	return srv
+}
+
+func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
+	srv.mux.ServeHTTP(w, r)
 }
 
 // Shorten handles the URL shortening requests.
```


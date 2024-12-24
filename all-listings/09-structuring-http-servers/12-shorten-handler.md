# Listing 9.12: Shorten handler

## [bite](https://github.com/inancgumus/gobyexample/blob/67b6d06d143571e8e522e9c2d317e9cba397a0ba/bite) / [link](https://github.com/inancgumus/gobyexample/blob/67b6d06d143571e8e522e9c2d317e9cba397a0ba/bite/link) / [server.go](https://github.com/inancgumus/gobyexample/blob/67b6d06d143571e8e522e9c2d317e9cba397a0ba/bite/link/server.go)

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
//
//	Status Code       Condition
//	201               The link is successfully shortened.
//	400               The request is invalid.
//	409               The link already exists.
//	405               The request method is not POST.
//	413               The request body is too large.
//	500               There is an internal error.
func (srv *Server) Shorten(links *Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lnk := Link{
			Key: LinkKey(r.PostFormValue("key")),
			URL: r.PostFormValue("url"),
		}
		if err := lnk.Validate(); err != nil {
			httpError(w, fmt.Errorf("validating link: %w: %w",
				err, bite.ErrInvalidRequest))
			return
		}
		if err := links.Create(r.Context(), lnk); err != nil {
			httpError(w, fmt.Errorf("creating link: %w", err))
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, lnk.Key)
	})
}

// Health serves the health check requests.
//
//	Status Code       Condition
//	200               The server is healthy.
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
@@ -17,9 +17,38 @@ type Server struct {
 // NewServer returns a new [Server].
 func NewServer(lg *slog.Logger) *Server {
 	return &Server{lg: lg}
 }
 
+// Shorten handles the URL shortening requests.
+//
+//	Status Code       Condition
+//	201               The link is successfully shortened.
+//	400               The request is invalid.
+//	409               The link already exists.
+//	405               The request method is not POST.
+//	413               The request body is too large.
+//	500               There is an internal error.
+func (srv *Server) Shorten(links *Store) http.Handler {
+	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
+		lnk := Link{
+			Key: LinkKey(r.PostFormValue("key")),
+			URL: r.PostFormValue("url"),
+		}
+		if err := lnk.Validate(); err != nil {
+			httpError(w, fmt.Errorf("validating link: %w: %w",
+				err, bite.ErrInvalidRequest))
+			return
+		}
+		if err := links.Create(r.Context(), lnk); err != nil {
+			httpError(w, fmt.Errorf("creating link: %w", err))
+			return
+		}
+		w.WriteHeader(http.StatusCreated)
+		fmt.Fprint(w, lnk.Key)
+	})
+}
+
 // Health serves the health check requests.
 //
 //	Status Code       Condition
 //	200               The server is healthy.
```


# Listing 9.11: Converting errors to status codes

## [bite](https://github.com/inancgumus/gobyexample/blob/cbf33db4709cd689a3459e243c239f9a7e46a597/bite) / [link](https://github.com/inancgumus/gobyexample/blob/cbf33db4709cd689a3459e243c239f9a7e46a597/bite/link) / [server.go](https://github.com/inancgumus/gobyexample/blob/cbf33db4709cd689a3459e243c239f9a7e46a597/bite/link/server.go)

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
@@ -1,9 +1,12 @@
 package link
 
 import (
+	"errors"
 	"fmt"
 	"log/slog"
 	"net/http"
+
+	"github.com/inancgumus/gobyexample/bite"
 )
 
 // Server is a URL shortener HTTP server.
@@ -23,3 +26,16 @@ func NewServer(lg *slog.Logger) *Server {
 func (srv *Server) Health(w http.ResponseWriter, r *http.Request) {
 	fmt.Fprintln(w, "OK")
 }
+
+func httpError(w http.ResponseWriter, err error) {
+	code := http.StatusInternalServerError
+	switch {
+	case errors.Is(err, bite.ErrInvalidRequest):
+		code = http.StatusBadRequest
+	case errors.Is(err, bite.ErrExists):
+		code = http.StatusConflict
+	case errors.Is(err, bite.ErrNotExist):
+		code = http.StatusNotFound
+	}
+	http.Error(w, err.Error(), code)
+}
```


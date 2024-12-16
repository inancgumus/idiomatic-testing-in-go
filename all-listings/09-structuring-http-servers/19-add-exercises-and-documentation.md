# Listing 9.19: Add exercises and documentation

## [bite](https://github.com/inancgumus/gobyexample/blob/366ab8018e0fdf51963209e567578174a54d3145/bite) / [link](https://github.com/inancgumus/gobyexample/blob/366ab8018e0fdf51963209e567578174a54d3145/bite/link) / [client.go](https://github.com/inancgumus/gobyexample/blob/366ab8018e0fdf51963209e567578174a54d3145/bite/link/client.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package link

import "net/http"

// Client is a URL shortener client.
//
// # Exercise #1
//
// Implement an HTTP client for the URL shortener server.
//
//   - Implement the client for the all endpoints of [Server].
//   - The client should be able to shorten and resolve URLs
//     by sending HTTP requests to the server.
//
// Tip: You can see the client's proposed API in this file.
//
//   - Bonus: Add a timeout option and support.
//   - Bonus: Add a retry mechanism.
//
// Warning:
//
// Avoid using the server's HTTP handler functions in the
// client. These exercises are about implementing a client
// for the server using the standard library's net/http
// package.
//
// # Exercise #2
//
// Create a program called shortlink that uses [Client]
// with the following command-line interface:
//
//   - shortlink -s <url> # shortens the given URL
//   - shortlink -r <key> # resolves the short URL
//   - shortlink health   # checks the health of the server
//
// Optional flags:
//
//   - -addr: The address of the server. Default: localhost:8080
//   - -timeout: The timeout for the client. Default: 10 seconds
//   - -retry: The number of retries for the client. Default: 3
//   - -retry-wait: The wait time between retries. Default: 1 second
//   - -log: The log level. Default: info
type Client struct{ client *http.Client } //nolint:unused

// NewClient creates and returns a new [Client].
func NewClient() *Client { return nil }

// Shorten shortens the given URL and returns a [Link].
func (c *Client) Shorten(url string) (Link, error) { return Link{}, nil }

// Resolve resolves the given short URL and returns a [Link].
func (c *Client) Resolve(key string) (Link, error) { return Link{}, nil }

// Health checks the health of the server.
func (c *Client) Health() error { return nil }
```

## [bite](https://github.com/inancgumus/gobyexample/blob/366ab8018e0fdf51963209e567578174a54d3145/bite) / [link](https://github.com/inancgumus/gobyexample/blob/366ab8018e0fdf51963209e567578174a54d3145/bite/link) / [doc.go](https://github.com/inancgumus/gobyexample/blob/366ab8018e0fdf51963209e567578174a54d3145/bite/link/doc.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
// Package link provides a link management [Server] and its [Client].
//
//   - The [Server] shortens URLs and resolves shortened URLs.
//   - The [Client] can shorten and resolve URLs. However, it is not
//     implemented yet and is left as an exercise.
//
// # Endpoints
//
// The service provides three handlers:
//   - [Shorten] - Shortens a URL.
//   - [Resolve] - Resolves a shortened URL.
//   - [Health] - Checks the health of the service.
//
// # Curl examples
//
// Shorten a URL:
//
//	$ curl localhost:8080/shorten -d '{"key":"inanc", "url":"https://x.com/inancgumus"}'
//
// Resolve a shortened URL:
//
//	$ curl localhost:8080/r/inanc
//
// Health check:
//
//	$ curl localhost:8080/health
package link
```

## [bite](https://github.com/inancgumus/gobyexample/blob/366ab8018e0fdf51963209e567578174a54d3145/bite) / [link](https://github.com/inancgumus/gobyexample/blob/366ab8018e0fdf51963209e567578174a54d3145/bite/link) / [server.go](https://github.com/inancgumus/gobyexample/blob/366ab8018e0fdf51963209e567578174a54d3145/bite/link/server.go)

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

	// Exercise: Since http.Handler is an exported type,
	// link.Server exposes an unnecessary Handler field
	// by embedding Handler directly.
	//
	// Declare a new unexported interface type in the
	// link package with a ServeHTTP method, and embed
	// that interface in Server instead.
	//
	// Hint: Go uses structural typing, so it doesn't
	// matter which interface provides ServeHTTP. Once
	// Server embeds this unexported interface, you can
	// still assign a ServeMux to that field.
	http.Handler
}

// NewServer returns a new [Server].
func NewServer(lg *slog.Logger, store *Store) *Server {
	mux := http.NewServeMux()
	srv := &Server{
		lg:      lg,
		Handler: mux,
	}
	mux.HandleFunc("POST /shorten", srv.Shorten(store))
	mux.HandleFunc("GET /r/{key}", srv.Resolve(store))
	mux.HandleFunc("GET /health", srv.Health)
	return srv
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
//
// Uses the "key" [http.Request.PathValue] to resolve the link.
//
//	Status Code       Condition
//	302               The link is successfully resolved.
//	400               The request is invalid.
//	404               The link does not exist.
//	500               There is an internal error.
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


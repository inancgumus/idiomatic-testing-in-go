package link

import (
	"fmt"
	"net/http"

	"github.com/inancgumus/gobyexample/bite"
	"github.com/inancgumus/gobyexample/bite/httpio"
)

// ------------------------------------------------------------
// NOTE TO THE READER:
//
// See server.interfaces.go for a server that uses interfaces.
// ------------------------------------------------------------

// Server is a URL shortener HTTP server.
type Server struct {
	http.Handler
	// add other fields here...
}

// NewServer returns a new URL shortener server.
func NewServer(links *Store) *Server {
	mux := http.NewServeMux()
	mux.Handle("POST /shorten", Shorten(links))
	mux.Handle("GET /r/{key}", Resolve(links))
	mux.HandleFunc("GET /health", Health)

	return &Server{
		Handler: mux,
	}
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
func Shorten(links *Store) httpio.Handler {
	return func(w http.ResponseWriter, r *http.Request) httpio.Handler {
		var link Link

		max := http.MaxBytesReader(w, r.Body, 4_096)
		if err := httpio.DecodeJSON(max, &link); err != nil {
			return httpio.Error(invalidRequest(err))
		}
		if err := links.Create(r.Context(), link); err != nil {
			return httpio.Error(err)
		}

		return httpio.Code(http.StatusCreated, httpio.JSON(map[string]string{
			"key": link.Key,
		}))
	}
}

// Resolve handles the URL resolving requests for the short links.
//
//	Status Code       Condition
//	302               The link is successfully resolved.
//	400               The request is invalid.
//	404               The link does not exist.
//	500               There is an internal error.
func Resolve(links *Store) httpio.Handler {
	return func(w http.ResponseWriter, r *http.Request) httpio.Handler {
		link, err := links.Retrieve(r.Context(), r.PathValue("key"))
		if err != nil {
			return httpio.Error(err)
		}
		http.Redirect(w, r, link.URL, http.StatusFound)

		return httpio.OK
	}
}

// Health handles the health check requests.
//
//	Status Code       Condition
//	200               The server is healthy.
func Health(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "OK") }

func invalidRequest(err error) error {
	return fmt.Errorf("%w: %s", bite.ErrInvalidRequest, err) //nolint:errorlint
}

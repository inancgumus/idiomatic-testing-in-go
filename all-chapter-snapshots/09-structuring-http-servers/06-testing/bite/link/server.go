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
	http.Handler
}

// NewServer returns a new [Server].
func NewServer(lg *slog.Logger, store *Store) *Server {
	mux := http.NewServeMux()
	srv := &Server{
		lg:      lg,
		Handler: mux,
	}
	mux.Handle("POST /shorten", srv.Shorten(store))
	mux.Handle("GET /r/{key}", srv.Resolve(store))
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

// Resolve handles the URL resolving requests for the shortened links.
//
// Uses the "key" [http.Request.PathValue] to resolve the link.
//
//	Status Code       Condition
//	302               The link is successfully resolved.
//	400               The request is invalid.
//	404               The link does not exist.
//	500               There is an internal error.
func (srv *Server) Resolve(links *Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := LinkKey(r.PathValue("key"))

		if err := key.Validate(); err != nil {
			httpError(w, fmt.Errorf("validating link: %w: %w",
				err, bite.ErrInvalidRequest))
			return
		}
		link, err := links.Retrieve(r.Context(), key)
		if err != nil {
			httpError(w, fmt.Errorf("retrieving link: %w", err))
			return
		}
		http.Redirect(w, r, link.URL, http.StatusFound)
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

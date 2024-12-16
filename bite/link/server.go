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

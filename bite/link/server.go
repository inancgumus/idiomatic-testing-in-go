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

package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/inancgumus/gobyexample/bite/link"
)

func main() {
	const addr = "localhost:8080"

	log := slog.New(slog.NewTextHandler(os.Stderr, nil)).With("app", "linkd")
	log.Info("starting", "addr", addr)

	linkServer := link.NewServer(log, new(link.Store))

	srv := &http.Server{
		Addr:        addr,
		Handler:     linkServer,
		ReadTimeout: 20 * time.Second,
		IdleTimeout: 40 * time.Second,
	}
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Error("server closed unexpectedly", "message", err)
	}
}

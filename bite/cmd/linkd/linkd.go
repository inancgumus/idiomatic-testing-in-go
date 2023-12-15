package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/inancgumus/gobyexample/bite/httpio"
	"github.com/inancgumus/gobyexample/bite/link"
	"github.com/inancgumus/gobyexample/bite/sqlx"
)

// ---------------------------------------------------------
// exercise: make main() testable as shown in /hit/cmd/cli
//
// main is intentionally written in a way that it's not
// testable. Your goal is to make it testable.
//
// - create a config type
// 		- flags should write to the config fields
// 		- env vars should write to the config fields
// 		- validate the config
// - create an env type
// - create the run and runServer functions
// - pass the env from main to a run function
// - run will call runServer with the env and the config
// - instead of logging in main, return and print the error
// - use os.Exit(1) to exit the program from main with an
//   error status code
// ---------------------------------------------------------

func main() {
	// exercise: make these flags configurable.
	const addr = "localhost:8080"
	const timeout = 10 * time.Second
	const dbDSN = "file:bite.db?mode=rwc"

	log := slog.New(&httpio.LogHandler{
		Handler: slog.NewTextHandler(os.Stderr, nil),
	})
	log = log.With("app", "linkd")
	slog.SetDefault(log)

	log.Info("starting the server", "addr", addr)

	ctx := context.Background()
	db, err := sqlx.Dial(ctx, sqlx.DefaultDriver, dbDSN)
	if err != nil {
		log.Error("connecting to database", "message", err)
		return
	}

	links := link.NewServer(link.NewStore(db))
	handler := http.TimeoutHandler(links, timeout, "timeout")
	handler = httpio.LoggingMiddleware(handler)
	handler = httpio.TraceMiddleware(handler)

	srv := &http.Server{
		Addr:        addr,
		Handler:     handler,
		ReadTimeout: timeout * 2,
		IdleTimeout: timeout * 4,
	}
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Error("server closed unexpectedly", "message", err)
	}
}

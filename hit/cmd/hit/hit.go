package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/inancgumus/gobyexample/hit"
)

const logo = `
 __  __     __     ______
/\ \_\ \   /\ \   /\__  _\
\ \  __ \  \ \ \  \/_/\ \/
 \ \_\ \_\  \ \_\    \ \_\
  \/_/\/_/   \/_/     \/_/`

func main() {
	e := &env{
		stdout: os.Stdout,
		stderr: os.Stderr,
		args:   os.Args,
	}
	if err := run(e); err != nil {
		os.Exit(1)
	}
}

func run(e *env) error {
	c := config{
		n: 100, // default request count
		c: 1,   // default concurrency level
	}
	if err := parseArgs(&c, e.args[1:], e.stderr); err != nil {
		return err
	}

	fmt.Fprintf(
		e.stdout,
		"%s\n\nSending %d requests to %q (concurrency: %d)\n",
		logo, c.n, c.url, c.c,
	)
	if e.dryRun {
		return nil
	}

	return runHit(e, &c)
}

// runHit makes the requests and prints the results.
func runHit(e *env, c *config) error {
	// runHit is like a mini program that handles its own errors
	// and prints them to the standard error. Returning also the
	// error is for convenience and testing purposes.
	handleErr := func(err error) error {
		if err != nil {
			fmt.Fprintf(e.stderr, "\nerror occurred: %v\n", err)
			return err
		}
		return nil
	}

	// exercise: define timeout flags and use them here
	const (
		timeout           = time.Hour
		timeoutPerRequest = 30 * time.Second
	)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, timeout)
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	defer stop()

	request, err := http.NewRequest(http.MethodGet, c.url, http.NoBody)
	if err != nil {
		return handleErr(fmt.Errorf("new request: %w", err))
	}
	client := &hit.Client{
		C:       c.c,
		RPS:     c.rps,
		Timeout: timeoutPerRequest,
	}
	sum := client.Do(ctx, request, c.n)
	sum.Fprint(e.stdout)

	if err = ctx.Err(); errors.Is(err, context.DeadlineExceeded) { // set the error
		return handleErr(fmt.Errorf("timed out in %s", timeout))
	}

	return handleErr(err)
}

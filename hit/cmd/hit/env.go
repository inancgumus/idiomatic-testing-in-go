package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/url"
	"strconv"
)

// env is the program's environment.
// It abstracts the standard output, standard error, command-line
// arguments to make the program independent of its environment.
type env struct {
	stdout io.Writer // stdout abstracts standard output
	stderr io.Writer // stderr abstracts standard error
	args   []string  // args are command-line arguments
	dryRun bool      // dryRun enables dry mode
}

// config holds the program's configuration.
type config struct {
	url string // url to send requests
	n   int    // n is the number of requests
	c   int    // c is the concurrency level
	rps int    // rps is the requests per second
}

// parseArgs parses command-line flags and positional arguments
// in args and updates c with the parsed values. c can provide
// default values for the flags and positional arguments before
// calling parseArgs.
func parseArgs(c *config, args []string, stderr io.Writer) error {
	fs := flag.NewFlagSet("hit", flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "usage: %s [options] url\n", fs.Name())
		fs.PrintDefaults()
	}

	fs.Var(asPositiveIntValue(&c.n), "n", "Number of requests")
	fs.Var(asPositiveIntValue(&c.c), "c", "Concurrency level")
	fs.Var(asPositiveIntValue(&c.rps), "rps", "Requests per second")
	if err := fs.Parse(args); err != nil {
		return err
	}
	c.url = fs.Arg(0)

	if err := validateArgs(c); err != nil {
		fmt.Fprintln(fs.Output(), err)
		fs.Usage()
		return err
	}

	return nil
}

// validateArgs validates c's fields for expected flag values and
// positional arguments. It returns an error if any of the fields
// are invalid.
func validateArgs(c *config) error {
	const urlArg = "argument url"

	u, err := url.Parse(c.url)
	if err != nil {
		return argError(c.url, urlArg, err)
	}
	if c.url == "" || u.Host == "" || u.Scheme == "" {
		return argError(c.url, urlArg, errors.New("requires a valid url"))
	}
	if c.n < c.c {
		err := fmt.Errorf(`should be greater than -c: "%d"`, c.c)
		return argError(c.n, "flag -n", err)
	}

	return nil
}

// argError returns an error message for an invalid argument.
func argError(value any, arg string, err error) error {
	return fmt.Errorf(`invalid value "%v" for %s: %w`, value, arg, err)
}

// positiveIntValue is a custom flag.Value that only
// accepts positive integer values.
type positiveIntValue int

// asPositiveInt returns an int pointer as a positiveIntValue.
func asPositiveIntValue(p *int) *positiveIntValue {
	return (*positiveIntValue)(p)
}

// Set sets the value of the positiveIntValue.
func (n *positiveIntValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		return err
	}
	if v <= 0 {
		return errors.New("should be greater than zero")
	}
	*n = positiveIntValue(v)
	return nil
}

// String returns the string representation of the value.
func (n *positiveIntValue) String() string {
	return strconv.Itoa(int(*n))
}

/* exercise: add a custom positiveDurationValue */

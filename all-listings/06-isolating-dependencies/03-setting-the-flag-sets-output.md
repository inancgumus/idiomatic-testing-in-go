# Listing 6.3: Setting the flag set's output

## [hit](https://github.com/inancgumus/gobyexample/blob/993cabf212c489996003563d3e7386e4e4c654f6/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/993cabf212c489996003563d3e7386e4e4c654f6/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/993cabf212c489996003563d3e7386e4e4c654f6/hit/cmd/hit) / [env.go](https://github.com/inancgumus/gobyexample/blob/993cabf212c489996003563d3e7386e4e4c654f6/hit/cmd/hit/env.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/url"
	"strconv"
)

type env struct {
	stdout io.Writer // stdout abstracts standard output
	stderr io.Writer // stderr abstracts standard error
	args   []string  // args are command-line arguments
	dryRun bool      // dryRun enables dry mode
}

type config struct {
	url string // url to send requests
	n   int    // n is the number of requests
	c   int    // c is the concurrency level
	rps int    // rps is the requests per second
}

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

func validateArgs(c *config) error {
	const urlArg = "url argument"

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

type positiveIntValue int

func asPositiveIntValue(p *int) *positiveIntValue {
	return (*positiveIntValue)(p)
}

func (n *positiveIntValue) String() string {
	return strconv.Itoa(int(*n))
}

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
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -23,8 +23,10 @@ type config struct {
 	rps int    // rps is the requests per second
 }
 
-func parseArgs(c *config, args []string) error {
+func parseArgs(c *config, args []string, stderr io.Writer) error {
 	fs := flag.NewFlagSet("hit", flag.ContinueOnError)
+	fs.SetOutput(stderr)
+
 	fs.Usage = func() {
 		fmt.Fprintf(fs.Output(), "usage: %s [options] url\n", fs.Name())
 		fs.PrintDefaults()
```


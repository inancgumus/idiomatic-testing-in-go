# Listing 5.10: Defining a custom flag type

## [hit](https://github.com/inancgumus/gobyexample/blob/0ba1d91f0fe1920079509d39981c1d4df561b3f5/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/0ba1d91f0fe1920079509d39981c1d4df561b3f5/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/0ba1d91f0fe1920079509d39981c1d4df561b3f5/hit/cmd/hit) / [env.go](https://github.com/inancgumus/gobyexample/blob/0ba1d91f0fe1920079509d39981c1d4df561b3f5/hit/cmd/hit/env.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"errors"
	"flag"
	"strconv"
)

type config struct {
	url string // url to send requests
	n   int    // n is the number of requests
	c   int    // c is the concurrency level
	rps int    // rps is the requests per second
}

func parseArgs(c *config, args []string) error {
	fs := flag.NewFlagSet("hit", flag.ContinueOnError)

	fs.StringVar(&c.url, "url", "", "HTTP server `URL` (required)")
	fs.Var(asPositiveIntValue(&c.n), "n", "Number of requests")
	fs.Var(asPositiveIntValue(&c.c), "c", "Concurrency level")
	fs.Var(asPositiveIntValue(&c.rps), "rps", "Requests per second")

	return fs.Parse(args)
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
@@ -16,10 +16,10 @@ type config struct {
 func parseArgs(c *config, args []string) error {
 	fs := flag.NewFlagSet("hit", flag.ContinueOnError)
 
-	fs.StringVar(&c.url, "url", "", "HTTP server URL ")
-	fs.IntVar(&c.n, "n", c.n, "Number of requests")
-	fs.IntVar(&c.c, "c", c.c, "Concurrency level")
-	fs.IntVar(&c.rps, "rps", c.rps, "Requests per second")
+	fs.StringVar(&c.url, "url", "", "HTTP server `URL` (required)")
+	fs.Var(asPositiveIntValue(&c.n), "n", "Number of requests")
+	fs.Var(asPositiveIntValue(&c.c), "c", "Concurrency level")
+	fs.Var(asPositiveIntValue(&c.rps), "rps", "Requests per second")
 
 	return fs.Parse(args)
 }
```


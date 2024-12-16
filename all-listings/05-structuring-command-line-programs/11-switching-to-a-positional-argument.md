# Listing 5.11: Switching to a positional argument

## [hit](https://github.com/inancgumus/gobyexample/blob/145fb1a6ab0b5ab894cb14e3a7337055f11bcb77/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/145fb1a6ab0b5ab894cb14e3a7337055f11bcb77/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/145fb1a6ab0b5ab894cb14e3a7337055f11bcb77/hit/cmd/hit) / [env.go](https://github.com/inancgumus/gobyexample/blob/145fb1a6ab0b5ab894cb14e3a7337055f11bcb77/hit/cmd/hit/env.go)

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

	return nil
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
@@ -3,6 +3,7 @@ package main
 import (
 	"errors"
 	"flag"
+	"fmt"
 	"strconv"
 )
 
@@ -15,13 +16,20 @@ type config struct {
 
 func parseArgs(c *config, args []string) error {
 	fs := flag.NewFlagSet("hit", flag.ContinueOnError)
+	fs.Usage = func() {
+		fmt.Fprintf(fs.Output(), "usage: %s [options] url\n", fs.Name())
+		fs.PrintDefaults()
+	}
 
-	fs.StringVar(&c.url, "url", "", "HTTP server `URL` (required)")
 	fs.Var(asPositiveIntValue(&c.n), "n", "Number of requests")
 	fs.Var(asPositiveIntValue(&c.c), "c", "Concurrency level")
 	fs.Var(asPositiveIntValue(&c.rps), "rps", "Requests per second")
+	if err := fs.Parse(args); err != nil {
+		return err
+	}
+	c.url = fs.Arg(0)
 
-	return fs.Parse(args)
+	return nil
 }
 
 type positiveIntValue int
```


# Listing 5.7: Parsing flags using a flag set

## [hit](https://github.com/inancgumus/gobyexample/blob/be1a37d86e6c4be0301071736f9c456f1f81811e/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/be1a37d86e6c4be0301071736f9c456f1f81811e/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/be1a37d86e6c4be0301071736f9c456f1f81811e/hit/cmd/hit) / [env.go](https://github.com/inancgumus/gobyexample/blob/be1a37d86e6c4be0301071736f9c456f1f81811e/hit/cmd/hit/env.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import "flag"

type config struct {
	url string // url to send requests
	n   int    // n is the number of requests
	c   int    // c is the concurrency level
	rps int    // rps is the requests per second
}

func parseArgs(c *config, args []string) error {
	fs := flag.NewFlagSet("hit", flag.ContinueOnError)

	fs.StringVar(&c.url, "url", "", "HTTP server URL ")
	fs.IntVar(&c.n, "n", c.n, "Number of requests")
	fs.IntVar(&c.c, "c", c.c, "Concurrency level")
	fs.IntVar(&c.rps, "rps", c.rps, "Requests per second")

	return fs.Parse(args)
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -1,10 +1,6 @@
 package main
 
-import (
-	"fmt"
-	"strconv"
-	"strings"
-)
+import "flag"
 
 type config struct {
 	url string // url to send requests
@@ -13,43 +9,13 @@ type config struct {
 	rps int    // rps is the requests per second
 }
 
-type parseFunc func(string) error
+func parseArgs(c *config, args []string) error {
+	fs := flag.NewFlagSet("hit", flag.ContinueOnError)
 
-func stringVar(p *string) parseFunc {
-	return func(s string) error {
-		*p = s
-		return nil
-	}
-}
+	fs.StringVar(&c.url, "url", "", "HTTP server URL ")
+	fs.IntVar(&c.n, "n", c.n, "Number of requests")
+	fs.IntVar(&c.c, "c", c.c, "Concurrency level")
+	fs.IntVar(&c.rps, "rps", c.rps, "Requests per second")
 
-func intVar(p *int) parseFunc {
-	return func(s string) error {
-		var err error
-		*p, err = strconv.Atoi(s)
-		return err
-	}
-}
-
-func parseArgs(c *config, args []string) error {
-	flagSet := map[string]parseFunc{
-		"url": stringVar(&c.url),
-		"n":   intVar(&c.n),
-		"c":   intVar(&c.c),
-		"rps": intVar(&c.rps),
-	}
-	for _, arg := range args {
-		fn, fv, ok := strings.Cut(arg, "=")
-		if !ok {
-			continue // wrong flag format
-		}
-		fn = strings.TrimPrefix(fn, "-")
-		parseValue, ok := flagSet[fn]
-		if !ok {
-			continue // not in flagSet
-		}
-		if err := parseValue(fv); err != nil {
-			return fmt.Errorf("invalid value %q for flag %s: %w", fv, fn, err)
-		}
-	}
-	return nil
+	return fs.Parse(args)
 }
```


# Listing 5.9: Implementing a custom dynamic value

## [hit](https://github.com/inancgumus/gobyexample/blob/95eed0e065e2adfb0beb02fb3aec0bd97a3c844a/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/95eed0e065e2adfb0beb02fb3aec0bd97a3c844a/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/95eed0e065e2adfb0beb02fb3aec0bd97a3c844a/hit/cmd/hit) / [env.go](https://github.com/inancgumus/gobyexample/blob/95eed0e065e2adfb0beb02fb3aec0bd97a3c844a/hit/cmd/hit/env.go)

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

	fs.StringVar(&c.url, "url", "", "HTTP server URL ")
	fs.IntVar(&c.n, "n", c.n, "Number of requests")
	fs.IntVar(&c.c, "c", c.c, "Concurrency level")
	fs.IntVar(&c.rps, "rps", c.rps, "Requests per second")

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
@@ -1,6 +1,10 @@
 package main
 
-import "flag"
+import (
+	"errors"
+	"flag"
+	"strconv"
+)
 
 type config struct {
 	url string // url to send requests
@@ -19,3 +23,26 @@ func parseArgs(c *config, args []string) error {
 
 	return fs.Parse(args)
 }
+
+type positiveIntValue int
+
+func asPositiveIntValue(p *int) *positiveIntValue {
+	return (*positiveIntValue)(p)
+}
+
+func (n *positiveIntValue) String() string {
+	return strconv.Itoa(int(*n))
+}
+
+func (n *positiveIntValue) Set(s string) error {
+	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
+	if err != nil {
+		return err
+	}
+	if v <= 0 {
+		return errors.New("should be greater than zero")
+	}
+	*n = positiveIntValue(v)
+
+	return nil
+}
```


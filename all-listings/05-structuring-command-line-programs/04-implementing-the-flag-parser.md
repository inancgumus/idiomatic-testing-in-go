# Listing 5.4: Implementing the flag parser

## [hit](https://github.com/inancgumus/gobyexample/blob/7f871d9038daeecc8c14aea478205a04439e98c6/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/7f871d9038daeecc8c14aea478205a04439e98c6/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/7f871d9038daeecc8c14aea478205a04439e98c6/hit/cmd/hit) / [env.go](https://github.com/inancgumus/gobyexample/blob/7f871d9038daeecc8c14aea478205a04439e98c6/hit/cmd/hit/env.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type config struct {
	url string // url to send requests
	n   int    // n is the number of requests
	c   int    // c is the concurrency level
	rps int    // rps is the requests per second
}

type parseFunc func(string) error

func stringVar(p *string) parseFunc {
	return func(s string) error {
		*p = s
		return nil
	}
}

func intVar(p *int) parseFunc {
	return func(s string) error {
		var err error
		*p, err = strconv.Atoi(s)
		return err
	}
}

func parseArgs(c *config, args []string) error {
	flagSet := map[string]parseFunc{
		"url": stringVar(&c.url),
		"n":   intVar(&c.n),
		"c":   intVar(&c.c),
		"rps": intVar(&c.rps),
	}
	for _, arg := range args {
		fn, fv, ok := strings.Cut(arg, "=")
		if !ok {
			continue // wrong flag format
		}
		fn = strings.TrimPrefix(fn, "-")
		parseValue, ok := flagSet[fn]
		if !ok {
			continue // not in flagSet
		}
		if err := parseValue(fv); err != nil {
			return fmt.Errorf("invalid value %q for flag %s: %w", fv, fn, err)
		}
	}
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
 
-import "strconv"
+import (
+	"fmt"
+	"strconv"
+	"strings"
+)
 
 type config struct {
 	url string // url to send requests
@@ -25,3 +29,27 @@ func intVar(p *int) parseFunc {
 		return err
 	}
 }
+
+func parseArgs(c *config, args []string) error {
+	flagSet := map[string]parseFunc{
+		"url": stringVar(&c.url),
+		"n":   intVar(&c.n),
+		"c":   intVar(&c.c),
+		"rps": intVar(&c.rps),
+	}
+	for _, arg := range args {
+		fn, fv, ok := strings.Cut(arg, "=")
+		if !ok {
+			continue // wrong flag format
+		}
+		fn = strings.TrimPrefix(fn, "-")
+		parseValue, ok := flagSet[fn]
+		if !ok {
+			continue // not in flagSet
+		}
+		if err := parseValue(fv); err != nil {
+			return fmt.Errorf("invalid value %q for flag %s: %w", fv, fn, err)
+		}
+	}
+	return nil
+}
```


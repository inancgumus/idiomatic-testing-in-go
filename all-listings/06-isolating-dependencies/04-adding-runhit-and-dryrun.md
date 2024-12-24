# Listing 6.4: Adding runHit and dryRun

## [hit](https://github.com/inancgumus/gobyexample/blob/fbf4778bea8914ff86d481f6a50ff15a248460ea/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/fbf4778bea8914ff86d481f6a50ff15a248460ea/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/fbf4778bea8914ff86d481f6a50ff15a248460ea/hit/cmd/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/fbf4778bea8914ff86d481f6a50ff15a248460ea/hit/cmd/hit/hit.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"fmt"
	"io"
	"os"
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

	if err := runHit(e.stdout, &c); err != nil {
		fmt.Fprintf(e.stderr, "\nerror occurred: %v\n", err)
		return err
	}

	return nil
}

func runHit(stdout io.Writer, c *config) error {
	/* TODO: integrate the hit package */
	return nil
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -2,6 +2,7 @@ package main
 
 import (
 	"fmt"
+	"io"
 	"os"
 )
 
@@ -26,17 +27,29 @@ func main() {
 func run(e *env) error {
 	c := config{
 		n: 100, // default request count
 		c: 1,   // default concurrency level
 	}
 	if err := parseArgs(&c, e.args[1:], e.stderr); err != nil {
 		return err
 	}
-
 	fmt.Fprintf(
 		e.stdout,
 		"%s\n\nSending %d requests to %q (concurrency: %d)\n",
 		logo, c.n, c.url, c.c,
 	)
+	if e.dryRun {
+		return nil
+	}
+
+	if err := runHit(e.stdout, &c); err != nil {
+		fmt.Fprintf(e.stderr, "\nerror occurred: %v\n", err)
+		return err
+	}
+
+	return nil
+}
 
+func runHit(stdout io.Writer, c *config) error {
+	/* TODO: integrate the hit package */
 	return nil
 }
```


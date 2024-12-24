# Listing 6.2: Calling run from main

## [hit](https://github.com/inancgumus/gobyexample/blob/a3a113848935bce7b33688869d3dfa1cf0cd960e/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/a3a113848935bce7b33688869d3dfa1cf0cd960e/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/a3a113848935bce7b33688869d3dfa1cf0cd960e/hit/cmd/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/a3a113848935bce7b33688869d3dfa1cf0cd960e/hit/cmd/hit/hit.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"fmt"
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

	return nil
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -13,16 +13,30 @@ const logo = `
   \/_/\/_/   \/_/     \/_/`
 
 func main() {
+	e := &env{
+		stdout: os.Stdout,
+		stderr: os.Stderr,
+		args:   os.Args,
+	}
+	if err := run(e); err != nil {
+		os.Exit(1)
+	}
+}
+
+func run(e *env) error {
 	c := config{
 		n: 100, // default request count
 		c: 1,   // default concurrency level
 	}
-	if err := parseArgs(&c, os.Args[1:]); err != nil {
-		os.Exit(1)
+	if err := parseArgs(&c, e.args[1:], e.stderr); err != nil {
+		return err
 	}
-	fmt.Printf(
+
+	fmt.Fprintf(
+		e.stdout,
 		"%s\n\nSending %d requests to %q (concurrency: %d)\n",
 		logo, c.n, c.url, c.c,
 	)
-	/* package hit integration here */
+
+	return nil
 }
```


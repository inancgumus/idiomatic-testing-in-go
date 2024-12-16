# Listing 5.5: Integrating the custom parser

## [hit](https://github.com/inancgumus/gobyexample/blob/205e1cae64f6979c1c2d8be3492ef51c5fff5f7b/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/205e1cae64f6979c1c2d8be3492ef51c5fff5f7b/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/205e1cae64f6979c1c2d8be3492ef51c5fff5f7b/hit/cmd/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/205e1cae64f6979c1c2d8be3492ef51c5fff5f7b/hit/cmd/hit/hit.go)

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

const usage = `
Usage:
  -url
       HTTP server URL (required)
  -n
       Number of requests
  -c
       Concurrency level
  -rps
       Requests per second`

func main() {
	var c config
	if err := parseArgs(&c, os.Args[1:]); err != nil {
		fmt.Printf("%s\n%s", err, usage)
		os.Exit(1)
	}
	fmt.Printf(
		"%s\n\nSending %d requests to %q (concurrency: %d)\n",
		logo, c.n, c.url, c.c,
	)
	/* package hit integration here */
}
```

## What's changed?

> [!TIP]
> The following diff shows the changes made to the file since the last commit.
> The lines starting with `+` show the new lines added, and the lines starting with `-` show the lines removed.
> The lines starting with `@@` show the line and column numbers of the changes.

```diff
@@ -1,6 +1,9 @@
 package main
 
-import "fmt"
+import (
+	"fmt"
+	"os"
+)
 
 const logo = `
  __  __     __     ______
@@ -21,7 +24,14 @@ Usage:
        Requests per second`
 
 func main() {
-	fmt.Printf("%s\n%s", logo, usage)
-
-	/* TODO: integrate package hit */
+	var c config
+	if err := parseArgs(&c, os.Args[1:]); err != nil {
+		fmt.Printf("%s\n%s", err, usage)
+		os.Exit(1)
+	}
+	fmt.Printf(
+		"%s\n\nSending %d requests to %q (concurrency: %d)\n",
+		logo, c.n, c.url, c.c,
+	)
+	/* package hit integration here */
 }
```


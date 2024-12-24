# Listing 5.6: Setting sensible defaults

## [hit](https://github.com/inancgumus/gobyexample/blob/3e7ae9da5937bc22b2f9497c65e223784c7c4441/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/3e7ae9da5937bc22b2f9497c65e223784c7c4441/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/3e7ae9da5937bc22b2f9497c65e223784c7c4441/hit/cmd/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/3e7ae9da5937bc22b2f9497c65e223784c7c4441/hit/cmd/hit/hit.go)

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
	c := config{
		n: 100, // default request count
		c: 1,   // default concurrency level
	}
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
@@ -24,14 +24,17 @@ Usage:
        Requests per second`
 
 func main() {
-	var c config
+	c := config{
+		n: 100, // default request count
+		c: 1,   // default concurrency level
+	}
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


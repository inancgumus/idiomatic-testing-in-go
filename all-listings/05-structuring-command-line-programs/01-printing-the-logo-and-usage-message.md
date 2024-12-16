# Listing 5.1: Printing the logo and usage message

## [hit](https://github.com/inancgumus/gobyexample/blob/95135aef9c1a6abe53b019d859e913c553effdc9/hit) / [cmd](https://github.com/inancgumus/gobyexample/blob/95135aef9c1a6abe53b019d859e913c553effdc9/hit/cmd) / [hit](https://github.com/inancgumus/gobyexample/blob/95135aef9c1a6abe53b019d859e913c553effdc9/hit/cmd/hit) / [hit.go](https://github.com/inancgumus/gobyexample/blob/95135aef9c1a6abe53b019d859e913c553effdc9/hit/cmd/hit/hit.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import "fmt"

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
	fmt.Printf("%s\n%s", logo, usage)

	/* TODO: integrate package hit */
}
```


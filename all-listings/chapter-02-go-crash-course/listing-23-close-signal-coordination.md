# Listing 2.23: Close signal coordination

## [concurrency](https://github.com/inancgumus/gobyexample/blob/bb1188bdce43d4921bb10b50bfbcc6cb8eecd86c/concurrency) / [barrier](https://github.com/inancgumus/gobyexample/blob/bb1188bdce43d4921bb10b50bfbcc6cb8eecd86c/concurrency/barrier) / [main.go](https://github.com/inancgumus/gobyexample/blob/bb1188bdce43d4921bb10b50bfbcc6cb8eecd86c/concurrency/barrier/main.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"fmt"

	"github.com/inancgumus/gobyexample/concurrency/syncx"
)

func main() {
	var sg syncx.SafeGroup

	wait := make(chan struct{})
	for range 10 {
		sg.Go(func() {
			<-wait
			fmt.Print("go!")
		})
	}
	// do other work
	close(wait)
	sg.Wait()
	// do other work
}
```


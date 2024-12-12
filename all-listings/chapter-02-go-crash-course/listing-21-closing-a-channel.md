# Listing 2.21: Closing a channel

## [concurrency](https://github.com/inancgumus/gobyexample/blob/fbc057c4c2cef2a64d91c3c63f8484c75d3e298f/concurrency) / [forrange](https://github.com/inancgumus/gobyexample/blob/fbc057c4c2cef2a64d91c3c63f8484c75d3e298f/concurrency/forrange) / [main.go](https://github.com/inancgumus/gobyexample/blob/fbc057c4c2cef2a64d91c3c63f8484c75d3e298f/concurrency/forrange/main.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	results := make(chan int)
	go func() {
		for n := range rand.N(100) + 1 { //nolint:gosec
			results <- max(1, n*2)
		}
		close(results)
	}()
	for {
		result, ok := <-results
		if !ok {
			break
		}
		fmt.Print(result, ".")
	}
}
```


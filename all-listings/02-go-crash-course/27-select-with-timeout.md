# Listing 2.27: Select with timeout

## [concurrency](https://github.com/inancgumus/gobyexample/blob/b34e95b56841bc6ae2ec890ff7bd46bec5d3b58b/concurrency) / [timeout](https://github.com/inancgumus/gobyexample/blob/b34e95b56841bc6ae2ec890ff7bd46bec5d3b58b/concurrency/timeout) / [main.go](https://github.com/inancgumus/gobyexample/blob/b34e95b56841bc6ae2ec890ff7bd46bec5d3b58b/concurrency/timeout/main.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	results := make(chan int, 1)
	go func() {
		time.Sleep(10 * time.Second)
		results <- 42
	}()
	select {
	case v := <-results:
		fmt.Println("result:", v)
	case <-time.After(5 * time.Second):
		fmt.Println("timed out")
	}
}
```


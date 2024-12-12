# Listing 2.17: WaitGroup to wait

## [concurrency](https://github.com/inancgumus/gobyexample/blob/2c75ad4bf9ee6e4b8f216601110bd4d87858e609/concurrency) / [waitgroup](https://github.com/inancgumus/gobyexample/blob/2c75ad4bf9ee6e4b8f216601110bd4d87858e609/concurrency/waitgroup) / [main.go](https://github.com/inancgumus/gobyexample/blob/2c75ad4bf9ee6e4b8f216601110bd4d87858e609/concurrency/waitgroup/main.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for id := range 10 {
		wg.Add(1)
		go func() {
			work(id + 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Print("main done.")
}

func work(id int) {
	time.Sleep(rand.N(10 * time.Second)) //nolint:gosec
	fmt.Printf("worker %d done.", id)
}
```


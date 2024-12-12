# Listing 2.16: Running concurrently

## [concurrency](https://github.com/inancgumus/gobyexample/blob/124d3b8131c9db783ceaae3b4a8db78acc21bd9e/concurrency) / [goroutines](https://github.com/inancgumus/gobyexample/blob/124d3b8131c9db783ceaae3b4a8db78acc21bd9e/concurrency/goroutines) / [main.go](https://github.com/inancgumus/gobyexample/blob/124d3b8131c9db783ceaae3b4a8db78acc21bd9e/concurrency/goroutines/main.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	go work(1)
	fmt.Print("main done.")
}

func work(id int) {
	time.Sleep(rand.N(10 * time.Second)) //nolint:gosec
	fmt.Printf("worker %d done.", id)
}
```


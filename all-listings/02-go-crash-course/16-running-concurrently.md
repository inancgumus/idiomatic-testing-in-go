# Listing 2.16: Running concurrently

## [concurrency](https://github.com/inancgumus/gobyexample/blob/d71a1d97200985feb46e5e564a0190c39efba18c/concurrency) / [goroutines](https://github.com/inancgumus/gobyexample/blob/d71a1d97200985feb46e5e564a0190c39efba18c/concurrency/goroutines) / [main.go](https://github.com/inancgumus/gobyexample/blob/d71a1d97200985feb46e5e564a0190c39efba18c/concurrency/goroutines/main.go)

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


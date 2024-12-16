# Listing 2.20: Sending and receiving

## [concurrency](https://github.com/inancgumus/gobyexample/blob/1fdbeb84f6996b25e353f2c28d7c2af628940036/concurrency) / [sendrecv](https://github.com/inancgumus/gobyexample/blob/1fdbeb84f6996b25e353f2c28d7c2af628940036/concurrency/sendrecv) / [main.go](https://github.com/inancgumus/gobyexample/blob/1fdbeb84f6996b25e353f2c28d7c2af628940036/concurrency/sendrecv/main.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import "fmt"

func calculate() int { return 1 }

func main() {
	results := make(chan int)
	for range 42 {
		go func() { results <- calculate() }()
	}
	var total int
	for range 42 {
		total += <-results
	}
	fmt.Println("the meaning of universe:", total)
}
```


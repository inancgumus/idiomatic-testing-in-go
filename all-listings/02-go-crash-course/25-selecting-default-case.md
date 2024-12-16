# Listing 2.25: Selecting default case

## [concurrency](https://github.com/inancgumus/gobyexample/blob/c99c800890b2121694a44edce0cb23fab718d728/concurrency) / [nonblocking](https://github.com/inancgumus/gobyexample/blob/c99c800890b2121694a44edce0cb23fab718d728/concurrency/nonblocking) / [main.go](https://github.com/inancgumus/gobyexample/blob/c99c800890b2121694a44edce0cb23fab718d728/concurrency/nonblocking/main.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import "fmt"

func isClosed[T any](done chan T) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	done := make(chan struct{})
	fmt.Print("open:", isClosed(done), ".")
	close(done)
	fmt.Print("open:", isClosed(done), ".")
}
```


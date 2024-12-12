# Listing 2.8: Calling methods

## [oop](https://github.com/inancgumus/gobyexample/blob/afb4cc24aeb33bde3c7557639d9506ce4bc808f3/oop) / [pointer-receivers](https://github.com/inancgumus/gobyexample/blob/afb4cc24aeb33bde3c7557639d9506ce4bc808f3/oop/pointer-receivers) / [main.go](https://github.com/inancgumus/gobyexample/blob/afb4cc24aeb33bde3c7557639d9506ce4bc808f3/oop/pointer-receivers/main.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import "fmt"

func main() {
	fileServer := &server{url: "file"}
	fileServer.check()
	fmt.Printf("is slow? %t\n", fileServer.slow())
}
```


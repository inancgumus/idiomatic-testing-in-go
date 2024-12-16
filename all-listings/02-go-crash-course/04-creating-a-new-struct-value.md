# Listing 2.4: Creating a new struct value

## [oop](https://github.com/inancgumus/gobyexample/blob/00f1a4a1299867f0eb3b8291e2f1ece8d8c95e32/oop) / [structs](https://github.com/inancgumus/gobyexample/blob/00f1a4a1299867f0eb3b8291e2f1ece8d8c95e32/oop/structs) / [main.go](https://github.com/inancgumus/gobyexample/blob/00f1a4a1299867f0eb3b8291e2f1ece8d8c95e32/oop/structs/main.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import "fmt"

func main() {
	fileServer := server{url: "file"}
	fmt.Printf(
		"url: %s response time: %d\n",
		fileServer.url, fileServer.responseTime,
	)
}
```


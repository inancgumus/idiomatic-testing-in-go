# Listing 2.2: Implementing the main package

## [packages](https://github.com/inancgumus/gobyexample/blob/c6efb740a76d52a6770afba159cdd0a3025c5a60/packages) / [hello.go](https://github.com/inancgumus/gobyexample/blob/c6efb740a76d52a6770afba159cdd0a3025c5a60/packages/hello.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"fmt"

	"github.com/inancgumus/gobyexample/packages/book"
)

func main() {
	fmt.Println(book.Title())
}
```


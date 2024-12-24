# Listing 2.15: A sequential function

## [concurrency](https://github.com/inancgumus/gobyexample/blob/d101d4ba813694896237f36a05267ce95c2cdf33/concurrency) / [sequential](https://github.com/inancgumus/gobyexample/blob/d101d4ba813694896237f36a05267ce95c2cdf33/concurrency/sequential) / [main.go](https://github.com/inancgumus/gobyexample/blob/d101d4ba813694896237f36a05267ce95c2cdf33/concurrency/sequential/main.go)

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

func work(id int) {
	time.Sleep(rand.N(10 * time.Second)) //nolint:gosec
	fmt.Printf("worker %d done.", id)
}

func main() {
	work(1)
	fmt.Print("main done.")
}
```


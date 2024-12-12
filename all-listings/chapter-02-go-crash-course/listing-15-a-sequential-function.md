# Listing 2.15: A sequential function

## [concurrency](https://github.com/inancgumus/gobyexample/blob/da344dcc555375018a32861ffae2f12adc5d0f59/concurrency) / [sequential](https://github.com/inancgumus/gobyexample/blob/da344dcc555375018a32861ffae2f12adc5d0f59/concurrency/sequential) / [main.go](https://github.com/inancgumus/gobyexample/blob/da344dcc555375018a32861ffae2f12adc5d0f59/concurrency/sequential/main.go)

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


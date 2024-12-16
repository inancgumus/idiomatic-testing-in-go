# Listing 2.19: Using SafeGroup

## [concurrency](https://github.com/inancgumus/gobyexample/blob/083a82dfc442c68bc26adb72a9ddd86df7352efb/concurrency) / [safegroup](https://github.com/inancgumus/gobyexample/blob/083a82dfc442c68bc26adb72a9ddd86df7352efb/concurrency/safegroup) / [main.go](https://github.com/inancgumus/gobyexample/blob/083a82dfc442c68bc26adb72a9ddd86df7352efb/concurrency/safegroup/main.go)

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

	"github.com/inancgumus/gobyexample/concurrency/syncx"
)

func main() {
	var sg syncx.SafeGroup
	for i := range 10 {
		sg.Go(func() { work(i + 1) })
	}
	sg.Wait()
	fmt.Print("main done.")
}

func work(id int) {
	time.Sleep(rand.N(10 * time.Second)) //nolint:gosec
	fmt.Printf("worker %d done.", id)
}
```


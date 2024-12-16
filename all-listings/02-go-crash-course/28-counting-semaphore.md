# Listing 2.28: Counting semaphore

## [concurrency](https://github.com/inancgumus/gobyexample/blob/1c02df8e92f626e33a4101ea9e7fa7edc9f98d10/concurrency) / [semaphore](https://github.com/inancgumus/gobyexample/blob/1c02df8e92f626e33a4101ea9e7fa7edc9f98d10/concurrency/semaphore) / [main.go](https://github.com/inancgumus/gobyexample/blob/1c02df8e92f626e33a4101ea9e7fa7edc9f98d10/concurrency/semaphore/main.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/gobyexample/concurrency/syncx"
)

func main() {
	type token struct{}
	tokens := make(chan token, 10)

	var sg syncx.SafeGroup
	for i := range 1000 {
		tokens <- token{}
		sg.Go(func() {
			upload(i)
			<-tokens
		})
	}
	sg.Wait()

	fmt.Println("done.")
}

func upload(image int) {
	fmt.Printf("%d.", image)
	time.Sleep(time.Second)
}
```


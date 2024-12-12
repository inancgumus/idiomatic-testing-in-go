# Listing 2.26: Buffered channel

## [concurrency](https://github.com/inancgumus/gobyexample/blob/287d1969c31fbc327064885fc24ead9de871aeea/concurrency) / [buffered](https://github.com/inancgumus/gobyexample/blob/287d1969c31fbc327064885fc24ead9de871aeea/concurrency/buffered) / [main.go](https://github.com/inancgumus/gobyexample/blob/287d1969c31fbc327064885fc24ead9de871aeea/concurrency/buffered/main.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import "fmt"

func main() {
	messages := make(chan string, 1)
	messages <- "hello"
	fmt.Println(<-messages)
	messages <- "hello"
	// messages <- "world"
}
```


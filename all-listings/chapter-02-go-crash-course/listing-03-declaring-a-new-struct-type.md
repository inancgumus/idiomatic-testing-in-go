# Listing 2.3: Declaring a new struct type

## [oop](https://github.com/inancgumus/gobyexample/blob/6c4a3485ea07ea4400c3e2da4b484b8af3f0cad6/oop) / [structs](https://github.com/inancgumus/gobyexample/blob/6c4a3485ea07ea4400c3e2da4b484b8af3f0cad6/oop/structs) / [server.go](https://github.com/inancgumus/gobyexample/blob/6c4a3485ea07ea4400c3e2da4b484b8af3f0cad6/oop/structs/server.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import "time"

type server struct {
	url          string
	responseTime time.Duration
}
```


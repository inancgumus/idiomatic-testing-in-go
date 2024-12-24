# Listing 2.7: Pointer receiver methods

## [oop](https://github.com/inancgumus/gobyexample/blob/234398940955da16ce0b24b96404cb384774507b/oop) / [pointer-receivers](https://github.com/inancgumus/gobyexample/blob/234398940955da16ce0b24b96404cb384774507b/oop/pointer-receivers) / [server.go](https://github.com/inancgumus/gobyexample/blob/234398940955da16ce0b24b96404cb384774507b/oop/pointer-receivers/server.go)

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

func (s *server) check()     { s.responseTime = 3 * time.Second }
func (s *server) slow() bool { return s.responseTime > 2*time.Second }
```


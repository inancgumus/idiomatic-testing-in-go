# Listing 2.7: Pointer receiver methods

## [oop](https://github.com/inancgumus/gobyexample/blob/f2c3897fad6ca576a410752fb5a83d5adc4d943e/oop) / [pointer-receivers](https://github.com/inancgumus/gobyexample/blob/f2c3897fad6ca576a410752fb5a83d5adc4d943e/oop/pointer-receivers) / [server.go](https://github.com/inancgumus/gobyexample/blob/f2c3897fad6ca576a410752fb5a83d5adc4d943e/oop/pointer-receivers/server.go)

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


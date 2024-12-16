# Listing 2.5: Methods with value receivers

## [oop](https://github.com/inancgumus/gobyexample/blob/f3cdacfe5e82c1be93c6d5946b8fd3190e71b2bb/oop) / [value-receivers](https://github.com/inancgumus/gobyexample/blob/f3cdacfe5e82c1be93c6d5946b8fd3190e71b2bb/oop/value-receivers) / [main.go](https://github.com/inancgumus/gobyexample/blob/f3cdacfe5e82c1be93c6d5946b8fd3190e71b2bb/oop/value-receivers/main.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package main

import "fmt"

type usage int

func (u usage) high() bool { return u >= 95 }
func (u usage) set(to int) { u = usage(to) } //nolint:staticcheck,ineffassign

func main() {
	var cpu usage = 99 // cpu is 99
	cpu.set(70)        // cpu is still 99
	fmt.Printf("cpu: %d%% high:%t\n", cpu, cpu.high())
}
```


# Listing 2.5: Methods with value receivers

## [oop](https://github.com/inancgumus/gobyexample/blob/837e26ee7bf651473c9dd6b2dc3ed7c1fbf96478/oop) / [value-receivers](https://github.com/inancgumus/gobyexample/blob/837e26ee7bf651473c9dd6b2dc3ed7c1fbf96478/oop/value-receivers) / [main.go](https://github.com/inancgumus/gobyexample/blob/837e26ee7bf651473c9dd6b2dc3ed7c1fbf96478/oop/value-receivers/main.go)

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


# Listing 2.18: A safer WaitGroup

## [concurrency](https://github.com/inancgumus/gobyexample/blob/b63221b2d5b6fd59446ea65eb6ed3d014a116085/concurrency) / [syncx](https://github.com/inancgumus/gobyexample/blob/b63221b2d5b6fd59446ea65eb6ed3d014a116085/concurrency/syncx) / [syncx.go](https://github.com/inancgumus/gobyexample/blob/b63221b2d5b6fd59446ea65eb6ed3d014a116085/concurrency/syncx/syncx.go)

> [!TIP]
> Each listing corresponds to a commit.
>
> Click the links above to see the file and its directory in their original locations and state as they were at the time of the commit.

```go
package syncx

import "sync"

// SafeGroup is a safe version of [sync.WaitGroup].
type SafeGroup struct{ wg sync.WaitGroup }

// Wait waits for the group of goroutines to finish.
func (sg *SafeGroup) Wait() { sg.wg.Wait() }

// Go runs the given function in a goroutine.
func (sg *SafeGroup) Go(fn func()) {
	sg.wg.Add(1)
	go func() {
		fn()
		sg.wg.Done()
	}()
}
```


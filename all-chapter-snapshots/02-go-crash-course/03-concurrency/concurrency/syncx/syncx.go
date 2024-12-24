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

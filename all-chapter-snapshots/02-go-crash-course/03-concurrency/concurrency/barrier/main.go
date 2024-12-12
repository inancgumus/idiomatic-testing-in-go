package main

import (
	"fmt"

	"github.com/inancgumus/gobyexample/concurrency/syncx"
)

func main() {
	var sg syncx.SafeGroup

	wait := make(chan struct{})
	stop := make(chan struct{})
	for range 10 {
		sg.Go(func() {
			select {
			case <-wait:
			case <-stop:
				fmt.Print("stopped.")
				return
			}
			fmt.Print("go!")
		})
	}
	// do other work
	close(stop)
	// Either close the stop channel or the wait channel
	// depending on your program's logic.
	sg.Wait()
	// do other work
}

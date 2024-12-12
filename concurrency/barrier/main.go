package main

import (
	"fmt"

	"github.com/inancgumus/gobyexample/concurrency/syncx"
)

func main() {
	var sg syncx.SafeGroup

	wait := make(chan struct{})
	for range 10 {
		sg.Go(func() {
			<-wait
			fmt.Print("go!")
		})
	}
	// do other work
	close(wait)
	sg.Wait()
	// do other work
}

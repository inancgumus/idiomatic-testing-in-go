package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/inancgumus/gobyexample/concurrency/syncx"
)

func main() {
	var sg syncx.SafeGroup
	for i := range 10 {
		sg.Go(func() { work(i + 1) })
	}
	sg.Wait()
	fmt.Print("main done.")
}

func work(id int) {
	time.Sleep(rand.N(10 * time.Second)) //nolint:gosec
	fmt.Printf("worker %d done.", id)
}

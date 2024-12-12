package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/gobyexample/concurrency/syncx"
)

func main() {
	type token struct{}
	tokens := make(chan token, 10)

	var sg syncx.SafeGroup
	for i := range 1000 {
		tokens <- token{}
		sg.Go(func() {
			upload(i)
			<-tokens
		})
	}
	sg.Wait()

	fmt.Println("done.")
}

func upload(image int) {
	fmt.Printf("%d.", image)
	time.Sleep(time.Second)
}

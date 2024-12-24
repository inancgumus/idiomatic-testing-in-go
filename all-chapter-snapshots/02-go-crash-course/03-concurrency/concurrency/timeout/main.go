package main

import (
	"fmt"
	"time"
)

func main() {
	results := make(chan int, 1)
	go func() {
		time.Sleep(10 * time.Second)
		results <- 42
	}()
	select {
	case v := <-results:
		fmt.Println("result:", v)
	case <-time.After(5 * time.Second):
		fmt.Println("timed out")
	}
}

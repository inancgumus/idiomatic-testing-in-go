package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	timeout := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		close(timeout)
	}()

	c := make(chan time.Duration, 2)
	go check("server1", c, timeout)
	go check("server2", c, timeout)

	select {
	case rt := <-c:
		fmt.Println("fastest response:", rt)
	case <-timeout:
		fmt.Println("timed out")
	}
}

func check(url string, c chan<- time.Duration, timeout <-chan struct{}) {
	d := rand.N(10 * time.Second)
	fmt.Printf("checking %s will take %s...\n", url, d)

	select {
	case <-time.After(d):
		c <- d
	case <-timeout:
		// do nothing here as the function will return
	}
}

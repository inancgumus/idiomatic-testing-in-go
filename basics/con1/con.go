package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	// incorrect: leaks a goroutine
	// c := make(chan time.Duration)

	c := make(chan time.Duration, 2)

	go check("server1", c)
	go check("server2", c)

	fmt.Println("fastest response:", <-c)

	// fmt.Println("average response:", (<-c + <-c) / 2)
}

func check(url string, c chan<- time.Duration) {
	d := rand.N(10 * time.Second)

	fmt.Printf("checking %s will take %s...\n", url, d)
	<-time.After(d)
	c <- d
}

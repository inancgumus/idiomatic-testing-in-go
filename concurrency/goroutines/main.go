package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	go work(1)
	fmt.Print("main done.")
}

func work(id int) {
	time.Sleep(rand.N(10 * time.Second)) //nolint:gosec
	fmt.Printf("worker %d done.", id)
}

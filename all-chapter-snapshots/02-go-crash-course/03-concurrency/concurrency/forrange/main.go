package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	results := make(chan int)
	go func() {
		for n := range rand.N(100) + 1 { //nolint:gosec
			results <- max(1, n*2)
		}
		close(results)
	}()
	for result := range results {
		fmt.Print(result, ".")
	}
}

package main

import "fmt"

func calculate() int { return 1 }

func main() {
	results := make(chan int)
	for range 42 {
		go func() { results <- calculate() }()
	}
	var total int
	for range 42 {
		total += <-results
	}
	fmt.Println("the meaning of universe:", total)
}

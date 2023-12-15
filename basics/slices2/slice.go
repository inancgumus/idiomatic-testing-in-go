package main

import "fmt"

func main() {
	hits := []int{20, 10, 5, 0, 40, 25}

	firstTwo := hits[:2]
	lastTwo := hits[len(hits)-2:]

	fmt.Println(firstTwo) // prints [20,10]
	fmt.Println(lastTwo)  // prints [40,25]
}

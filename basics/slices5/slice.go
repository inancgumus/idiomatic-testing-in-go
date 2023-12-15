package main

import "fmt"

func main() {
	hits := fetchHits()
	fmt.Println(hits)
}

func fetchHits() []int {
	hits := make([]int, 0, 1_000)
	hits = append(hits, 1, 2, 3)
	hits = append(hits, 4, 5, 6)
	// ...continue appending more elements...
	return hits
}

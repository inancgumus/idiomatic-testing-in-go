package main

func main() {
	_ = fetchHits()
}

func fetchHits() []int {
	var hits []int //nolint:prealloc

	for i := range 4 {
		hits = append(hits, i+1)
	}

	return hits
}

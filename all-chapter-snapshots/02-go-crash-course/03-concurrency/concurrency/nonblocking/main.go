package main

import "fmt"

func isClosed[T any](done chan T) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	done := make(chan struct{})
	fmt.Print("open:", isClosed(done), ".")
	close(done)
	fmt.Print("open:", isClosed(done), ".")
}

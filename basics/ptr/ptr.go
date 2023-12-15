package main

import "fmt"

func main() {
	counter := 42

	incrVal(counter) // does not change counter
	fmt.Println(counter) // prints 42

	incrPtr(&counter)
	fmt.Println(counter) // prints 43
}

func incrVal(c int) {
	c++ //nolint:ineffassign
} 

func incrPtr(c *int) {
	*c++
}

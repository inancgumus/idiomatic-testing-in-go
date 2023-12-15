package main

import "fmt"

func main() {
	usage := map[string]float64{
		"gateway": 75.,
		"auth":    50.5,
	}

	usage["health"] = 99.0
	fmt.Println(usage["health"]) // prints 99

	fmt.Println(len(usage)) // prints 3

	delete(usage, "health")

	v, ok := usage["health"]
	fmt.Println(v, ok) // prints 0 false

	clear(usage)
	fmt.Println(len(usage)) // prints 0
}

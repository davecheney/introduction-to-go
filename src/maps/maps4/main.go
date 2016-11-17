package main

import "fmt"

func main() {
	// map of planets to their number of moons
	moons := map[string]int{
		"Mercury": 0,
		"Venus":   0,
		"Earth":   1,
		"Mars":    2,
		"Jupiter": 67,
	}

	fmt.Println("Earth:", moons["Earth"])
	fmt.Println("Neptune:", moons["Neptune"])
}

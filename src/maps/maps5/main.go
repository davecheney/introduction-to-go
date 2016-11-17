package main

import "fmt"

func main() {
	moons := map[string]int{"Mercury": 0, "Venus": 0, "Earth": 1, "Mars": 2, "Jupiter": 67}

	n, found := moons["Earth"]
	fmt.Println("Earth:", n, "Found:", found)

	n, found = moons["Neptune"]
	fmt.Println("Neptune:", n, "Found:", found)
}

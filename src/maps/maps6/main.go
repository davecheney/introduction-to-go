package main

import "fmt"

func main() {
	moons := map[string]int{"Mercury": 0, "Venus": 0, "Earth": 1, "Mars": 2, "Jupiter": 67}

	const planet = "Mars"

	n, found := moons[planet]
	fmt.Println(planet, n, "Found:", found)

	delete(moons, planet)

	n, found = moons[planet]
	fmt.Println(planet, n, "Found:", found)
}

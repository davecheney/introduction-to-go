package main

import "fmt"

func main() {
	cities := map[string]int{
		"Tokyo": 33200000, "New York": 17800000,
		"Sau Paulo": 17700000, "Delhi": 14300000,
		"Moscow": 10500000,
	}

	for name, pop := range cities { // HL
		fmt.Println("City:", name, "Population", pop)
	}
}

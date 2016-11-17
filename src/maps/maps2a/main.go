package main

import (
	"fmt"
)

func main() {
	// Maps can have keys of almost any type.
	//
	// Create a map from the names of days, eg "Monday", "Tuesday", to their
	// numbers. 1 for Monday, 2 for Tuesday.
	daynumber := make(map[string]int)

	daynumber["Monday"] = 1
	daynumber["Tuesday"] = 2
	daynumber["Wednesday"] = 3
	daynumber["Thursday"] = 4
	daynumber["Friday"] = 5
	daynumber["Saturday"] = 6
	daynumber["Sunday"] = 7

	// The correct answer should print 3 6
	fmt.Println(daynumber["Wednesday"], daynumber["Saturday"])
}

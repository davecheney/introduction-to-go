package main

import "fmt"

// START OMIT
func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	for i, p := range primes {
		i++
		switch i {
		case 1:
			fmt.Println("The", i, "st prime is", p)
		case 2:
			fmt.Println("The", i, "nd prime is", p)
		case 3:
			fmt.Println("The", i, "rd prime is", p)
		default:
			fmt.Println("The", i, "th prime is", p)
		}
	}
}

// END OMIT

package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	for i, p := range primes { // HL
		fmt.Println("The", i, "th prime is", p)
	}
}

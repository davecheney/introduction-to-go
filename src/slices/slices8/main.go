package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11}
	fmt.Println(len(primes), primes)

	primes = append(primes, 13)
	fmt.Println(len(primes), primes)

	primes = append(primes, 17, 19, 23)
	fmt.Println(len(primes), primes)
}

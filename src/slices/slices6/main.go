package main

import "fmt"

func main() {
	primes := make([]int, 10)
	primes[0] = 2
	primes[1] = 3
	primes[2] = 5
	primes[3] = 7
	primes[4] = 11
	primes[5] = 13
	primes[6] = 17
	primes[7] = 19
	primes[8] = 23
	primes[9] = 29
	fmt.Println(primes)
}

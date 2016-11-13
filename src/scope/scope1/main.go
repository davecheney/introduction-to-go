package main

import "fmt"

func f(x int) { // HL
	for x := 0; x < 10; x++ { // HL
		fmt.Println(x)
	}
}

var x int // HL

func main() {
	var x = 200 // HL
	f(x)
}

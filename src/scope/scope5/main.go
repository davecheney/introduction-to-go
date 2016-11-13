package main

import "fmt"

// START1 OMIT
var x = 100 // HL

func main() {
	var x = 200 // HL
	fmt.Println(x)
}

// END1 OMIT

// START2 OMIT
func f() {
	var x = 99 // HL
	if x > 90 {
		x := 60 // HL
		fmt.Println(x)
	}
}

// END2 OMIT

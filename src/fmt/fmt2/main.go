package main

import "fmt"

// START OMIT
func main() {
	name := "David"
	age := 27

	fmt.Printf("Hello my name is %d, my age is %d", name)

	_ = age // keep the compiler happy
}

// END OMIT

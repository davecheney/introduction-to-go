package main

import "fmt"

func main() {
	x := 100
	for i := 0; i < 5; i++ {
		x := i // HL
		fmt.Println(x)
	}
	fmt.Println(x)
}

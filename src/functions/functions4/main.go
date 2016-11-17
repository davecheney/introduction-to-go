package main

import "fmt"

// Double returns a number two times larger than i.
func A(i int) int {
	return i * 2
}

// Swap switches the values of a and b
func Swap(a int, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(A(200))

	a, b := Swap(7, 9)
	fmt.Println(a, b)
}

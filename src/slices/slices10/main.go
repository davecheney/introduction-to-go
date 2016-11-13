package main

import "fmt"

func main() {
	a := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	fmt.Println(a)

	b := a[0:5]
	for i := 0; i < len(b); i++ {
		b[i] = b[i] * -1
	}
	fmt.Println(b)

	fmt.Println(a)
}

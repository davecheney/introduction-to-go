package main

import "fmt"

func f() {
	x := 200
	fmt.Println("inside f: x =", x)
}

func main() {
	x := 100
	fmt.Println("inside main: x =", x)
	f()
	fmt.Println("inside main: x =", x)
}

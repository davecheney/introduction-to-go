package main

import "fmt"

func f() {
	x := 200
	fmt.Println("inside f: x =", x)
}

var x = 50

func main() {
	fmt.Println("inside main: x =", x)
	x := 100
	fmt.Println("inside main: x =", x)
	f()
	fmt.Println("inside main: x =", x)
}

package main

import "fmt"

func main() {
	var π = 3.14159
	var radius = 6371.0 // radius of the Earth in km
	var circumference = 2 * π * radius

	fmt.Println("Circumference of the earth is", circumference, "km")
}

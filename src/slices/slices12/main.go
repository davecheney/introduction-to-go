package main

import "fmt"

func main() {
	brothers := []string{"chico", "harpo", "groucho", "gummo", "zeppo"}
	fmt.Println(brothers)

	wellknown := brothers[0:3]
	fmt.Println(wellknown)
}

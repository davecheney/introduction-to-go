package main

import "fmt"

// START OMIT
func main() {
	days := make(map[int]string)
	days[1] = "Monday"
	days[2] = "Tuesday"
	days[3] = "Wednesday"
	days[4] = "Thursday"
	days[5] = "Friday"
	days[6] = "Saturday"
	days[7] = "Sunday"
	fmt.Println(days)
}

// END OMIT

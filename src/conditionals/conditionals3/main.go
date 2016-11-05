package main

func main() {
	// START OMIT
	var i = 1
	for {
		// if i > 10 { // HL
		//	break // HL
		// } // HL
		if i%2 == 0 {
			println(i)
		}
		i++
	}
	// END OMIT
}

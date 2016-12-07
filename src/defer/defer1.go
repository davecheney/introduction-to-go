package count

import (
	"bufio"
	"log"
	"os"
)

// START OMIT
func CountLines(path string) int {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // HL

	sc := bufio.NewScanner(f)
	var lines int
	for sc.Scan() {
		lines++
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

// END OMIT

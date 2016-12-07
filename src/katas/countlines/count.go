package count

import (
	"bufio"
	"log"
	"os"
)

func CountLines(path string) int {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(f)
	var lines int
	for sc.Scan() {
		lines++
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	f.Close()
	return lines
}

package main

import (
	"bufio"
	"io"
	"log"
	"strings"
)

// ReadAll reads all the lines of text from r and returns
// all the data read as a string
func ReadAll(r io.Reader) string {
	sc := bufio.NewScanner(r)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if sc.Err() != nil {
		log.Fatal(sc.Err())
	}
	return strings.Join(lines, "\n")
}

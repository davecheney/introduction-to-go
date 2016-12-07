package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func CountLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	var lines int
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

func CountFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	lines, err := CountLines(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\t%d\n", path, lines)
}

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		CountFile(arg)
	}
}

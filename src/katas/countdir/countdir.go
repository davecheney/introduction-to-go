package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
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

func CountDir(dir string) {
	d, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()
	entries, err := d.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		if strings.HasSuffix(e.Name(), ".txt") {
			CountFile(filepath.Join(dir, e.Name()))
		}
	}
}

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		CountDir(arg)
	}
}

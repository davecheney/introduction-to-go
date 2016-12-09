package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// START OMIT
type Person struct {
	Name    string
	City    string
	Country string
}

func main() {
	p := Person{Name: "Dave", City: "Sydney", Country: "Australia"}
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b.String())
}

// END OMIT

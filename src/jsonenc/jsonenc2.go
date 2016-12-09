package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// START OMIT
type Person struct {
	Name    string `json:"name"`
	City    string `json:",omitempty"`
	Country string
}

func main() {
	p := Person{Name: "Dave", Country: "Australia"}
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b.String())
}

// END OMIT

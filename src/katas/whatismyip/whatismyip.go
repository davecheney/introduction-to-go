package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://httpbin.org/html")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	result := make(map[string]interface{})
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&result); err != nil {
		log.Fatal(err)
	}

	for key, value := range result {
		fmt.Println("Key:", key, "Value:", value)
	}
}

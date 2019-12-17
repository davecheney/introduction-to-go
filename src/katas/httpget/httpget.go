package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s $URL", os.Args[0])
	}

	fmt.Println("Args: ", os.Args[0], os.Args[1])

	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}

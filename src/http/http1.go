package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the index page")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	http.ListenAndServe(":8000", mux)
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func greet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "今日は %s", name)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/greet/{name}", greet)
	http.ListenAndServe(":8000", r)
}

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func greet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	firstname := vars["firstname"]
	lastname := vars["lastname"]
	fmt.Fprintf(w, "今日は %s %s", firstname, lastname)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/greet/{firstname}/{lastname}", greet)
	http.ListenAndServe(":8000", r)
}

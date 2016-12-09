package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

var booksdir string

type JSONResult struct {
	Title string `json:"title"`
	Lines int    `json:"lines,string"`
}

func bookController(w http.ResponseWriter, r *http.Request) {
	// get the name of the book
	vars := mux.Vars(r)
	book, ok := vars["book"]
	if !ok {
		http.Error(w, "no book specified", 404)
		return
	}

	// turn the name of the book into a file path
	path := filepath.Join(booksdir, book)

	// count the lines in the book
	lines, err := countLines(path)

	// if there was an error, report it back to the client and return
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// create our JSON result value
	res := &JSONResult{
		Title: book,
		Lines: lines,
	}

	// create a JSON encoded using w to write JSON back to the client
	enc := json.NewEncoder(w)
	enc.Encode(res)
}

func countLines(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var lines int
	for sc.Scan() {
		lines++
	}
	err = sc.Err()
	return lines, err
}

func main() {
	// check that os.Args[1] has the books directory
	if len(os.Args) < 2 {
		log.Fatal("usage: %s /path/to/books", os.Args[0])
	}

	// record the books directory
	booksdir = os.Args[1]

	// set up our router
	r := mux.NewRouter()

	// register /books/{book} handler
	r.HandleFunc("/books/{book}", bookController)

	// listen and serve for ever
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

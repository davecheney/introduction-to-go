package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const templ = `<html><head><title>Hello</title></head><body>
Hello {{ .RemoteAddr }}
You sent me a {{ .Method }} request for {{ .URL }}
</body></html>`

func HelloServer(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	t := template.Must(template.New("html").Parse(templ))
	t.Execute(w, req)
}

func main() {
	fmt.Println("please connect to localhost:7777/hello")
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":7777", nil))
}

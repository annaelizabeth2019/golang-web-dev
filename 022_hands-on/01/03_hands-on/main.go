package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	// accidentally did both 03 and 05 in this file
	http.Handle("/", http.HandlerFunc(foo))
	http.Handle("/dog/", http.HandlerFunc(bar))
	http.Handle("/me/", http.HandlerFunc(myName))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	err := tpl.Execute(w, req)
	if err != nil {
		log.Fatalln(err)
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bar ran")
}

func myName(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "anna")
}

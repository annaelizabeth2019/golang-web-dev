package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandleFunc(foo))
	http.Handle("/dog/", http.HandleFunc(bar))
	http.Handle("/me/", http.HandleFunc(anna))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func bar(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bar ran")
}

func anna(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "something.gohtml", "anna")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

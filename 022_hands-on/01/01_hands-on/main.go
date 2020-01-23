package main

import "net/http"

import "io"

func writeName(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Anna")
}



func main() {
	http.HandleFunc("/me/", writeName)

	http.ListenAndServe(":8080", nil)
}

package main

import (
	"net/http"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

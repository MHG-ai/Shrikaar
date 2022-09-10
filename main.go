package main

import (
	"net/http"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir("./forms")))
	http.Handle("/", http.FileServer(http.Dir("./")))
	// http.Handle("/home", http.FileServer(http.Dir("./index.html")))
	http.ListenAndServe(":80", nil)
}

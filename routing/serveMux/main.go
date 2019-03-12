package main

import (
	"io"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome Home!!! This is better!")
}

func about(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "About. Well Theres still alot we dont know about")
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	http.ListenAndServe(":8080", nil)
}

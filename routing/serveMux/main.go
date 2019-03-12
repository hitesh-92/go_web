package main

import (
	"io"
	"net/http"
)

type home int

func (h home) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome Home!!! This is better!")
}

type about int

func (a about) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "About. Well Theres still alot we dont know about")
}

func main() {
	var h home
	var a about

	mux := http.NewServeMux()
	mux.Handle("/", h)
	mux.Handle("/about", a)

	http.ListenAndServe(":8080", mux)
}

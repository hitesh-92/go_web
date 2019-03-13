package main

import (
	"io"
	"net/http"
)

func root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<p>Welcome. Check out <a href='/dog'>dog</a> and <a href='/me'>me</a> </p>")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<p>Woof.</p> <a href='/'>Home</p>")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<p>My name is Hitesh</p> <a href='/'>Home</p>")
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}

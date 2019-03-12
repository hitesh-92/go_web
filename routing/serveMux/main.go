package main

import (
	"io"
	"net/http"
)

type testing int

func (t testing) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	switch req.URL.Path {

	case "/":
		io.WriteString(w, "check out: /home or /about")

	case "/about":
		io.WriteString(w, "About Us. This is for routing!")

	case "/home":
		io.WriteString(w, "Welcome Home! Nice to be back?")

	}
}

func main() {
	var a testing
	http.ListenAndServe(":8080", a)
}

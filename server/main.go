package main

import (
	"fmt"
	"net/http"
)

type testing int

func (t testing) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Just testing for now")
}

func main() {
	var t testing
	http.ListenAndServe(":8080", t)
}

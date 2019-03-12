package main

import (
	"html/template"
	"log"
	"net/http"
)

type testing int

func (t testing) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()

	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var a testing
	http.ListenAndServe(":8080", a)
}

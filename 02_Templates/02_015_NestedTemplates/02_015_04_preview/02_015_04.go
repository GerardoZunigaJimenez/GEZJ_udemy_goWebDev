package main

import (
	"html/template"
	"net/http"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

/**
go run 02_015_04.go
localhost:8080
 */
func main() {

	http.HandleFunc("/", index)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

package main

import (
	"net/http"
	"log"
	"html/template"
	"fmt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

/*
Serve the files in the "starting-files" folder
To get your images to serve, use only this:
	fs := http.FileServer(http.Dir("public"))
Hint: look to see what type FileServer returns, then think it through.
 */
func main() {
	fs := http.FileServer(http.Dir("public"))
	fmt.Printf("fs type: %T\n", fs)
	http.Handle("/pics/", fs)
	http.HandleFunc("/", dogs)
	http.ListenAndServe(":8080", nil)
}

func dogs(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

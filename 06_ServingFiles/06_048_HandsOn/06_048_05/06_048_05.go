package main

import (
	"html/template"
	"net/http"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

/*
Serve the files in the "starting-files" folder
To get your images to serve, use:
	func StripPrefix(prefix string, h Handler) Handler
	func FileServer(root FileSystem) Handler
Constraint: you are not allowed to change the route being used for images in the template file
 */
func main() {
	http.HandleFunc("/", foo)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

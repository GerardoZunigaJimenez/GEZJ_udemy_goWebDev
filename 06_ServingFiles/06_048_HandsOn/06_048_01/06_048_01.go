package main

import (
	"net/http"
	"io"
	"html/template"
)

/*
ListenAndServe on port 8080 of localhost
For the default route "/" Have a func called "foo" which writes to the response "foo ran"

For the route "/dog/" Have a func called "dog" which parses a template called "dog.gohtml" and writes to the response "

This is from dog
" and also shows a picture of a dog when the template is executed.
Use "http.ServeFile" to serve the file "dog.jpeg"
 */

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

 func main(){
 	http.HandleFunc("/", foo)
 	http.HandleFunc("/dog", dog)
 	http.HandleFunc("/dog.jpg", dogPic)
 	http.ListenAndServe(":8080", nil)
 }

 func foo(w http.ResponseWriter, r *http.Request){
	 w.Header().Set("Content-Type", "text/html; charset=utf-8")
	 io.WriteString(w, `<strong>FOO RAN</strong><br>`)
 }

 func dog(w http.ResponseWriter, r *http.Request){
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
 }

 func dogPic(w http.ResponseWriter, r *http.Request){
 	http.ServeFile(w, r,"toby.jpg")
 }
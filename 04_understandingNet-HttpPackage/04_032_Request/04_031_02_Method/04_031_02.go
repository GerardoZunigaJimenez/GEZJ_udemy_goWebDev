package main

import (
	"net/http"
	"log"
	"html/template"
	"net/url"
)

var tpl *template.Template

type hotdog int

func init(){
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

// Within this example any hot dog element is a HTTP Handler, because we are implementing this method.
/*
localhost:8080
 */
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil{
		log.Fatalln(err)
	}
	data := struct{
		Method string
		Submissions url.Values
	}{
		r.Method,
		r.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main(){
	var d hotdog
	http.ListenAndServe(":8080",d)
}

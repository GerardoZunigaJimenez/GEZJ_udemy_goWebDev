package main

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
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
	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
	fmt.Fprintln(w,"Andy code you want in this func")
}

func main(){
	var d hotdog
	http.ListenAndServe(":8080",d)
}

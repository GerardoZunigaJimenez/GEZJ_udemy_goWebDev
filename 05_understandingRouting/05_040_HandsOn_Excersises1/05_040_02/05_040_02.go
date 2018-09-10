package main

import (
	"net/http"
	"html/template"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

/*
Take the previous program in the previous folder and change it so that:
a template is parsed and served
you pass data into the template
 */
func main() {

	http.HandleFunc("/", root)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}

func root(res http.ResponseWriter, req *http.Request) {
	parseError(res, tpl.ExecuteTemplate(res, "index.gohtml", nil) )
}

func dog(res http.ResponseWriter, req *http.Request) {
	parseError(res, tpl.ExecuteTemplate(res, "dog.gohtml", nil) )
}

func me(res http.ResponseWriter, req *http.Request) {
	parseError(res, tpl.ExecuteTemplate(res, "me.gohtml", "Geras") )
}

func parseError(w http.ResponseWriter, err error){
	if err != nil{
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		log.Fatalln(err)
	}
}
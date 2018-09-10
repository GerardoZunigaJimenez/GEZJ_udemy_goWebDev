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
Take the previous program and change it so that:
func main uses http.Handle instead of http.HandleFunc
Contstraint: Do not change anything outside of func main
 */
func main() {

	http.Handle("/", http.HandlerFunc(root))
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/me", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}

func root(res http.ResponseWriter, req *http.Request) {
	parseError(res, tpl.ExecuteTemplate(res, "index.gohtml", nil) )
}

func dog(res http.ResponseWriter, req *http.Request) {
	parseError(res, tpl.ExecuteTemplate(res, "dog.gohtml", nil) )
}

func me(res http.ResponseWriter, req *http.Request) {
	parseError(res, tpl.ExecuteTemplate(res, "me.gohtml", "Jeringas") )
}

func parseError(w http.ResponseWriter, err error){
	if err != nil{
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		log.Fatalln(err)
	}
}
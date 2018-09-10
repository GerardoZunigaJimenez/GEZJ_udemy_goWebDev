package main

import (
	"net/http"
	"io"
)


func d(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "dog dog dogggy")
}

func c(res http.ResponseWriter, req *http.Request){
	io.WriteString(res, "cat cat cat")
}

//using default server mux
func main(){

	http.Handle("/dog/",http.HandlerFunc(d))
	http.Handle("/cat/",http.HandlerFunc(c))


	http.ListenAndServe(":8080", nil)
}
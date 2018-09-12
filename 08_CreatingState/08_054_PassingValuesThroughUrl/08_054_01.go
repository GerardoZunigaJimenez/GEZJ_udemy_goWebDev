package main

import (
	"net/http"
	"io"
)

/*
visit this page:
http://localhost:8080/?q=dog
 */
func main(){
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func foo(w http.ResponseWriter, r *http.Request){
	v := r.FormValue("q")
	io.WriteString(w, "Do my search: "+v)
}


package main

import (
	"net/http"
	"fmt"
)

/*
ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default ServeMux:

"/" "/dog/" "/me/

Add a func for each of the routes.

Have the "/me/" route print out your name.
 */
func main(){

	http.HandleFunc("/", root)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/me", me)

	http.ListenAndServe(":8080", nil)
}

func root(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Root")
}

func dog(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "dog")
}

func me(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "my Name is Gerardo Zuniga")
}

package main

import (
	"net/http"
	"fmt"
)

type hotdog int


// Within this example any hot dog element is a HTTP Handler, because we are implementing this method.
/*
localhost:8080
 */
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Andy code you want in this func")
}

func main(){
	var d hotdog
	http.ListenAndServe(":8080",d)
}

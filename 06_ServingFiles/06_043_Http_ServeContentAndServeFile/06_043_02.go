package main

import (
	"io"
	"net/http"
)

func main() {
	//We are defying 2 url that we will consume like functions
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jgp", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jgp">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "toby.jpg")
}

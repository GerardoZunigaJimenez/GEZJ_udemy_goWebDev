package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	//We are defying 2 url that we will consume like functions
	http.HandleFunc("/", dog)
	http.HandleFunc("/whiteDog", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//This is a reference from other http url
	io.WriteString(w, `
	<img src="/whiteDog">
	`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
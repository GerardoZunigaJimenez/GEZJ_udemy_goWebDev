package main

import (
	"io"
	"net/http"
	"os"
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
	f, err := os.Open("toby.jpg")
	if err != nil {
		//To launch status codes
		http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err !=nil{
		//To launch status codes
		http.Error(w, "file not found", http.StatusNotFound)
	}
	http.ServeContent(w, req,f.Name(), fi.ModTime(), f)
}
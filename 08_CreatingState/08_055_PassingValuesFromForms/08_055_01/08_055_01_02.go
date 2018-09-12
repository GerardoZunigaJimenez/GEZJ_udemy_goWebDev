package main

import (
"net/http"
"io"
)

/*
Comparing to send the values through body or post

localhost:8080/body - body
localhost:8080/url -  url
 */
func main() {
	http.HandleFunc("/body", body)
	http.HandleFunc("/url", url)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

/*
post
body
*/
func body(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//Change Method type between GET/POST
	io.WriteString(w, `
	<form method="post">
	<input type="text" name="q">
	<input type="submit">
	</form>
	<br>`+ v)
}

/*
get
url
 */
func url(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//Change Method type between GET/POST
	io.WriteString(w, `
	<form method="get">
	<input type="text" name="q">
	<input type="submit">
	</form>
	<br>`+ v)
}


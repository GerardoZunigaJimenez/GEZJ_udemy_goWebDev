package main

import (
	"net/http"
	"io"
)

/*
Now we're going to look at passing values through form submission.
So in some ways some it's a form the values can get sent to the server in two ways

One way it could go with the request body the payload right of the request that the values that would be some being submitted
from that form can just be in the body of that request go into the server.

The other way we can get values from the users submission of a form is by sending those values through
the URL.
 */
func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

/*
post
body

get
url
 */
func foo(w http.ResponseWriter, req *http.Request) {
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

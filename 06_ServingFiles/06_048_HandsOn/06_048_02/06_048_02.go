package main

import (
	"net/http"
)

/*
Serve the files in the "starting-files" folder
Use "http.FileServer"
 */


 func main(){
	 http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
 }
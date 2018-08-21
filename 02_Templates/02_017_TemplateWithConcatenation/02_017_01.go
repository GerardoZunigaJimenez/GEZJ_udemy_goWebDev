package main

import "fmt"

func main() {
	name := "Todd McLeod"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<h1>` + name + `</h1>
	</body>
	</html>
	`

	fmt.Println(tpl)
	/*
	exec go run 02_017_01.go > index.html
	 */
}

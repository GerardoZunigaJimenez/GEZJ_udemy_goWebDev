package main

import (
	"fmt"
	"os"
	"log"
	"io"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	str := fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8">
			<title>Hello World!</title>
			</head>
			<h1>` + name + `</h1>
			</body>
			</html>
			` )

	nf, err := os.Create("index.html")
	if err != nil{
		log.Fatalln("error creatig file", err)
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))

	/*
	Execute at console:
		go run 02_007_03.go Gerassssss
	 */
}

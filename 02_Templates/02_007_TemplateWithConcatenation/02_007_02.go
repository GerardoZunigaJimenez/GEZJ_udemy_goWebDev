package main

import (
	"fmt"
	"os"
	"log"
	"io"
	"strings"
)

func main() {
	name := "Gerardo zia"

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
}

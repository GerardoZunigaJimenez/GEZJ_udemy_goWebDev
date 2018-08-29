package main

import (
	"text/template"
	"log"
	"os"
)

type Page struct {
	Title, Heading, Input string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	home := Page{
		Title:   "Nothing Escaped",
		Heading: "Nothing is escaped with text/template",
		Input:   `<script>alert("Yow!");</script>`,
	}
	err := tpl.Execute(os.Stdout, home)
	if err != nil {
		log.Fatalln(err)
	}
}

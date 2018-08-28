package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	type user struct{
		Name string
		Motto string
		Admin bool
	}

	u1:= user{
		Name : "Buddha",
		Motto : "The belief of no beliefs",
		Admin : false,
	}
	u2:= user{
		Name : "Gandhi",
		Motto : "Be the change",
		Admin : true,
	}
	u3:= user{
		Name : "",
		Motto : "Nobody",
		Admin : true,
	}
	users := []user{u1,u2,u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}

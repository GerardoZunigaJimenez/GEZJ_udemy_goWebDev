package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model string
	Doors int
}

type items struct{
	Wisdom []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	a := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	b := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	c := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	d := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	e := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	f := car{
		"Toyota",
		"Corolla",
		4,
	}

	g := car{
		"Ford",
			"F150",
			2,
	}

	sages := []sage{a, b, c, d, e}
	cars := []car{f,g}

	data := items{
		Wisdom:sages,
		Transport:cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
package main

import (
	"log"
	"os"
	"text/template"
	"fmt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml","tpl_KeyValue.gohtml"))
}

func main() {
	sages := map[string]string{
		"India":    "Ghandi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhamad",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("\n")
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl_KeyValue.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}

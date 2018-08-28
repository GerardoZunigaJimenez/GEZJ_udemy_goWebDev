package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

/**
the way to pass data to the template is with {{.}}
where:
	dot is the data that we are passing

If you're coming from other server side programming but you pass in one piece of data.
But fortunately we can pass in a piece of data which is aggregate data type.
Right.

So it's agrah aggregate data structure.
So we can pass in a struct which has a whole bunch of fields and those fields could be you know of values
of type map or type slice or of type another struct.
And so really we can pass in as much as we want but you only get a pass in one piece of data or data
structure and then you get access that when you get into your template.
 */
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}

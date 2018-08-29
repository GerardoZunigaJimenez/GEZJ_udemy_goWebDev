package main

import (
	"text/template"
	"os"
	"log"
)

const (
	Southern string = "Southern"
	Central  string = "Central"
	Northern string = "Northern"
)

type hotel struct {
	Name    string
	Address string
	City    city
	Zip     int32
	Region  string
}
type city struct {
	Name  string
	State state
}

type state struct {
	Name    string
	Country country
}

type country struct {
	Name string
}

var tpl *template.Template

func createHotelList() []hotel {
	return []hotel{
		hotel{Name: "Extended Day America", Address: "Laurel St 415", City: city{Name: "Carlsbad", State: state{Name: "California", Country: country{Name: "USA",}}}, Zip: 90200, Region: Southern},
		hotel{Name: "Holiday Inn", Address: "Carlsbad Boulevard 741", City: city{Name: "Carlsbad", State: state{Name: "California", Country: country{Name: "USA",}}}, Zip: 90200, Region: Central},
		hotel{Name: "Marriot", Address: "Carlsbad ST 302", City: city{Name: "Carlsbad", State: state{Name: "California", Country: country{Name: "USA",}}}, Zip: 90200, Region: Central},
		hotel{Name: "San Marcos Hotel", Address: "San Marcos Boulevard 78", City: city{Name: "Carlsbad", State: state{Name: "California", Country: country{Name: "USA",}}}, Zip: 90200, Region: Northern},
		hotel{Name: "America Inn", Address: "Any Place at Carlsbad", City: city{Name: "Carlsbad", State: state{Name: "California", Country: country{Name: "USA",}}}, Zip: 90200, Region: Southern},
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

/*
go run 02_018_03.go > index.html
*/
func main() {
	hotels := createHotelList()

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil{
		log.Fatalln(err)
	}
}

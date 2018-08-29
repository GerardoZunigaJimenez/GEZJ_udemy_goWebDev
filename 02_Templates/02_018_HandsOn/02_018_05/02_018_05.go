package main

import (
	"time"
	"text/template"
	"os"
	"log"
	"encoding/csv"
	"strconv"
)

type record struct {
	Date time.Time
	Open float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func parseCSV(absFilePath string) ([]record, error) {
	src, err := os.Open(absFilePath)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer src.Close()

	rdr := csv.NewReader(src)

	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var records = []record{}

	for i, row := range rows {
		if i == 0 {
			continue
		}
		date, err := time.Parse("yyyy-mm-dd", row[0])
		open, err := strconv.ParseFloat(row[1], 64)

		if err != nil{
			log.Fatalln(err)
			return nil,err
		}
		records = append(records, record{date, open})
	}
	return records,nil
}

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	//parse csv
	records, err := parseCSV("table.csv")
	if err != nil{
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, records)
	if err != nil{
		log.Fatalln(err)
	}
}

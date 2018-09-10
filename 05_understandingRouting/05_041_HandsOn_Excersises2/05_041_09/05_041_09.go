package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"io"
	"html/template"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type rest struct {
	methodType string
	url        string
	gohtml     string
}

var sRest = []rest{
	rest{"GET", "/", "get.gohtml"},
	rest{"GET","/apply","getApply.gohtml"},
	rest{"POST","/apply","postApply.gohtml"},
}
/*
Add code to WRITE to the connection.
 */
func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	cCounter := 0

	for {
		conn, err := l.Accept()
		cCounter++

		if err != nil {
			log.Println(err)
			continue
		}

		go func(con net.Conn, cCounter int) {
			defer conn.Close()
			var i int
			var reqR rest

			scanner := bufio.NewScanner(con)

			for scanner.Scan() {
				ln := scanner.Text()
				if ln == "" {
					// when ln is empty, header is done
					fmt.Println("")
					break
				}
				if i == 0 {
					// we're in REQUEST LINE
					xs := strings.Fields(ln)
					reqR.methodType = xs[0]
					reqR.url = xs[1]
				}
				fmt.Printf("\tThe Connection #%v saiys: '%v'\n", cCounter, ln)
				i++
			}

			io.WriteString(con, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(con, "Content-Length: %d\r\n")
			fmt.Fprint(con, "Content-Type: text/html\r\n")
			io.WriteString(con, "\r\n")
			swtichURL(con,reqR,sRest)
		}(conn, cCounter)
		// we never get here
		// we have an open stream connection
		// how does the above reader know when it's done?
		fmt.Println("New Connection Appears: ", cCounter)
	}
}

func swtichURL(w io.Writer, input rest, posibilities []rest) {
	for _, p := range posibilities {
		if p.methodType == input.methodType && p.url == input.url {
			tpl.ExecuteTemplate(w, p.gohtml, nil)
			fmt.Printf("\tURL found: %v - %v\n", p.methodType, p.url)
			return
		}
	}
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	fmt.Printf("\tINDEX because url not found: %v - %v\n", input.methodType, input.url)
}

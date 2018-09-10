package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	_, uri := request(conn)

	switch(uri){
	case "/hello":
		respond(conn,"hello little fishes")
	case "/bye":
		respond(conn,"I don't wanna go Mr Stark")
	case "/contact":
		respond(conn,"I don't wanna go Mr Stark")
	case "/greaterCommons":
		respond(conn,"LearningGO")
	default :
		respond(conn,"ERRORRRRRR")
	}
}

func request(conn net.Conn) (string, string) {
	i := 0
	var method, uri string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			method = strings.Fields(ln)[0]
			uri = strings.Fields(ln)[1]
			fmt.Println("\t***METHOD", method)
			fmt.Println("\t***URL", uri)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	fmt.Println("\n\n")
	return method, uri
}

func respond(conn net.Conn, response string) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title>`+response+`</title></head><body>
	<strong>`+response+`</strong><br>
	<a href="/hello">index</a><br>
	<a href="/bye">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/contact">apply</a><br>
	<form method="POST" action="/apply">
	<input type="submit" value="apply">
	</form>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func respondError(conn net.Conn, response string) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title>`+response+`</title></head><body>
	<strong>`+response+`</strong><br>
	<a href="/hello">index</a><br>
	<a href="/bye">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/greaterCommons">apply</a><br>
	<form method="POST" action="/apply">
	<input type="submit" value="apply">
	</form>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 404")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
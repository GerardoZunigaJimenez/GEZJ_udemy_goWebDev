package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"io"
	"strings"
)

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

			scanner := bufio.NewScanner(con)
			cln := 0
			var method, url string

			for scanner.Scan() {
				ln := scanner.Text()
				cln++
				if ln == "" {
					// when ln is empty, header is done
					fmt.Println("")
					break
				}
				if cln <= 2{
					switch cln{
					case 1:
						method = strings.Split(ln, " ")[0]
					case 2:
						url = strings.Split(ln, " ")[1]
					}
				}
				fmt.Printf("\tThe Connection #%v saiys: '%v'\n", cCounter, ln)
			}

			body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
			body += "\nThe method that you sent was: " + method
			body += "\nwith the URL: " + url
			io.WriteString(con, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(con, "Content-Length: %d\r\n", len(body))
			fmt.Fprint(con, "Content-Type: text/plain\r\n")
			io.WriteString(con, "\r\n")
			io.WriteString(con, body)
		}(conn, cCounter)
		// we never get here
		// we have an open stream connection
		// how does the above reader know when it's done?
		fmt.Println("New Connection Appears: ", cCounter)
	}
}

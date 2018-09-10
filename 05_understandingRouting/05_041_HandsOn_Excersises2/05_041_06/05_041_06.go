package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"io"
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
			for scanner.Scan() {
				ln := scanner.Text()
				if ln == "" {
					// when ln is empty, header is done
					fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
					break
				}
				fmt.Printf("\tThe Connection #%v saiys: '%v'\n", cCounter, ln)
			}

			body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
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

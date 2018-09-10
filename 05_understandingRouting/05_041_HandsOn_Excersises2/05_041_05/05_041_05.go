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


		go func(con net.Conn, cCounter int ) {
			defer conn.Close()

			scanner := bufio.NewScanner(con)
			for scanner.Scan() {
				ln := scanner.Text()
				if ln != "" {
					fmt.Printf("\tThe Connection #%v saiys: '%v'\n",cCounter,ln)
					io.WriteString(conn, fmt.Sprintf("\tHello Connection #%v, the message that you sent was '%v'\n",cCounter, ln))
				} else {
					fmt.Printf("The End is comming for Connection #%v\n",cCounter)
					break
				}
			}
		}(conn, cCounter)


		// we never get here
		// we have an open stream connection
		// how does the above reader know when it's done?
		fmt.Println("New Connection Appears: ", cCounter)
		//Message for the Connection
		io.WriteString(conn, fmt.Sprintf("I see your connection #%v\n",cCounter) )
	}
}
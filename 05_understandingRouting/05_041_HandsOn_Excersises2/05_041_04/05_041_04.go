package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"io"
)

/*
Building upon the code from the previous problem:
Extract the code you wrote to READ from the connection using bufio.NewScanner into its own function called "serve".

Pass the connection of type net.Conn as an argument into this function.

Add "go" in front of the call to "serve" to enable concurrency and multiple connections.
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
					fmt.Printf("\tThe Connection #%v saiys: %v\n",cCounter,ln)
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
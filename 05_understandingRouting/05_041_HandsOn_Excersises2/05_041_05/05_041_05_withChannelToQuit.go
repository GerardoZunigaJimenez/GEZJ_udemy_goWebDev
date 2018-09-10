package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"io"
)

var kill = "$KILL$"
var q = make(chan bool)
/*
It is having problems to kill all sons when a KiLL command arrived
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

		go handleConnection(conn, cCounter)

		fmt.Println("New Connection Appears: ", cCounter)
		//Message for the Connection
		io.WriteString(conn, fmt.Sprintf("I see your connection #%v.  If you want to close the server please send the following message: '%v'\n", cCounter, kill))

		//This code is used to kill the server from a client message
		select {
		case v, ok := <-q:
			if v && ok {
				close(q)
				l.Close()
				return
			}
		default:
			continue

		}
	}
	//This Code will never be Printed
	fmt.Println("Infinity War is over")
}

func handleConnection(con net.Conn, cCounter int) {

	defer con.Close()
	killServer := false
	scanner := bufio.NewScanner(con)

	for scanner.Scan() {
		ln := scanner.Text()
		if ln == kill {
			killServer = true
			break
		}
		if ln != "" {
			fmt.Printf("\tThe Connection #%v saiys: '%v'\n", cCounter, ln)
			io.WriteString(con, fmt.Sprintf("\tHello Connection #%v, the message that you sent was '%v'\n", cCounter, ln))
		} else {
			fmt.Printf("The End is comming for Connection #%v\n", cCounter)
			break
		}
	}

	if killServer {
		fmt.Printf("I don't wanna go Mr Stark!\n")
		q <- true
		return
	}
}

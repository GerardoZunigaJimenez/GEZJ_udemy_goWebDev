package main

import (
	"net"
	"log"
	"fmt"
)

/*
go run 03_022.go
go run 03_024_02.go
*/

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
}

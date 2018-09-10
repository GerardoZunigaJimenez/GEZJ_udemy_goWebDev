package main

import (
	"net"
	"log"
	"io/ioutil"
	"fmt"
)

/*
go run "03_021.go"
go run "03_024_01"
 */
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}
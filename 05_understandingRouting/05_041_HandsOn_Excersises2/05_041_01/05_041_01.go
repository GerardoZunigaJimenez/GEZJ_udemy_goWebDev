package main

import (
	"net"
	"log"
	"io"
)

/*
Create a basic server using TCP.
The server should use net.Listen to listen on port 8080.

Remember to close the listener using defer.

Remember that from the "net" package you first need to LISTEN, then you need to ACCEPT an incoming connection.

Now write a response back on the connection.

Use io.WriteString to write the response: I see you connected.

Remember to close the connection.

Once you have all of that working, run your TCP server and test it from telnet (telnet localhost 8080).
 */

 func main(){
 	l, err := net.Listen("tcp",":8080")
 	parseError(err)
	 defer l.Close()

 	for{
 		con, err := l.Accept()
 		parseError(err)

 		io.WriteString(con,"I see you connected")

 		con.Close()
	}
 }

 func parseError(err error){
 	if err != nil{
		log.Fatalln(err)
	}
 }



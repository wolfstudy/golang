package main

import (
	"net"
	"fmt"
	"time"
	"log"
)

//golang solution c10k question.
func handle(conn net.Conn) {
	fmt.Fprintf(conn, "%s", time.Now().String())
	conn.Close()
}

func main() {
	l, err := net.Listen("tcp", "8080")
	if err!=nil {
		log.Fatal(err)
	}

	for  {
		conn,err:= l.Accept()
		if err!=nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

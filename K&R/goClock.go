package main

import (
	"net"
	"io"
	"time"
	"log"
)

func main() {
	lister, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lister.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		handlerConn(conn)
	}
}

func handlerConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:44:32\n"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}

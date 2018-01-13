package main

import (
	"net"
	"fmt"
	"time"
)

//golang solution c10k question.
func handle (conn net.Conn) {
	fmt.Fprintf(conn, "%s", time.Now().String())
	conn.Close()
}

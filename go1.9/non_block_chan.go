package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("recevied message:", msg)
	default:
		fmt.Println("no received message")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send message:", msg)
	default:
		fmt.Println("no message send")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	case sin := <-signals:
		fmt.Println("received signals:", sin)
	default:
		fmt.Println("no activity")
	}
}

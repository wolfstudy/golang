package main

import "fmt"

func main() {
	message := make(chan string)

	go func() {
		message <- "ping"
	}()

	msg := <- message

	fmt.Println("chan mse is :", msg)
}

package main

import "fmt"

func hello()  {
	fmt.Println("hello heihei")
}

func main() {
	go hello()
	fmt.Println("main func")
}

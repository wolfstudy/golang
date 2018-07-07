package main

import (
	l "container/list"
	"time"
	"fmt"
)

func main() {
	list()
}

var name = "aaa"

func list() {
	names := l.New()
	t := time.Now()
	for i := 1; i <= 1000000; i++ {
		_ = names.PushFront(name)
	}
	fmt.Println("list: " + time.Now().Sub(t).String())
}

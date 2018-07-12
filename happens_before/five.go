package main

import "sync"

var l sync.Mutex
var aaaaa string

func ffff() {
	aaaaa = "hello, world"
	l.Unlock()
}

//使用lock的方式解决
func main() {
	l.Lock()
	go ffff()
	l.Lock()
	print(aaaaa)
}


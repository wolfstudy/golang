package main

import "sync"

var aaaaaa string
var once sync.Once

func setup() {
	aaaaaa = "hello, world"
}

func doprint() {
	once.Do(setup)
	print(aaaaaa)
}

func twoprint() {
	go doprint()
	go doprint()
}

func main() {
	twoprint()
}

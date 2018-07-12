package main

import (
	_ "net/http/pprof"
	"log"
	"net/http"
)

var a string

func f() {
	print(a)
}

func hello() {
	a = "hello, world"
	go f()
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	hello() //main 在hello调用之前已经返回

	// localhost:6060/debug/pprof
	select {
	}
}

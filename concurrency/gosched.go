package main

import (
	"sync"
	"runtime"
	"fmt"
	"go/src/strconv"
	"time"
)

var syn =sync.Mutex{}


func test1() {
	for i := 0; i < 10; i++ {
		syn.Lock()
		runtime.Gosched()
		fmt.Println("test1 print:" + strconv.Itoa(i))
		syn.Unlock()
	}
}

func test2() {
	time.Sleep(500 * time.Microsecond)
	for i := 0; i < 10; i++ {
		syn.Lock()
		fmt.Println("test2 print:" + strconv.Itoa(i))
		syn.Unlock()
	}
}

func main() {
	go test1()
	go test2()
	time.Sleep(4 * time.Second)
}

package main

import (
	"sync/atomic"
	"time"
	"runtime"
)

//GODEBUG=schedtrace=1000,scheddetail=1 ./pmg
//GOMAXPROCS=1 GODEBUG=schedtrace=1000 ./pmg

var count int64

func test() {
	atomic.AddInt64(&count, 1)
	defer atomic.AddInt64(&count, -1)


	runtime.LockOSThread()
	defer runtime.LockOSThread()

	time.Sleep(time.Second)//1s
}

func main() {
	for i := 0; i <= 2000; i++ {
		go test()
	}

	for {
		time.Sleep(time.Second)
	}
}

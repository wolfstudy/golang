package main

import (
	"time"
	"fmt"
)

/*
	通过goroutines，channel和ticker都可以优雅地支持速率限制。
 */

func main() {

	//限制我们传入的请求
	request := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		request <- i
	}
	close(request)

	limiter := time.Tick(time.Millisecond * 200)

	//阻塞接收  1 request every 200 milliseconds.
	for req := range request {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//burstyLimiter通道将允许最多3个事件的突发
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

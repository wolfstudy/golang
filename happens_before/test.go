package main

import (
	"sync"
	"time"
)

//这里把buffered channel作为semaphore来使用，表面上看最多允许一个goroutine对count进行++和--
//根据Go语言的内存模型，对count变量的访问并没有形成临界区。
func main() {
	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ch <- true
			count++
			time.Sleep(time.Millisecond)
			count--
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
}

package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup
	go func() {
		time.Sleep(3 * time.Second)
		close(jobs)//panic: send on closed channel
	}()

	go func() {
		for i := 0; ; i++ {
			jobs <- i
			fmt.Println("produce:%d", i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range jobs {
			fmt.Println("consume:%d",i)
		}
	}()
	wg.Wait()
}

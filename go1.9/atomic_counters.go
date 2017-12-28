package main

import (
	"sync/atomic"
	"time"
	"fmt"
)

//using the sync/atomic package for atomic counters accessed by multiple goroutines.
func main() {
	var ops uint64 = 0

	for i := 0; i <= 50 ; i++ {
		go func() {
			for  {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

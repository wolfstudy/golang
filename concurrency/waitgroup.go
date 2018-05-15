package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup

func doSome(i int) {
	fmt.Println(i)
	wg.Done()
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go doSome(i)
	}
	wg.Wait()
}

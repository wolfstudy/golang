package main

import (
	"sync"
	"fmt"
	"time"
)

var once sync.Once

func main() {
	for i, v := range make([]string, 3) {
		once.Do(onces)
		fmt.Println("count", v, "----", i)
	}
	for i := 0; i < 3; i++ {
		go func() {
			once.Do(onced)
			fmt.Println("666")
		}()
	}

	time.Sleep(4*1000)
}

func onces() {
	fmt.Println("onces")
}

func onced() {
	fmt.Println("onced")
}
//整个程序，只会执行onces()方法一次,onced()方法是不会被执行的。

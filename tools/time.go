package main

import (
	"time"
	"fmt"
)


//计算时间，主要注意time.Since（）的用法
func main() {
	StartCac()
}

func StartCac() {
	t1 := time.Now() // get current time
	//logic handlers
	for i := 0; i < 1000; i++ {
		fmt.Print("*")
	}
	elapsed := time.Since(t1)
	fmt.Println()
	fmt.Println("App elapsed: ", elapsed)
}
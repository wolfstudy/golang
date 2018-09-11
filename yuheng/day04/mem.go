package main

import (
	"fmt"
	"time"
)

func main() {

	var x [1 << 20 * 100]byte

	for i := 0; i < len(x); i++ {
		x[i] = byte(i % 255)

		if i%(1<<20) == 0 {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}
}

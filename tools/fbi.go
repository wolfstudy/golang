package main

import "fmt"

//golang 实现斐波那契数列 感受 golang 的优美
func fib(c chan int, n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
		c <- a
	}

	close(c)
}

func main() {
	c := make(chan int)
	go fib(c, 10)

	for x := range c {
		fmt.Println(x)
	}
}

package main

func fib (c chan int, n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
		c <- a
	}

	close(c)
}

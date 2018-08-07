package main

import "fmt"

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	d := fib()()
	// Function calls are evaluated left-to-right.
	fmt.Println(f(), f(), f(), f(), f(), f())
	fmt.Println(d, d, d, d, d)
	fmt.Println(fib()(), fib()(), fib()(), fib()(), fib()())
}

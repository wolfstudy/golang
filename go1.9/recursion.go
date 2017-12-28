package main

import "fmt"

//递归
func fact (n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}

func main() {
	a := fact(6)
	fmt.Println(a)
}

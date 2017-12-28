package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func pplus(a int, b int, c int) int {
	return a + b + c
}

func main() {
	sum := plus(1,2)
	fmt.Println("sums:",sum)

	sums := pplus(1,2,3)
	fmt.Println("sums:",sums)

}

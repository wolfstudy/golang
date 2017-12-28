package main

import "fmt"

func val() (int,int)  {
	return 3, 8
}
func main() {
	a, b := val()
	fmt.Println("r1:",a)
	fmt.Println("r2:",b)

	_, c := val()
	fmt.Println("r3:",c)
}

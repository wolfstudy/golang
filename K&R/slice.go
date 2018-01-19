package main

import "fmt"

func main() {
	s := []int{5, 6, 7, 8, 9}

	//为什么 比较二者的区别
	s = remove(s, 2)
	fmt.Println(s)
	//[5 6 8 9 9]
	res := remove(s, 2)
	fmt.Println(res)
	//[5 6 8 9]
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

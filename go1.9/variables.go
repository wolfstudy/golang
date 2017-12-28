package main

import "fmt"

func main() {
	var a string = "wolf4j"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	//整数的初始值为0
	var e int
	fmt.Println(e)

	//另一种初始化的方式，一般开发中使用较多
	f := "bitmain"
	fmt.Println(f)
}

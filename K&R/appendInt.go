package main

import "fmt"

//appendInt 一次只能追加一个元素

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d,cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

//追加一个元素
func appendInt(x []int, y int) []int {
	var z []int

	zLen := len(x) + 1
	if zLen <= cap(x) {
		z = x[:zLen]
	} else {
		zCap := zLen
		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}
		z = make([]int, zLen, zCap)
		copy(z, x)
	}

	z[len(x)] = y

	return z
}

//拓展 ...int表示接收变长的slice
func appendInt1(x []int, y ...int) []int {
	var z []int

	//fix
	zLen := len(x) + len(y)

	if zLen <= cap(x) {
		z = x[:zLen]
	} else {
		zCap := zLen
		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}
		z = make([]int, zLen, zCap)
		copy(z, x)
	}

	//fix
	z[len(x):] = y

	return z
}

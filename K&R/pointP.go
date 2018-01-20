package main

import "fmt"

func main() {

	x1 := 1
	p := &x1
	fmt.Println(*p)
	fmt.Println(p) // 打印的是否是内存地址，还是一个hashcode

	*p = 2
	fmt.Println(x1)

	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)

	//icpr 需要记住一个点，虽然是指针自增，但是增加的是这个指针指向的值，这个指针不自己跑
	v := 1
	res := icpr(&v)
	fmt.Println("===============icpr============\n",res)

}

//go是基于c语言的，所以可以访问到内存地址，不像java只能访问到hashcode，需要注意

func icpr(p *int) int {
	*p++
	return *p
}

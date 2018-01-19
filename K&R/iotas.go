package main

import "fmt"

//iota 只能在常量（const）中使用
func main() {
	//const (
	//	a = iota
	//	b
	//	c
	//	d = iota
	//	e
	//	_
	//	_
	//	f
	//)
	//
	//fmt.Println("a = %v", a)
	//fmt.Println("b = %v", b)
	//fmt.Println("c = %v", c)
	//fmt.Println("d = %v", d)
	//fmt.Println("e = %v", e)
	//fmt.Println("f = %v", f)

	const (
		_   = 1 << (10 * iota)
		KiB
		MiB
		GiB
		TiB
		PiB
		EiB
		//ZiB
		//YiB
	)

	fmt.Println(KiB)
	fmt.Println(MiB)
	fmt.Println(GiB)
	fmt.Println(TiB)
	fmt.Println(PiB)
	fmt.Println(EiB)
	//fmt.Println(ZiB)
	//fmt.Println(YiB)

}

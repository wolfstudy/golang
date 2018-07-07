package main

import (
	"runtime"
	"fmt"
	"time"
)

//func A() *int {
//	x := 100
//	return &x
//}
//
//func B() [3]int {
//	return [3]int{1, 2, 3}
//}
//
//func C() []int {
//	return []int{1, 2, 3}
//}

//func D() {
//	x := make([]int, 3)
//	x[0] = 0x100
//	x[1] = 0x200
//	x[2] = 0x300
//	println(x[2])
//}

func E(x *int) {
	*x = 123

	runtime.SetFinalizer(x, func(*int) { println("finalizer") })
	runtime.GC()
	//runtime.KeepAlive(x)
	for {
		time.Sleep(time.Second)
		fmt.Println("aaa")
	}
}

func main() {
	//D()

	x := 100
	E(&x)
	fmt.Scanln()
}

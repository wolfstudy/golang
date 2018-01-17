package main

import (
	"fmt"
	"unsafe"
)

//Sizeof函数可以返回一个类型所占用的内存大小，这个大小只有类型有关，和类型对应的变量存储的内容大小无关，
func main() {
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(int8(0)))
	fmt.Println(unsafe.Sizeof(int16(10)))
	fmt.Println(unsafe.Sizeof(int32(10000000)))
	fmt.Println(unsafe.Sizeof(int64(10000000000000)))
	fmt.Println(unsafe.Sizeof(int(10000000000000000)))
}

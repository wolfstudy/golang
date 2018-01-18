package main

import (
	"fmt"
	"unsafe"
)

//内存对齐  Mac64位
type user1 struct {
	b byte
	i int32
	j int64
}
type user2 struct {
	b byte
	j int64
	i int32
}
type user3 struct {
	i int32
	b byte
	j int64
}
type user4 struct {
	i int32
	j int64
	b byte
}
type user5 struct {
	j int64
	b byte
	i int32
}
type user6 struct {
	j int64
	i int32
	b byte
}

func main() {
	var u1 user1
	var u2 user2
	var u3 user3
	var u4 user4
	var u5 user5
	var u6 user6
	fmt.Println("u1 size is ",unsafe.Sizeof(u1))
	fmt.Println("u2 size is ",unsafe.Sizeof(u2))
	fmt.Println("u3 size is ",unsafe.Sizeof(u3))
	fmt.Println("u4 size is ",unsafe.Sizeof(u4))
	fmt.Println("u5 size is ",unsafe.Sizeof(u5))
	fmt.Println("u6 size is ",unsafe.Sizeof(u6))
}

/*
u1 size is  16
u2 size is  24
u3 size is  16
u4 size is  24
u5 size is  16
u6 size is  16

这说明：

1，内存对齐影响struct的大小
2，struct的字段顺序影响struct的大小
综合以上两点，我们可以得知，不同的字段顺序，最终决定struct的内存大小，所以有时候合理的字段顺序可以减少内存的开销。
 */
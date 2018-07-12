package main

import (
	"container/list"
	"fmt"
)

//其根本原因是e是一个Element类型的指针，当然其也可能为nil，
// 但是golang中list包中函数没有对其进行是否为nil的检查，变默认其非nil进行操作，所以这种情况下，便可能出现程序崩溃。

func main() {
	li := list.New()

	li.PushBack(1)
	fmt.Println(li.Front().Value)//1

	value := li.Remove(li.Front())//1
	fmt.Println(value)

	value1 := li.Remove(li.Front())//panic: runtime error: invalid memory address or nil pointer dereference
	fmt.Println(value1)
}

package main

import (
	"fmt"
	"sync"
)

/*
	竞态条件:其中一个并发操作尝试读取变量，在这期间的某一时刻，另一个并发操作尝试同时写入相同的变量。
	在golang语言中，我们可以使用一个go的关键字来运行一个并发的函数，我们一般叫做gorutine
 */

func main() { //main函数本身就是一个主线程，
	var memoryAccess sync.Mutex
	var data int

	//会另起一个线程
	go func() {
		memoryAccess.Lock()
		data++ //尝试访问变量 A操作
		memoryAccess.Unlock()
	}()

	//time.Sleep(100000) // 这种操作强烈不建议
	memoryAccess.Lock()
	if data == 0 { //尝试访问变量 B操作
		fmt.Println("The value is 0") //C 操作
	} else {
		fmt.Printf("the value is %d", data)
	}
	memoryAccess.Unlock()
}

//运行结果可能有三种：
//1. The value is 0；B操作在A操作之前
//2. 什么也不打印；A操作在B操作之前
//3. The value is 1；B 操作在A 操作之前，但是A 操作在C 操作之前

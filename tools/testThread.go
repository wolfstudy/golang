package main

import "fmt"

func hello()  {
	fmt.Println("hello heihei")
}

func main() {
	go hello()
	fmt.Println("main func")
}

/**
该程序只会输出文本 main function。我们启动的 Go 协程究竟出现了什么问题？
1.启动一个新的协程时，协程的调用会立即返回。
与函数不同，程序控制不会去等待 Go 协程执行完毕。
在调用 Go 协程之后，程序控制会立即返回到代码的下一行，忽略该协程的任何返回值。
 */
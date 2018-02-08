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
 */
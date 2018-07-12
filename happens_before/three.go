package main

var c = make(chan int, 10)
var aaa string

func ff() {
	aaa = "hello, world"
	c <- 0
}

//要让gorutine显示执行的话，使用channel作为同步信号，阻塞
//The closing of a channel happens before a receive that returns a zero value because the channel is closed.
func main() {
	go ff()
	<-c
	print(aaa)
}

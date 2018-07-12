package main

//as above, but with the send and receive statements swapped and using an unbuffered channel

//var cc = make(chan int)

//If the channel were buffered (e.g., c = make(chan int, 1))
// then the program would not be guaranteed to print "hello, world".
var cc = make(chan int, 1)
var aaaa string

func fff() {
	aaaa = "hello, world"
	<-cc
}

//is also guaranteed to print "hello, world".
func main() {
	go fff()
	cc <- 0
	print(aaaa)
}

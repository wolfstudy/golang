package main

var aa string

func hello1() {
	go func() { aa = "hello" }()
	print(aa)
}

//If the effects of a goroutine must be observed by another goroutine,
// use a synchronization mechanism such as a lock or channel communication
// to establish a relative ordering.

//refer to three.go
func main() {
	hello1()
}


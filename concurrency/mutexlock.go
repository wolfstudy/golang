package main

import (
	"sync"
	"fmt"
)

//all goroutines are asleep - deadlock!
func main() {
	var l *sync.Mutex
	l = new(sync.Mutex)
	l.Lock()
	fmt.Println("1")
	l.Lock()
}

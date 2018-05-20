package main

import (
	"sync/atomic"
	"fmt"
)

//使用另一种方式实现sync.once
type Onces struct {
	done int32
}
func (o *Onces) Do(f func()) {
	if atomic.LoadInt32(&o.done) == 1 {
		return
	}
	// Slow-path.
	if atomic.CompareAndSwapInt32(&o.done, 0, 1) {
		f()
	}
}

func main() {
	var once1 Onces
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once1.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

//var once1 Onces
//func main() {
//	for i, v := range make([]string, 10) {
//		once1.Do(oncess)
//		fmt.Println("count", v, "---->", i)
//	}
//	for i := 0; i < 10; i++ {
//		go func() {
//			once1.Do(onceds)
//			fmt.Println("666")
//		}()
//	}
//
//	time.Sleep(40000)
//}
//
//func oncess() {
//	fmt.Println("onces")
//}
//
//func onceds() {
//	fmt.Println("onced")
//}

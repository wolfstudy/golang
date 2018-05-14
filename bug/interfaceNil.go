package main

import "fmt"

/*
Nil接口与接口中有一个零指针不同。
 */

type Cat interface {
	Meow()
}

type Tabby struct{}

func (*Tabby) Meow() {
	fmt.Println("meow")
}

func GetACat() Cat {
	var myTabby *Tabby = nil
	// Oops, we forgot to set myTabby to a real value
	return myTabby
}

func main() {
	if GetACat() == nil {
		fmt.Errorf("Forgot to return a real cat")
	}
}

/*
Guess what? 上面的测试不会检测到零指针！这是因为接口用作指针，所以GetACat有效地返回一个指向零指针的指针。千万不要做GetACat做的操作。
使用错误值很容易做到这一点。
 */

package main

import (
	"errors"
	"fmt"
)

var (
	ErrDidNotWork = errors.New("did not work")
)

func testSome(reallyDoIt bool) (err error) {
	//var result string
	if reallyDoIt {
		result, err := trytheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

//函数返回值中（err error）已经定义err，但是在函数内部，又重新定义err，返回值中的err会遮蔽 := 定义的err。

func trytheThing() (string, error) {
	return "hahah", errors.New("hello world")
}

func main() {
	b := true
	e := testSome(b)
	fmt.Println(e) //nil
}

//上面的函数总是会返回一个nil错误，因为内部err变量“shadows”函数范围变量。
//fix:
//使用var result string或者不要使用 :=

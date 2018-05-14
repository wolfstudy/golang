package main

import "fmt"

func test(b *bool) bool {
	*b = false
	return *b
}

func main() {
	bb := false
	v := test(&bb)
	fmt.Printf("======%v", v)
}

/*

上述代码主要测试一个问题：关于参数是值传递还是引用传递，如果是值传递的话，是没有办法修改那个值的，
因为在参数调用时，会把这个值copy一份过来，此时他的值是改变了，但是他们的地址是不一样的，所以，并没有真正修改那个内部的值
但是传指针进去，就不会出现上述问题，虽然在函数调用的时候也会copy一份，但是他们前后的地址是一样的，他们指向同一块内存。
所以这个时候修改之后的值，会真正的修改。
 */

package main

import (
	"reflect"
	"fmt"
)

//反射 typeOF
//在Go的反射定义中，任何接口都会由两部分组成的，一个是接口的具体类型，一个是具体类型对应的值
//标准库为我们提供两种类型来分别表示他们reflect.Value和reflect.Type，
// 并且提供了两个函数来获取任意对象的Value和Type。


//reflect.TypeOf可以获取任意对象的具体类型，
func main() {
	u := User{"make",20}
	t := reflect.TypeOf(u)
	fmt.Println(t)
}

type User struct {
	Name string
	Age  int
}

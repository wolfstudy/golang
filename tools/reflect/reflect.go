package main

import (
	"reflect"
	"fmt"
)

//反射 typeOF
//在Go的反射定义中，任何接口都会由两部分组成的，一个是接口的具体类型，一个是具体类型对应的值
//标准库为我们提供两种类型来分别表示他们reflect.Value和reflect.Type，
// 并且提供了两个函数来获取任意对象的Value和Type。



func main() {
	u := User{"make",20}

	//reflect.TypeOf可以获取任意对象的具体类型，
	t := reflect.TypeOf(u)
	fmt.Println(t)

	//reflect.TypeOf函数接受一个空接口interface{}作为参数，所以这个方法可以接受任何类型的对象。
	v := reflect.ValueOf(u)
	fmt.Println(v)

}

type User struct {
	Name string
	Age  int
}

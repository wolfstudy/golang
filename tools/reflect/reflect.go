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
	u := User{"make", 20}

	//通过reflect.ValueOf函数把任意类型的对象转为一个reflect.Value

	//reflect.TypeOf可以获取任意对象的具体类型，
	t := reflect.TypeOf(u)
	fmt.Println(t)

	//reflect.TypeOf函数接受一个空接口interface{}作为参数，所以这个方法可以接受任何类型的对象。
	v := reflect.ValueOf(u)
	fmt.Println(v)

	//将上述过程逆转回来

	//reflect.Value为我们提供了Inteface方法来帮我们做这个事情
	u1 := v.Interface().(User)
	fmt.Println(u1)

	//这里可以还原的原因是因为在Go的反射中，把任意一个对象分为reflect.Value和reflect.Type，
	// 而reflect.Value又同时持有一个对象的reflect.Value和reflect.Type,
	// 所以我们可以通过reflect.Value的Interface方法实现还原。

	//如何从一个reflect.Value获取对应的reflect.Type。
	t1 := v.Type()
	fmt.Println(t1)
}

type User struct {
	Name string
	Age  int
}

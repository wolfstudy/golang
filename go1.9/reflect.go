package main

import (
	"reflect"
	"fmt"
)

func main() {
	var num = 1

	v := reflect.ValueOf(num)
	fmt.Println("num value is:", v)
	pv := reflect.ValueOf(&num)
	fmt.Println("num value addr is:", pv)

	t := reflect.TypeOf(num)
	fmt.Println("num type is: ", t)

	type User struct {
		Id   int    `default:"12121"`
		Name string `default:"wolf"`
		Age  int    `default:"12"`
	}

	//u := &User{
	//	Id:   1212,
	//	Name: "wolf",
	//	Age:  12,
	//}
	getType := reflect.TypeOf(User{})
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(User{})
	fmt.Println("get all Fields is:", getValue)

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i)
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

}

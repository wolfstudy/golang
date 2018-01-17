package main

import (
	"reflect"
	"fmt"
)

func main() {
	u := User{"make",20}
	t := reflect.TypeOf(u)
	fmt.Println(t)
}

type User struct {
	Name string
	Age  int
}

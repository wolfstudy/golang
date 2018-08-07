package main

import "fmt"

func main() {
	fmt.Println(GetUser())
}

func GetUser() *user {
	u := &user{
		name: "aaa",
		age:  12,
	}
	return u
}

type user struct {
	name string
	age  int
}

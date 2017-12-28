package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"wolf4j",23})
	fmt.Println(person{name:"make",age:21})
	fmt.Println(person{name:"john"})
	fmt.Println(&person{name:"mari",age:21})

	s := person{
		name:"Ann",
		age:32,
	}
	fmt.Println("name is:", s.name)

	sp := &s
	fmt.Println("age is:", sp.age)

	sp.age = 41
	fmt.Println("reSet age is:", sp.age)
}

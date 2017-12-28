package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 8

	fmt.Println("map:",m)

	v1 := m["k1"]
	fmt.Println("value:",v1)
	fmt.Println("length:",len(m))

	delete(m,"k2")
	fmt.Println("map:",m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

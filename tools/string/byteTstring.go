package main

import "fmt"

func main() {
	c := [100]byte{'a','b','c'}
	fmt.Println("C values is :",len(c), c[:4])
	g := CToGoString(c[:4])
	fmt.Println("G values is :",len(g), g)
}

func CToGoString (c []byte) string {
	n := -1
	for index, value := range c {
		if value == 0 {
			break
		}
		n = index
	}
	return string(c[:n+1])
}

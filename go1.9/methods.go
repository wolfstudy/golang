package main

import "fmt"

type rect struct {
	width   int
	height  int
}

func (r *rect) area () int {
	return r.height * r.width
}

func (r rect) perim () int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width:10,height:20}

	fmt.Println("The area is:", r.area())
	fmt.Println("The perim is:",r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("priem:",rp.perim())
}

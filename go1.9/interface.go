package main

import (
	"math"
	"fmt"
)

type geometry interface {
	area()  float64
	perim() float64
}

type rects struct {
	width   float64
	height  float64
}

type circle struct {
	radius float64
}

func (r rects) area() float64 {
	return r.height * r.width
}

func (r rects) perim() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rects{width:10,height:2}
	c := circle{radius:12}

	measure(r)
	measure(c)
}

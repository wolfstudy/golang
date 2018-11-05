package main

import (
	"github.com/davecgh/go-spew/spew"
	"fmt"
)

type People struct {
	Student
}

type Student struct {
	name string
	age  int
	sex  bool
	tea  *Teach //数值不会被拷贝
}

type Teach struct {
	name string
	age  int
}

func main() {
	s := &Student{
		name: "wolf",
		age:  21,
		sex:  true,
		tea: &Teach{
			name: "yujian",
			age:  12,
		},
	}

	p := &People{*s}
	spew.Dump(p)

	// tmp copy
	newP := *p
	newP.name = "qiwei"
	newP.age = 40
	newP.tea.age = 31 //如果不是基础类型 浅拷贝会污染原来的结构，需要注意

	spew.Dump(newP)
	fmt.Println("浅copy之后")
	spew.Dump(p)


	fmt.Println("分割线")
	newp := &People{*s}

	news := newp.deepCopy()
	news.tea.age = 31
	news.tea.name = "shabi"

	fmt.Println("deep copy before")
	spew.Dump(newp)
	fmt.Println("tmp fix value")
	spew.Dump(news)
	fmt.Println("deep copy after")
	spew.Dump(newp)

}

func (p *People) deepCopy() *Student {
	s := &Student{name: p.name, age: p.age, sex: p.sex}
	newtea := &Teach{name: p.tea.name, age: p.tea.age}
	s.tea = newtea
	return s
}

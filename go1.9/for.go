package main

import "fmt"

func main() {

	fmt.Println("循环的第一种方式：\n")
	i := 1
	for i <= 3  {
		fmt.Println(i)
		i++
	}

	fmt.Println("循环的第二种方式：\n")
	for j :=7 ; j <= 9 ; j++ {
		fmt.Println(j)
	}

	fmt.Println("循环的第三种方式：\n")
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("循环的第四种方式：\n")
	for n := 0 ; n <= 5 ; n++ {
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}

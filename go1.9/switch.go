package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("write", i ,"is")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Tuesday, time.Wednesday:
		fmt.Println("today is weekend")
	default:
		fmt.Println("today is weekday")
		
	}

	t := time.Now()
	switch  {
	case t.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}

	whatFuck := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i am boolean")
		case int:
			fmt.Println("i am int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatFuck(true)
	whatFuck(1)
	whatFuck("heheda")
}

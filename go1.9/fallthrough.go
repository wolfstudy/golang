package main

import "fmt"

/*
   fallthrough 语句是不退出当前case 语段去实行下一个case 语段，且不去判断case 条件的去实行。
 */

func main() {
	switch a := 3; {
	case a >= 2:
		fmt.Println(">=2")
		fallthrough
	case a >= 3:
		fmt.Println(">=3")
		fallthrough
	case a >= 4:
		fmt.Println(">=4")
		fallthrough
	case a >= 5:
		fmt.Println(">=5")
		fallthrough
	default:
		fmt.Println("default")
	}
}
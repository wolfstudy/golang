package main

import "fmt"

//The following code is looking for '\0',
func main() {
	b := []byte{1, 2, 3, 0, 0, 0}
	fmt.Println(FirstZero(b))
}

func FirstZero(b []byte) int {
	min, max := 0, len(b)
	for  {
		if min+1 == max {
			return max
		}
		mid := (min + max) / 2
		if b[mid] == '\000' {
			max = mid
		}else {
			min = mid
		}
	}
	return len(b)
}

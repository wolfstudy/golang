package main

import "fmt"

/*
 	一位一位的反转
 */
func main() {
	s := "e4d1b71e6ce59f1af3876a730558df2c3b99257eb538e2f6f283379df95f868401000000"
	fmt.Println(len(s))
	fmt.Println("string: ", reverseString(s))
}

func reverseString (s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes) - 1; from < to ; from , to = from + 1, to - 1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

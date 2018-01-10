package main

import "fmt"

/*
	小端到大端的转化 16进制，两位一起。btc 的大小端工具 //by wolf4j
 */
func main() {
	s := "3a1c430e5c8719ef66525461a374d33849d60c07c5d198b16bdf788e3e046de6"
	fmt.Println("length:", len(s))
	fmt.Println("string:", revertString(s))
}

func revertString(s string) string {
	len := len(s)

	result := make([]byte, 0, len)

	for i := len - 2; i >= 0; i-- {
		if i%2 == 0 {
			result = append(result, s[i], s[i+1])
		}
	}
	return string(result)
}

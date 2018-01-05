package main

import "fmt"

/*
	小端到大端的转化 16进制，两位一起。
 */
func main() {
	s := "dad4b83b82ea8a4e6100ebb9eff81bc25364c3531d9033afb2a35433c91dfc28"
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

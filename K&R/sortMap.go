package main

import (
	"sort"
	"fmt"
)

//为了提高map的顺序，map是无序的。
func main() {
	m := make(map[string]string)

	m["hello"] = "echo hello"
	m["hi"] = "echo hi"
	m["haha"] = "echo haha"
	m["hehe"] = "echo hehe"
	m["heihei"] = "echo heihei"
	m["xixi"] = "echo xixi"

	sortKeys := make([]string, 0)
	for index, _ := range m {
		sortKeys = append(sortKeys, index)
	}

	sort.Strings(sortKeys)

	//for k, v := range m {
	//	fmt.Printf("key is:%s, value is:%s \n", k, v)
	//}

	for _, keys := range sortKeys {
		fmt.Printf("k=%v\t v=%v\n", keys, m[keys])
	}
}

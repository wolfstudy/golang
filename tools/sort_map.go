package main

import (
	"fmt"
	"sort"
)

/*
	golang 中的map是无序的，使用以下方法 可以实现map有序。
	先做排序，排序好以后，再去遍历一次
 */

func main() {
	m := make(map[string]string)
	m["hello"] = "echo hello"
	m["world"] = "echo world"
	m["go"] = "echo go"
	m["is"] = "echo is"
	m["cool"] = "echo cool"

	sorted_keys := make([]string, 0)
	for k, _ := range m {
		sorted_keys = append(sorted_keys, k)
	}

	// sort 'string' key in increasing order
	sort.Strings(sorted_keys)

	for _, k := range sorted_keys {
		fmt.Printf("k=%v, v=%v\n", k, m[k])
	}

}

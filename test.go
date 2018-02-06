package main

import (
	"strconv"
	"fmt"
)

func main() {
	i, err := strconv.Atoi("-42")
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
}

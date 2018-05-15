package main

import (
	"os"
	"fmt"
)

func main() {
	f1, err := os.Stat("/Users/wolf4j/go/src/golang/os/stat.txt")
	if os.IsNotExist(err) {
		return
	}
	fmt.Println(f1.IsDir())
	fmt.Println(f1.Mode())
	fmt.Println(f1.Name())
	fmt.Println(f1.Size())

}

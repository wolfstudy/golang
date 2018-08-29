package main

import (
	"os"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	tmp := make([]byte, 1)
	file, err := os.OpenFile("/Users/wolf4j/go/src/golang/read.txt", 1, os.ModePerm)
	if err != nil {
		panic("wandan")
	}
	n, err := file.Read(tmp)
	if err != nil {
		panic("wandan")
	}
	spew.Dump(n)
}

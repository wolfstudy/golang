package main

import "os"

func main() {
	err := os.Remove("/Users/wolf4j/go/src/golang/os/stat2.txt")
	if err != nil {
		return
	}
}

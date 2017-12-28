package main

import (
	"os"
	"fmt"
)

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func main() {
	f := createFile("/tmp/defer.txt")
	//defer的底层数据结构使用stack来实现的，只有函数在最后执行完的时候，调用它
	defer closeFile(f)
	writeFile(f)
}

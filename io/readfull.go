package main

import (
	"io"
	"os"
	"fmt"
)

//返回的是读到的字节数，io.ReadFull从reader数据源中读取一定的字节数填充到buf中。
func main() {
	buf := make([]byte, 1000)
	file, _ := os.Open("/Users/wolf4j/go/src/golang/gocheck.md")

	n, err := io.ReadFull(file, buf)
	if err != nil {
		fmt.Errorf("err:%v",err)
	}
	fmt.Println(n)
}
//1. 如果 data.txt 是空，则返回 EOF 错误。
//
// 2. 如果 data.txt 少于 10 个字节， 则返回 ErrUnexpectedEOF 错误。并返回读取的字节。
//
// 3. 其他情况，正常返回 10 个字节，错误是 nil 。

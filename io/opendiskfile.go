package main

import (
	"fmt"
	"os"
)

func checkERR(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	//dat, err := ioutil.ReadFile("/tmp/dat")
	//checkERR(err)
	//fmt.Print(string(dat))
	////open只是以只读的方式打开
	//f, err := os.Open("/tmp/dat")
	//checkERR(err)

	//=============================
	//openfile
	// os.O_RDONLY // 只读
	// os.O_WRONLY // 只写
	// os.O_RDWR // 读写
	// os.O_APPEND // 往文件中添建（Append）
	// os.O_CREATE // 如果文件不存在则先创建
	// os.O_TRUNC // 文件打开时裁剪文件
	// os.O_EXCL // 和O_CREATE一起使用，文件不能存在
	// os.O_SYNC // 以同步I/O的方式打开
	//f1, err := os.OpenFile("/tmp/eat", os.O_RDONLY, os.ModePerm)
	//defer f1.Close()
	//checkERR(err)
	//ret, err := f1.Seek(3, 0)
	//bb := make([]byte, 8)
	//checkERR(err)
	//n, err := f1.Read(bb)
	//checkERR(err)
	//fmt.Printf("ret:%d , read:%d, value: %s", ret, n, string(bb))

	f2, err := os.OpenFile("/tmp/eat", os.O_WRONLY, os.ModePerm)
	defer f2.Close()
	checkERR(err)
	ret2, err := f2.Seek(3, 0)
	checkERR(err)
	fmt.Printf("ret2:%d\n", ret2)
	bb2 := []byte("Bytes\n")
	nn2, err := f2.Write(bb2)
	checkERR(err)
	fmt.Printf("value nn2:%v\n", nn2)
	ret22, err := f2.Seek(8, 0)
	checkERR(err)
	fmt.Printf("ret22:%d\n", ret22)
	//=============================

	//b1 := make([]byte, 5)
	//n1, err := f.Read(b1)
	//checkERR(err)
	//fmt.Printf("%d bytes: %s\n", n1, string(b1))
	//
	//o2, err := f.Seek(6, 0)
	//checkERR(err)
	//b2 := make([]byte, 2)
	//n2, err := f.Read(b2)
	//checkERR(err)
	//fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
	//
	//o3, err := f.Seek(6, 0)
	//checkERR(err)
	//b3 := make([]byte, 2)
	//n3, err := io.ReadAtLeast(f, b3, 2)
	//checkERR(err)
	//fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	//
	//_, err = f.Seek(0, 0)
	//checkERR(err)
	//r4 := bufio.NewReader(f)
	//b4, err := r4.Peek(5)
	//checkERR(err)
	//fmt.Printf("5 bytes: %s\n", string(b4))
	//
	//f.Close()
}

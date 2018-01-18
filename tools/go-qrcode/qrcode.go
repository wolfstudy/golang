package main

import (
	"github.com/skip2/go-qrcode"
)

//go-qrcode 是一个基于go的二维码生成库

func main() {
	qrcode.WriteFile("http://www.flysnow.org/",qrcode.Medium,256,"./blog_qrcode.png")
}

/*
生成二维码的函数
func WriteFile(content string, level RecoveryLevel, size int, filename string) error

WriteFile函数的原型定义如上，它有几个参数，大概意思如下：

1，content表示要生成二维码的内容，可以是任意字符串。
2，level表示二维码的容错级别，取值有Low、Medium、High、Highest。
3，size表示生成图片的width和height，像素单位。
4，filename表示生成的文件名路径。
*/

 //自定义二维码

//func main() {
//	qr,err:=qrcode.New("http://www.flysnow.org/",qrcode.Medium)
//	if err != nil {
//		log.Fatal(err)
//	} else {
//		qr.BackgroundColor = color.RGBA{50,205,50,255}
//		qr.ForegroundColor = color.White
//		qr.WriteFile(256,"./blog_qrcode.png")
//	}
//}

package main

import (
	"time"
	"fmt"
)

func main() {
	//初始化定时器
	t := time.NewTimer(2 * time.Second) //返回一个time的指针

	//当前时间
	now := time.Now()
	fmt.Printf("now time: %v", now)

	expire := <-t.C
	fmt.Printf("Expiration time: %v.\n", expire)
}

## 完整的编译过程

1. 预处理： 展开所有头文件和宏定义，删除注释，处理条件指令
2. 编译： 对预处理的文件进行词法、语法、语义分析，生成汇编指令
3. 汇编： 生成机器指令，输出目标文件（.o）
4. 链接： 合并目标文件，生成最终可执行文件

## go的编译参数：

* gcflags 传给编译器的
* ldflags 传给链接器的

* `go tool compile -h` 查看编译命令
* `go tool link -h` 查看链接命令

### 编译和链接

1. 链接器：合并组装多个目标文件，包括地址和空间分配，符号决议和重定位操作。
2. 合并方式：
    * 扫描所有目标文件，合并相似段，并建立全局符号表
    * 通过重定位表重新定位地址
3. go可执行文件默认使用静态链接
4. 静态链接和动态链接的区别

#### 同样打印hello world，go和c生成可执行文件的大小比较

c： 8.5KB
```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# cat hello.c
#include <stdio.h>
#include <stdlib.h>

int main()
{
	printf("hello world!");
	return 0;
}
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# ls -lh hello
-rwxr-xr-x 1 root root 8.5K Aug 30 10:59 hello
```

ldd可以帮助我们去查看代码引用了哪些链接库

```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# ldd hello
	linux-vdso.so.1 =>  (0x00007fff19527000)
	libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fedb0010000)
	/lib64/ld-linux-x86-64.so.2 (0x000055760469b000)
```


go: 2.0MB
```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# cat hello.go
package main

import "fmt"

func main(){
	fmt.Println("hello world!")
}
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# ls -lh hello_go
-rwxr-xr-x 1 root root 2.0M Aug 30 11:01 hello_go
```

原因：c默认使用的是动态链接，但是go使用的是静态链接，它直接把运行时所需要的指令copy到go可执行文件内部。


### 静态编译

`ar rs libtest.a test.o`
```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# ar rs libtest.a test.o
ar: creating libtest.a
```

```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# file libtest.a
libtest.a: current ar archive
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# gcc -o test111 test.c libtest.a
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# ./test111
768
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# ldd test111
	linux-vdso.so.1 =>  (0x00007ffc30b07000)
	libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007efd319d2000)
	/lib64/ld-linux-x86-64.so.2 (0x0000562e3d2b4000
```

### 动态编译

```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# gcc -fPIC -shared -o libtest.so test.c; gcc -I. -o test_so test.c ./libtest.so
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# ls
add  add.go  hello  hello.c  hello.go  hello_go  libtest.a  libtest.so  test  test.c  test.go  test.o  test111  test2  test_so
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# ldd test_so
```

### 压缩go的可执行文件


```
root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# ls -lh hello_go
-rwxr-xr-x 1 root root 1.2M Aug 30 11:22 hello_go

root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# strip hello_go

root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# upx -9 hello_go
                       Ultimate Packer for eXecutables
                          Copyright (C) 1996 - 2013
UPX 3.91        Markus Oberhumer, Laszlo Molnar & John Reiser   Sep 30th 2013

        File size         Ratio      Format      Name
   --------------------   ------   -----------   -----------
   1219240 ->    452068   37.08%  linux/ElfAMD   hello_go

Packed 1 file.

root@iZm5eakzlq9sm4ycm9lw0wZ:~/rxl/go/day01# ls -lh hello_go
-rwxr-xr-x 1 root root 442K Aug 30 11:23 hello_go
```
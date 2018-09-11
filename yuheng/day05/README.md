# GO语言概览

在我们现在的世界，只有两种东西，一种叫做指令，一种叫做数据，除此之外，其他东西都是假的。

test.go

```
var z = 0x100

func main() {
	z++
	fmt.Println(z)// 引用的全局变量

	z := "string" //重定义
	fmt.Println(z)//引用的是局部变量
}
```

编译： `go build -o test test.go && ./test`

```
257
string
```

反汇编：`go tool objdump -s "main\.main" test`

```
 wolf4j@xiaolong  ~/go/src/golang/yuheng/day05   master ●✚ go tool objdump -s "main\.main" test
TEXT main.main(SB) /Users/wolf4j/go/src/golang/yuheng/day05/test.go
  test.go:13            0x1093630               65488b0c25a0080000      MOVQ GS:0x8a0, CX
  test.go:13            0x1093639               483b6110                CMPQ 0x10(CX), SP
  test.go:13            0x109363d               0f8600010000            JBE 0x1093743
  test.go:13            0x1093643               4883ec68                SUBQ $0x68, SP
  test.go:13            0x1093647               48896c2460              MOVQ BP, 0x60(SP)
  test.go:13            0x109364c               488d6c2460              LEAQ 0x60(SP), BP
  test.go:14            0x1093651               488b05707a0900          MOVQ 0x97a70(IP), AX //短地址，在编译器就确定的地址
  test.go:14            0x1093658               48ffc0                  INCQ AX
  test.go:14            0x109365b               488905667a0900          MOVQ AX, 0x97a66(IP)
  test.go:15            0x1093662               48c744245000000000      MOVQ $0x0, 0x50(SP)
  test.go:15            0x109366b               48c744245800000000      MOVQ $0x0, 0x58(SP)
  test.go:15            0x1093674               488d05e5fb0000          LEAQ 0xfbe5(IP), AX
  test.go:15            0x109367b               48890424                MOVQ AX, 0(SP)
  test.go:15            0x109367f               488d05427a0900          LEAQ 0x97a42(IP), AX
  test.go:15            0x1093686               4889442408              MOVQ AX, 0x8(SP)
  test.go:15            0x109368b               e8a093f7ff              CALL runtime.convT2E64(SB)
  test.go:15            0x1093690               488b442418              MOVQ 0x18(SP), AX
  test.go:15            0x1093695               488b4c2410              MOVQ 0x10(SP), CX
  test.go:15            0x109369a               48894c2450              MOVQ CX, 0x50(SP)
  test.go:15            0x109369f               4889442458              MOVQ AX, 0x58(SP)
  test.go:15            0x10936a4               488d442450              LEAQ 0x50(SP), AX
  test.go:15            0x10936a9               48890424                MOVQ AX, 0(SP)
  test.go:15            0x10936ad               48c744240801000000      MOVQ $0x1, 0x8(SP)
  test.go:15            0x10936b6               48c744241001000000      MOVQ $0x1, 0x10(SP)
  test.go:15            0x10936bf               e8bc8effff              CALL fmt.Println(SB)
  test.go:18            0x10936c4               488d05dd090300          LEAQ 0x309dd(IP), AX
  test.go:18            0x10936cb               4889442430              MOVQ AX, 0x30(SP)
  test.go:18            0x10936d0               48c744243806000000      MOVQ $0x6, 0x38(SP)
  test.go:18            0x10936d9               48c744244000000000      MOVQ $0x0, 0x40(SP)
  test.go:18            0x10936e2               48c744244800000000      MOVQ $0x0, 0x48(SP)
  test.go:18            0x10936eb               488d05ee010100          LEAQ 0x101ee(IP), AX
  test.go:18            0x10936f2               48890424                MOVQ AX, 0(SP)
  test.go:18            0x10936f6               488d442430              LEAQ 0x30(SP), AX
  test.go:18            0x10936fb               4889442408              MOVQ AX, 0x8(SP)
  test.go:18            0x1093700               e8ab93f7ff              CALL runtime.convT2Estring(SB)
  test.go:18            0x1093705               488b442410              MOVQ 0x10(SP), AX
  test.go:18            0x109370a               488b4c2418              MOVQ 0x18(SP), CX
  test.go:18            0x109370f               4889442440              MOVQ AX, 0x40(SP)
  test.go:18            0x1093714               48894c2448              MOVQ CX, 0x48(SP)
  test.go:18            0x1093719               488d442440              LEAQ 0x40(SP), AX
  test.go:18            0x109371e               48890424                MOVQ AX, 0(SP)
  test.go:18            0x1093722               48c744240801000000      MOVQ $0x1, 0x8(SP)
  test.go:18            0x109372b               48c744241001000000      MOVQ $0x1, 0x10(SP)
  test.go:18            0x1093734               e8478effff              CALL fmt.Println(SB)
  test.go:19            0x1093739               488b6c2460              MOVQ 0x60(SP), BP
  test.go:19            0x109373e               4883c468                ADDQ $0x68, SP
  test.go:19            0x1093742               c3                      RET
  test.go:13            0x1093743               e8b8abfbff              CALL runtime.morestack_noctxt(SB)
  test.go:13            0x1093748               e9e3feffff              JMP main.main(SB)
  :-1                   0x109374d               cc                      INT $0x3
  :-1                   0x109374e               cc                      INT $0x3
  :-1                   0x109374f               cc                      INT $0x3
```

关掉系统优化开关：`go build -gcflags "-N -l" -o test test.go`
启动`gdb`调试：`gdb test`

## 类型系统
go的类型系统比较好，跨平台特性优良，有指定的int8\16\32\64，int 32位os 4byte，64位 8byte
go里面的引用类型有很大的欺骗性，slice本身就是一个数组。它的传递依然是基于复制的，引用类型是slice内部有一个指针指向另外一段数据，
本身是一个复合结构体。

## 控制语句

* for
* if...else //比较特殊和py有点像，条件这一块可以执行一个初始化操作，也可以执行一个函数
* switch..case

## 函数

* 支持多返回值，但是没有元组类型，所以返回的时候需要去接收，或者可以忽略掉`_`
* 直接返回匿名函数：闭包

## 错误处理

* 输出一个错误的返回值
* panic recover去捕获，类似Java结构化异常的这种错误处理思想。

## 复合类型

* 数组 //go不支持动态数组，在编译器必须确定下来，但是可以使用slice来代替。（数组是一个完整的连续的内存，slice不是数组）
* struct //go里面没有class这种东西，它全部是结构，所以他不是纯粹的这种op语言，只实现了一部分功能，基于接口实现的多态不是一回事
* 我们可以给其（struct）定义相对应的方法，类似于this指针的传递。
* 什么时候用方法？什么时候用函数？

```
type User struct{
    name string
    age int
}

//方法属于类型，类型不需要实例化
func (u User) test(x int){

}

//函数没有实例化一说
func test(u User, x int){

}
```
上面两种实现的方式都实现一个功能。

## 接口

* 接口代表了一种调用的约定，op语言需要显示声明接口的内容，多态里面有一个继承体系，子类可以拥有父类的全部属性和方法（子类可以以父类的身份出现-->多态）
* 抽象类：只能被继承，不能被初始化，没有办法实例化
* 抽象类和接口有什么不同？
    *
* override、 new、 overload、 virtual_method


## 并发

* mutliproccess
* thread
* gorutine

# go开发环境

## 工作空间

工作空间：有三个目录组成：
* src 存放源码目录
* bin 存放编译后的二进制文件
* pkg 存放`.a`文件

**是为每一个项目创建一个工作空间，还是所有的都放到一起。**

* go build: 不管你是否存在增量，都会重新编译一遍
* go install：增量编译,你需要知道上一次编译的结果放在哪里，有一个diff的过程，如果没有修改，就不会重新编译，intsall会把文件copy到pkg的目录下。它是通过`buildid`来
判断是否有过修改。增量编译的好处是，当项目比较大的时候，我们没必要花费更多的时间来重新编译。直接引用上次保存的`.a`文件。
    * 这里需要区分什么是编译什么是链接，只有地址改变的时候，才需要链接。
> 加 -x 参数可以查看编译的详细过程

## 编译参数

go build

```
-o 可执行文件的文件名（默认与目录同名）
-a 强制重新编译所有包（含标准库）
-p 并行编译所使用的CPU的核数量
-v 显示需要编译的包的名字
-n 仅显示编译命令，但不执行
-x 显示正在执行的编译命令
-work 显示临时工作目录，完成后不删除
-race 启动数据竞争检查（仅支持amd64）
-gcflags 编译器参数
-ldflags 链接器参数
```


gcflags:`go tool compile -h`

```
-B 禁用越界检查
-N 禁用优化
-l 禁用内联
-u 禁用unsafe
-S 输出汇编代码
-m 输出优化信息
```

ldflags：`go tool link -h`

```
-s 禁用符号表
-w 禁用DRAWF调试信息
-X 设置字符串全局变量值   -X ver="0.99"
-H 设置可执行文件格式    -H windowsgui
```

清理所有编译后的数据：`go clean`

> -x 可以输出详细信息，看究竟删除了一些什么东西

## 交叉编译

跨平台执行，使用环境变量指定你执行的操作系统。

`GOOS=windows go build` win-x86

```
day05.exe: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows
```

`GOOS=windows GOARCH=386 go build -o test.exe`

`GOOS=darwin go build`

## 依赖包下载

`go get + path`

> -x -v 可以输出详细信息

下载的时候，它默认会帮你执行编译，编译后的东西放到pkg文件下，这个时候没必要我们在手动go build去执行。

```
 wolf4j@xiaolong  ~/go  cd pkg/
 wolf4j@xiaolong  ~/go/pkg  ll
total 0
drwxr-xr-x  12 wolf4j  staff   384B  6  8 14:17 darwin_amd64
```

交叉编译的时候我们需要区分不同的os，所以需要使用`darwin_amd64`这样一个目录来区分。


# go 测试

## 单元测试

数据和逻辑测试分离

推荐包：`github.com/smartystreets/goconvey/convey `

## 性能测试

待续...

## 调试

GDB delve

```
l  显示源码
b 打断点
r 执行
info args/locals 打印相关信息
c 继续执行
p 打印
x 查看内存数据
disass 反汇编
info files 显示源码的信息

source /usr/local/go/src/runtime/runtime-gdb.py 载入gdb给go包装的一些命令（一般没必要）
```

## 检查

#### 代码检查：

* linter `gometalinter`

#### 依赖包检查

* depscheck

# go 杂记

## 从0开始学习go

首先需要一个工作空间

可执行文件的入口包使用的是`main`，main包是有特殊含义的，main函数没有参数，没有返回值。

### 变量

- 局部变量
- 全局变量 （var name string）

语法糖： var x = 0x100（交给编译器去推导类型）

- 定义：定义明确告诉编译器需要为这个变量分配一块内存，内存的大小是这个类型的大小。是有一个明确的内存分配行为。
- 声明：在某个地方已经有人定义好了这么一个变量，你可以直接使用，声明是那条狗是你们家的，前提是这条狗已经存在了。


```
package main

var z int

func main() {
	z = 123
	println(z)//引用全局变量

	z := "abc"
	println(z) //引用当前局部变量
}
```

输出：

```
123
abc
```

短地址一般是在编译期间就确定的。

c语言的类型系统是比较混乱的，需要根据os的位数来进行判断。go在类型上就是很确定的，类型声明足够清楚，所以跨平台特性做的不错。这个东西和指针有关系，指针的长度一般和地址总线是相等的，通常会和整数类型相同。

###### 引用类型

有很大的欺骗性，切片本身就是一个结构体。指的是内部引用了一个指针指向另一个数据，本身是一个复合结构体。所谓的引用类型，就是初始化的时候，用一个make函数把内部的属性初始化，也类似于一个语法糖。

- slice
- map
- channel

###### 控制语句

- if （类似与python，可以允许有初始化的操作）

```
package main

var z int

func main() {
	z = 123
	if println(z); z>0{
		z++
	}
}
```

- for
- switch


go吸收了很多动态语言的做法，支持多个返回值，go语言没有元组类型，所以你的返回值必须要有人去接收，或者用`_` 来忽略掉这个返回值。

###### 闭包函数


```
package main

func test(x int) func() {
	return func() {
		println(x)
	}
}

func main() {
	test(100)()
}
```

###### 错误处理

使用 panic来处理，recover来捕获这个异常，这个有点像`try...except`还有一个defer语句有点像`finally`,defer使用的是栈的结构来实现的。

###### 基本类型

- 数组（go不支持动态数组，编译期间就必须确定数组的长度，但是我们可以使用切片来实现动态数组的功能，切片的底层实现是一个数组。）

go只实现了很有限的封装，没有op语言的那种继承和多态，至于通过结构实现的继承是另一回事。

什么时候用方法，什么时候用函数？


```
type User struct {
	name string
	age  int
}

func (u User) test(x int) {
	u.age += x
}

func test(u User, x int) {
	u.age += x
}
```

- 方法：属于类型，
- 函数：

**只有对象才需要被实例化**


###### 接口

多态里面有一个继承体系，子类可以以父类的身份出现。

函数调用的时候利用op语言的多态，还有一种是利用接口，接口不能初始化，不能实例化。还有一个东西叫抽象类。

什么时候用接口，什么时候用抽象类？

类是继承，接口是实现。

op语言的一些关键字：

- override
- overload
- virtual method：
- new
- abstract method：


###### 并发

go语言把这件事情做的精巧，它默认为所有的事情都是并发的。

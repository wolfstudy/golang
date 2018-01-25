## unsafe使用

**三个函数：**

 * func Alignof（variable ArbitraryType）uintptr
 * func Offsetof（selector ArbitraryType）uintptr
 * func Sizeof（variable ArbitraryType）uintptr
 
 ```$xslt
sizeof()函数返回操作数在内存中的字节大小，参数可以是任意类型的表达式，   
但是它并不会对表达式进行求值。一个Sizeof函数调用是一个对应  
uintptr类型的常量表达式，因此返回的结果可以用作数组类型的  
长度大小，或者用作计算其他的常量。
```

**一种类型：**

 * 类型Pointer * ArbitraryType
 
> 1、这里，ArbitraryType不是一个真正的类型，它只是一个占位符。  
  2、与Golang中的大多数函数不同，上述三个函数的调用将始终在编译时求值，而不是运行时。
     这意味着它们的返回结果可以分配给常量。

 

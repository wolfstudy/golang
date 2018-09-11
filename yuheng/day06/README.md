寄存器很特殊，没有地址的概念，我们简单给它起了一个名字：BP、SP、AX、BX等，它有一些特殊的规则，由编译器和cpu来决定的。

Native code 就是一些二进制的数据，cpu只能处理二进制的数据，其他数据在cpu看来是不存在的

汇编才是这个世界上最简单的语言，在汇编的世界里面只有数据，没有什么类、结构体的抽象，他只是把一个地方的数据搬移到另一个地方。

## 寄存器

* 通用寄存器：AX, BX, CX, SI, DI(BP, SP) 有8个
* 堆栈寄存器：BP（基地址）SP（栈顶）维护函数调用时候的堆栈
* 指令寄存器：IP《你要告诉cpu（cpu不知道上下文，需要有人来记录这个状态），下一条执行的指令在哪里》、PC（程序计数器）==CS:IP 《地址》
* 段寄存器：CS（指向代码段）DS（指向数据段）
* RAX、EAX、AX、AH、AL

![image.png](https://upload-images.jianshu.io/upload_images/6967649-427fd587b9939ff8.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

一个CPU去读取一个指令的时候，需要三条总线：地址、数据、控制

8086寄存器的寻址空间只有64KB，那么大于64KB的数据如何读取超过64KB的地址呢。可以采用段+偏移量

c数据+代码，这是cpu执行的=两种东西，

* CS+EA（偏移量）===》.text
* data ===>DS+EA

指针是标准一个常量，地址这东西并不存在，只是一个抽象的数字，指针通常是需要分配内存的，我们把这个抽象的数字存到某一个地方就是指针。

![image.png](https://upload-images.jianshu.io/upload_images/6967649-4d353cbc2b372a92.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)







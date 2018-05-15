package main

func main() {

}

/*
new and make 都是用来分配空间。
内置函数 new 分配空间。传递给new 函数的是一个类型，不是一个值。返回值是 指向这个新分配的零值的指针。
内建函数 make 分配并且初始化 一个 slice, 或者 map 或者 chan 对象。 并且只能是这三种对象。 和 new 一样，第一个参数是 类型，不是一个值。
但是make 的返回值就是这个类型（即使一个引用类型），而不是指针。 具体的返回值，依赖具体传入的类型。

Slice : 第二个参数 size 指定了它的长度，此时它的容量和长度相同。
你可以传入第三个参数 来指定不同的容量值，但是必须不能比长度值小。
比如: make([]int, 0, 10)
Map: 根据size 大小来初始化分配内存，不过分配后的 map 长度为0。 如果 size 被忽略了，那么会在初始化分配内存的时候 分配一个小尺寸的内存。
Channel: 管道缓冲区依据缓冲区容量被初始化。如果容量为 0 或者被 忽略，管道是没有缓冲区的。

new 的作用是 初始化 一个指向类型的指针 (*T)， make 的作用是为 slice, map 或者 channel 初始化，并且返回引用 T
make(T, args)函数的目的与new(T)不同。它仅仅用于创建 Slice, Map 和 Channel，并且返回类型是 T（不是T*）的一个初始化的（不是零值）的实例。

++++++++++++++++++++++++++++++
    func newInt *int {
        var i int
        return &i   //为何可以返回局部变量呢？
    }
    someInt := newInt()

go有逃逸分析，虽然表面上看i是在栈上进行分配的，但是当你return的时候，编译器会检测，将变量i逃逸到栈上进行分配（在函数内联的时候会进行处理）

*/
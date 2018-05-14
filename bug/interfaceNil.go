package main

import "fmt"

/*
Nil接口与接口中有一个零指针不同。
 */

type Cat interface {
	Meow()
}

type Tabby struct{}

func (*Tabby) Meow() {
	fmt.Println("meow")
}

func GetACat() Cat {
	var myTabby *Tabby = nil
	// Oops, we forgot to set myTabby to a real value
	return myTabby
}

func main() {
	if GetACat() == nil {
		fmt.Errorf("Forgot to return a real cat")
	}
}

/*
Guess what? 上面的测试不会检测到零指针！这是因为接口用作指针，所以GetACat有效地返回一个指向零指针的指针。千万不要做GetACat做的操作。
使用错误值很容易做到这一点。


Under the covers，接口被实现为两个元素，一个类型和一个值。该值称为接口的动态值，是一个任意的具体值，类型是值的类型。
例如：对于int 3，用接口表示如下：（int，3）。
只有在内部值和类型都未设置时（nil，nil），接口值才为零。特别注意，一个nil接口将始终保持为nil类型。
如果我们在接口值中存储一个类型为* int的零指针，则无论指针的值是什么，内部类型都将是* int：（* int，nil）。
因此，即使内部指针为零，这样的接口值也不会为零。
这种情况可能会引起混淆，并且在接口值（如错误返回值）内存储一个零值时会出现这种情况：
伪代码：
func returnError() error {
	var p *MyError = nil
	if bad() {
		p = ErrBad
	}
	return p // Will always return a non-nil error.
}
如果一切顺利，函数返回一个nil p，所以返回值是一个error的接口值（* MyError，nil）。这意味着如果调用者将返回的错误与nil比较，
即使没有发生任何错误，它总会看起来好像有错误。为了给调用者返回一个适当的nil错误，函数必须返回一个明确的nil：

func returnsError() error {
	if bad() {
		return ErrBad
	}
	return nil
}

对于返回错误的函数总是在签名中使用错误类型而不是像* MyError这样的具体类型，以帮助确保正确创建错误，这是一个好主意。
作为一个例子，os.Open返回一个错误，即使不是nil，它总是具体的类型* os.PathError。

无论何时使用接口，都会出现类似于这里描述的情况。请记住，如果程序中存储了任何具体值，程序将不会为nil。

 */
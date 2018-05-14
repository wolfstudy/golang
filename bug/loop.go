package main

import "fmt"

//在golang中有三个地方比较容易犯错。


/*
Loop variables are scoped outside the loop.（循环变量的范围跑到循环之外了）
 */

func pprint(pi *int)  {
	fmt.Println("pprint 打印。。。",*pi)
}

func main() {

	for i := 0; i < 10; i++ {
		defer fmt.Println("第一个print。。。",i)// OK; prints 9 ... 0
		defer func() {
			fmt.Println("第二个print。。。",i)//WRONG; prints "10" 10 times
		}()

		defer func(i int) {
			fmt.Println("第三个print。。。",i)// OK
		}(i)

		defer pprint(&i)// WRONG; prints "10" 10 times
		go fmt.Println("第四个print。。。",i)// OK; prints 0 ... 9 in unpredictable order
		go func() {
			fmt.Println("第五个print。。。",i)// WRONG; totally unpredictable.
		}()
	}

}

/*
每个人都希望这些值被限制在循环内部，但是每次迭代都会重复使用相同的内存位置。这意味着你绝不能让上面的key，value或i的地址逃脱循环。
匿名函数（）{/ *用i * /}做一些事情（一个“闭包”）是一个微妙的方式来获取一个变量的地址
 */
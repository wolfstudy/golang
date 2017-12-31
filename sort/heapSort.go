package main

import "fmt"

func main() {
	a := []int{34, 52, 12, 45, 56, 10, 35}
	HeatSort(a)
	fmt.Println(a)
}

//交换
func swap(i, j int,a[] int)  {
	a[i], a[j] = a[j], a[i];
}

//调整节点
func headAdjust(a []int,start, end int)  {
	temp := a[start]

	for i := 2*start+1; i <= end; i*=2  {
		// 左右孩子的节点 2i+1. 2i+2

		//选择出左右孩子较大的下标
		if i < end && a[i] < a[i+1] {
			i++
		}

		if temp <= a[i] {
			break //已为大顶堆 保持稳定性
		}

		a[start] = a[i] //将子节点上移
		start = i //重置，开启下一轮
	}

	a[start] = temp //插入正确的位置
}

//节点排序
func HeatSort(a []int) {
	length := len(a)
	if length == 0 {
		return
	}

	for i := length / 2; i >= 0; i-- {
		headAdjust(a, i, length-1)
	}

	for j := length - 1; j >= 0; j-- {
		swap(0, j, a)
		headAdjust(a, 0, j-1)
	}
}

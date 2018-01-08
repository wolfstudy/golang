package main

import "fmt"

/*
堆是一棵顺序存储的完全二叉树。
其中每个结点的关键字都不大于其孩子结点的关键字，这样的堆称为小根堆。
其中每个结点的关键字都不小于其孩子结点的关键字，这样的堆称为大根堆。
举例来说，对于n个元素的序列{R0, R1, ... , Rn}当且仅当满足下列关系之一时，称之为堆：
(1) Ri <= R2i+1 且 Ri <= R2i+2 (小根堆)
(2) Ri >= R2i+1 且 Ri >= R2i+2 (大根堆)
*/

func main() {
	a := []int{34, 52, 12, 45, 56, 10, 35}
	HeatSort(a)
	fmt.Println(a)
}

//交换函数
func swap(i, j int, a [] int) {
	a[i], a[j] = a[j], a[i];
}

//调整节点
func headAdjust(a []int, start, end int) {
	temp := a[start]

	for i := 2*start + 1; i <= end; i *= 2 {
		// 左右孩子的节点 2i+1. 2i+2

		//选择出左右孩子较大的下标
		if i < end && a[i] < a[i+1] {
			i++
		}

		if temp <= a[i] {
			break //已为大顶堆 保持稳定性
		}

		a[start] = a[i] //将子节点上移
		start = i       //重置，开启下一轮
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

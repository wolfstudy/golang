package main

import "fmt"

/*
该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。

归并排序其实要做两件事：
（1）“分解”——将序列每次折半划分。
（2）“合并”——将划分后的序列段两两合并后排序。
我们先来考虑第二步，如何合并？
在每次合并过程中，都是对两个有序的序列段进行合并，然后排序。
这两个有序序列段分别为 R[low, mid] 和 R[mid+1, high]。
先将他们合并到一个局部的暂存数组R2中，带合并完成后再将R2复制回R中。
为了方便描述，我们称 R[low, mid] 第一段，R[mid+1, high] 为第二段。
每次从两个段中取出一个记录进行关键字的比较，将较小者放入R2中。最后将各段中余下的部分直接复制到R2中。
经过这样的过程，R2已经是一个有序的序列，再将其复制回R中，一次合并排序就完成了。
 */


//从上到下排序
func MergeSortUTD(s []int) {
	aux := make([]int, len(s))
	mergeSortUTD(s, aux, 0, len(s)-1)

}

//从下到上排序
func MergeSortDTU(s []int) {
	aux := make([]int, len(s)) //辅助切片
	n := len(s)
	for sz := 1; sz < n; sz *= 2 {
		for lo := 0; lo < n-sz; lo += 2 * sz {
			merge(s, aux, lo, lo+sz-1, min(lo+2*sz-1, n-1))
		}
	}
}

func mergeSortUTD(s, aux []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := (lo + hi) >> 1
	mergeSortUTD(s, aux, lo, mid)
	mergeSortUTD(s, aux, mid+1, hi)
	merge(s, aux, lo, mid, hi)
}

//真正交换的过程
func merge(s, aux []int, lo, mid, hi int) {
	for k := lo; k <= hi; k++ {
		aux[k] = s[k]
	}

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			s[k] = aux[j]
			j++
		} else if j > hi {
			s[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			s[k] = aux[j]
			j++
		} else {
			s[k] = aux[i]
			i++
		}
	}

}

//求最小值函数
func min(i, j int) int {
	if j < i {
		return j
	}
	return i
}

//测试
func main() {
	fmt.Println("排序前：")
	s := []int{9,0,6,5,8,2,1,7,4,3}
	fmt.Println(s)
	MergeSortUTD(s)
	fmt.Println("排序后：")
	//MergeSortDownToUp(s)
	fmt.Println(s)
}

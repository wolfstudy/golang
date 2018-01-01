package main

import "fmt"

/*
 缩小增量排序
先将整个待排元素序列分割成若干个子序列（由相隔某个“增量”的元素组成的）分别进行直接插入排序，
然后依次缩减增量再进行排序，待整个序列中的元素基本有序（增量足够小）时，
再对全体元素进行一次直接插入排序。因为直接插入排序在元素基本有序的情况下（接近最好情况），效率是很高的，
 */


 //交换
func swap1 (i, j int,a []int)  {
	a[i], a[j] = a[j], a[i]
}

//排序
func ShellSort(a []int){
	n := len(a)
	
	h := 1
	for h < n/3 {
		h = 3*h +1  //寻找合适的时间间隔h
	}

	for h >= 1 {
		//将数组变为间隔h个元素有序
		for i := h; i < n; i++ {
			//间隔h插入排序
			for j := i; j >= h && a[j] < a[j - h] ; j -= h {
				swap1(j, j-h, a)
			}
		}

		h /= 3
	}
}

//测试
func main() {
	s := []int{9,0,6,5,8,2,1,7,4,3}
	fmt.Println(s)
	ShellSort(s)
	fmt.Println(s)
}
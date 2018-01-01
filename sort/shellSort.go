package main

import "fmt"

func swap1 (i, j int,a []int)  {
	a[i], a[j] = a[j], a[i]
}

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

func main() {
	s := []int{9,0,6,5,8,2,1,7,4,3}
	fmt.Println(s)
	ShellSort(s)
	fmt.Println(s)
}
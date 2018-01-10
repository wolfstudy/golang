package main

import "fmt"

func insertSort(nums []int)  {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			j := i - 1
			temp := nums[i]
			for j >= 0 && nums[j] > temp {
				nums[j+1] = nums[j]
				j--
			}
			nums[j+1] = temp
		}
	}
}

func main() {
	nums := []int{3,6,1,5,7,2,8,4,9}
	fmt.Println("初始状态：",nums)
	insertSort(nums)
	fmt.Println("排序后状态：",nums)
}
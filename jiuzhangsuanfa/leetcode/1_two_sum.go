package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	var ret []int

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				ret = []int{i, j}
			}
		}
	}
	return ret
}

//func MapTwoSum(nums []int, target int) []int {
//	m := make(map[int]int)
//	var ret []int
//
//	for key, value := range nums {
//		m[key] = value
//	}
//}

func main() {
	nums := []int{2, 7, 11, 15}
	ret := TwoSum(nums, 9)
	fmt.Println(ret)
}

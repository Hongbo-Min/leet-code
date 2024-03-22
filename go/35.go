package main

import "fmt"

func InsertLocation(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			return middle
		}
	}
	return right + 1
}

// 搜索插入位置
func main() {
	nums := []int{1, 3, 5, 6}
	target := 8
	res := InsertLocation(nums, target)
	fmt.Printf("res: %v\n", res)
}

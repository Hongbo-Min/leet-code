package main

import (
	"fmt"
)

func getRightBorder(nums []int, target int) int {
	left, right := 0, len(nums)-1
	rightBorder := -2
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
			rightBorder = left
		}
	}
	return rightBorder
}

func getLeftBorder(nums []int, target int) int {
	left, right := 0, len(nums)-1
	leftBorder := -2
	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] >= target {
			right = middle - 1
			leftBorder = right
		} else {
			left = middle + 1
		}
	}
	return leftBorder
}

func searchRange(nums []int, target int) []int {
	rightBorder := getRightBorder(nums, target)
	leftBorder := getLeftBorder(nums, target)
	if rightBorder-leftBorder > 1 {
		return []int{leftBorder + 1, rightBorder - 1}
	}
	if leftBorder == -2 || rightBorder == -2 {
		return []int{-1, -1}
	}
	return []int{-1, -1}
}

// 有序数组中查找元素区间
func main() {
	nums := []int{1}
	target := 1
	res := searchRange(nums, target)
	fmt.Printf("res: %#v\n", res)
}

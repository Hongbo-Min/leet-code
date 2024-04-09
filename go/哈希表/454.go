package main

import "fmt"

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	res := 0
	mark := make(map[int]int)
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			mark[num1+num2]++
		}
	}
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			res += mark[-num3-num4]
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{-1, -2}
	nums3 := []int{-1, 2}
	nums4 := []int{0, 2}
	fmt.Printf("fourSumCount(nums1, nums2, nums3, nums4): %v\n", fourSumCount(nums1, nums2, nums3, nums4))
}

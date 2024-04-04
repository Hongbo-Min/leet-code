package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]struct{})
	for _, num := range nums1 {
		nums1Map[num] = struct{}{}
	}
	res := []int{}
	resMap := make(map[int]struct{})
	for _, num := range nums2 {
		if _, ok := nums1Map[num]; ok {
			resMap[num] = struct{}{}
		}
	}
	for key := range resMap {
		res = append(res, key)
	}
	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Printf("intersection(nums1, nums2): %v\n", intersection(nums1, nums2))
}

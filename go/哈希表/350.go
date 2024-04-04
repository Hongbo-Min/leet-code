package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	nums1Map := make(map[int]int)
	for _, num := range nums1 {
		nums1Map[num]++
	}
	res := []int{}
	for _, num := range nums2 {
		if count, ok := nums1Map[num]; ok {
			if count >= 1 {
				res = append(res, num)
				nums1Map[num]--
			}
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Printf("intersect(nums1, nums2): %v\n", intersect(nums1, nums2))
}

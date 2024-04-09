package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		n1 := nums[i]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {

			n2 := nums[j]
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				n3, n4 := nums[left], nums[right]
				sum := n1 + n2 + n3 + n4
				if sum == target {
					res = append(res, []int{n1, n2, n3, n4})
					for left < right && nums[left] == n3 {
						left++
					}
					for left < right && nums[right] == n4 {
						right--
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return res
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Printf("fourSum(nums, target): %v\n", fourSum(nums, target))
}

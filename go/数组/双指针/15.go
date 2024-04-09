package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		num1 := nums[i]
		if num1 > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left, right := i+1, len(nums)-1; left < right; {
			num2, num3 := nums[left], nums[right]
			if num1+num2+num3 == 0 {
				res = append(res, []int{num1, num2, num3})
				for left < right && nums[left] == num2 {
					left++
				}
				for left < right && nums[right] == num3 {
					right--
				}
			} else if num1+num2+num3 < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("threeSum(nums): %v\n", threeSum(nums))
}

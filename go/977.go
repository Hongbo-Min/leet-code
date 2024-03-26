package main

import "fmt"

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	k := len(res) - 1
	for i, j := 0, k; i <= j; {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			res[k] = nums[i] * nums[i]
			i++
			k--
		} else {
			res[k] = nums[j] * nums[j]
			j--
			k--
		}
	}
	return res
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Printf("sortedSquares(nums): %v\n", sortedSquares(nums))
}

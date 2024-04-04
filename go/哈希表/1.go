package main

import "fmt"

func twoSum(nums []int, target int) []int {
	mark := make(map[int]int)
	for i, num := range nums {
		need := target - num
		if index, ok := mark[need]; ok {
			return []int{i, index}
		}
		mark[num] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Printf("twoSum(nums, 9): %v\n", twoSum(nums, 9))
}

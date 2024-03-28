package main

import "fmt"

func removeElement(nums []int, val int) int {
	slowPointer := 0
	for fastPointer := 0; fastPointer < len(nums); fastPointer++ {
		if nums[fastPointer] != val {
			nums[slowPointer] = nums[fastPointer]
			slowPointer++
		}
	}
	return slowPointer
}

// 移除元素-双指针
func main() {
	val := 2
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Printf("removeElement(nums, val): %v\n", removeElement(nums, val))
}

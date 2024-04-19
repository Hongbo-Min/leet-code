package main

import "fmt"

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	val := nums[left]
	i := left
	j := right
	for i < j {
		for i < j && nums[j] >= val {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] < val {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = val
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func main() {
	nums := []int{1, 6, 2, 8, 9, 10, 3, 4}
	quickSort(nums, 0, len(nums)-1)
	for i, num := range nums {
		fmt.Printf("%d : %d\n", i, num)
	}
}

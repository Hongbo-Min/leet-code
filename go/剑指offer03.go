/* 找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
*/

/*
输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3
*/

package main

import "fmt"

func findRepeatNumber(nums []int) (res int) {
	numberMap := make(map[int]int)
	for _, v := range nums {
		numberMap[v]++
		if numberMap[v] >= 2 {
			res = v
			return
		}
	}

	return
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	res := findRepeatNumber(nums)
	fmt.Printf("res: %v\n", res)
}

package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt
	curWinCount := 0
	leftIndex := 0
	for rightIndex := 0; rightIndex < len(nums); rightIndex++ {
		curWinCount += nums[rightIndex]
		for curWinCount >= target {
			subLen := rightIndex - leftIndex + 1
			if subLen < res {
				res = subLen
			}
			curWinCount -= nums[leftIndex]
			leftIndex++
		}
	}
	if res == math.MaxInt {
		return 0
	}
	return res
}

func main() {
	println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

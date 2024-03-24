package main

import "fmt"

func isPerfectSquare(num int) bool {
	left, right := 1, num
	for left <= right {
		middle := left + (right-left)/2
		count := middle * middle
		if count == num {
			return true
		} else if count > num {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return false
}

// 有效的完全平方数
func main() {
	fmt.Printf("isPerfectSquare(16): %v\n", isPerfectSquare(16))
}

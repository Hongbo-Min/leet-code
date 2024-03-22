package main

import "fmt"

func mySqrt(x int) int {
	left, right := 1, x
	for left <= right {
		middle := left + (right-left)/2
		if middle*middle > x {
			right = middle - 1
		} else if middle*middle < x {
			left = middle + 1
		} else {
			return middle
		}
	}
	return right
}

// x的平方根
func main() {
	fmt.Printf("mySqrt(8): %v\n", mySqrt(8))
}

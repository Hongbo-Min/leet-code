package main

import "fmt"

func calcNum(n int) int {
	var res int
	for n > 0 {
		num := n % 10
		res += num * num
		n /= 10
	}
	return res
}

func isHappy(n int) bool {
	mark := make(map[int]struct{})
	for {
		num := calcNum(n)
		if num == 1 {
			return true
		}
		if _, ok := mark[num]; ok {
			return false
		} else {
			mark[num] = struct{}{}
		}
		n = num
	}
}

func main() {
	fmt.Printf("isHappy(19): %v\n", isHappy(19))
}

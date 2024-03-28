package main

import "fmt"

func totalFruit(fruits []int) int {
	res := -1
	count := make(map[int]int)
	left := 0
	for right, k := range fruits {
		count[k]++
		for len(count) > 2 {
			kind := fruits[left]
			count[kind]--
			if count[kind] == 0 {
				delete(count, kind)
			}
			left++
		}
		res = max(res, right-left+1)
		fmt.Printf("count: %v\n", count)
	}
	return res
}

func main() {
	println(totalFruit([]int{1, 0, 1, 4, 1, 4, 1, 2, 3}))
}

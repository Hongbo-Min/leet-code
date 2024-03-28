package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, 0)
	top, bottom := 0, m-1
	left, right := 0, n-1
	num := 1
	target := m * n
	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
			num++
			if num > target {
				return res
			}
		}
		top++
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
			num++
			if num > target {
				return res
			}
		}
		right--
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
			num++
			if num > target {
				return res
			}
		}
		bottom--
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
			num++
			if num > target {
				return res
			}
		}
		left++
	}
	return res
}

func main() {
	fmt.Printf("spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}): %v\n", spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}
